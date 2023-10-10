package stats

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fbiville/markdown-table-formatter/pkg/markdown"
	"github.com/go-bridget/mig/db"

	"github.com/TykTechnologies/exp/cmd/go-fsck/model"
)

func stats(cfg *options) error {
	// Read the exported go-fsck.json data.
	defs, err := model.ReadFile(cfg.inputFile)
	if err != nil {
		return err
	}

	// Loop through function definitions and collect referenced
	// symbols from imported packages. Globals may also reference
	// imported packages so this is incomplete at the moment.

	refs := []SymbolReference{}

	for _, def := range defs {
		for _, fn := range def.Funcs {
			symbols := listUsedSymbols(fn)
			if len(symbols) > 0 {
				refs = append(refs, symbols...)
			}
		}

		// Convert package refs into full import paths.
		if cfg.full {
			importMap := def.Imports.Map()
			for k, v := range refs {
				long, ok := importMap[v.Import]
				if ok {
					refs[k].Import = long
				}
			}
		}
	}

	ctx := context.Background()

	// Aggregations are easier in SQL... the following block of
	// code uses an sqlite in-memory database to do some math.
	conn, err := db.ConnectWithOptions(ctx, db.Options{
		Credentials: db.Credentials{
			DSN:    ":memory:",
			Driver: "sqlite",
		},
	})

	if err != nil {
		return err
	}

	createSymbolTables := strings.Join([]string{
		"CREATE TABLE IF NOT EXISTS symbol_reference (",
		"id integer primary key,",
		"import text,",
		"symbol text,",
		"used_by text",
		");",
	}, "\n")

	conn.MustExec(createSymbolTables)

	sql := "insert into symbol_reference (id, import, symbol, used_by) values ($1, $2, $3, $4)"
	id := 1
	for _, ref := range refs {
		conn.MustExec(sql, id, ref.Import, ref.Symbol, ref.UsedBy)
		id++
	}

	results := []struct {
		Import string
		Symbol string
		Used   int `db:"ref_count"`
	}{}

	if cfg.filter != "" {
		sql = "select import, symbol, count(used_by) ref_count from symbol_reference  where import like ? group by import, symbol order by ref_count desc"
		if err := conn.Select(&results, sql, "%"+cfg.filter+"%"); err != nil {
			return err
		}
	} else {
		sql = "select import, symbol, count(used_by) ref_count from symbol_reference group by import, symbol order by ref_count desc"
		if err := conn.Select(&results, sql); err != nil {
			return err
		}
	}

	// Encode aggregated results as json.
	if cfg.json {
		b, err := json.Marshal(results)
		if err != nil {
			return err
		}
		fmt.Println(string(b))
		return nil
	}

	// Encode aggregated results as markdown.
	table := [][]string{}
	for _, result := range results {
		table = append(table, []string{result.Import, result.Symbol, fmt.Sprint(result.Used)})
	}

	t, err := markdown.NewTableFormatterBuilder().WithPrettyPrint().Build("Import", "Symbol", "Used").Format(table)
	if err != nil {
		return err
	}

	fmt.Println(t)

	return nil
}

func listUsedSymbols(decl *model.Declaration) []SymbolReference {
	if len(decl.References) == 0 {
		return nil
	}

	symbols := []SymbolReference{}
	for pkg, refs := range decl.References {
		for _, ref := range refs {
			symbols = append(symbols, SymbolReference{
				Import: pkg,
				Symbol: ref,
				UsedBy: decl.Name,
			})
		}
	}

	return symbols
}
