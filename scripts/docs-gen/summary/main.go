package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"os"
	"sort"
	"strings"
	"unicode"

	"github.com/spf13/pflag"
)

func main() {
	_ = start()
}

type config struct {
	exclude string
	noSort  string
}

func start() error {
	var cfg config

	pflag.StringVar(&cfg.exclude, "exclude", "", "Exclude version from docs")
	pflag.StringVar(&cfg.noSort, "no-sort", "Config", "Exclude type from sorting")
	pflag.Parse()

	if err := processConfig(cfg); err != nil {
		return err
	}

	return nil
}

type TypeDeclaration struct {
	Added   []string `json:"added"`
	Removed []string `json:"removed"`

	Name     string `json:"name"`
	Path     string `json:"path"`
	Type     string `json:"type"`
	Tag      string `json:"tag"`
	Doc      string `json:"doc"`
	JSONName string `json:"json_name"`

	Seen bool `json:"-"`
}

func processConfig(cfg config) error {
	data := []*TypeDeclaration{}

	b, err := os.ReadFile("tyk-config-index.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	noSort := []*TypeDeclaration{}
	mustSort := []*TypeDeclaration{}

	for _, row := range data {
		if row.Path == cfg.noSort || strings.HasPrefix(row.Path, cfg.noSort+".") {
			noSort = append(noSort, row)
		} else {
			mustSort = append(mustSort, row)
		}
	}

	sort.Slice(mustSort, func(i, j int) bool {
		return mustSort[i].JSONName < mustSort[j].JSONName
	})

	data = append(noSort, mustSort...)

	listPath(data, "Config.", "", cfg.exclude)
	return nil
}

func listPath(data []*TypeDeclaration, pathPrefix string, prefix string, exclude string) {
	for _, def := range data {
		if def.JSONName == "" {
			continue
		}

		var typeName string
		if prefix != "" {
			typeName = prefix + "." + def.JSONName
		} else {
			typeName = def.JSONName
		}

		var envTypeName string
		if prefix != "" {
			envTypeName = prefix + "." + def.Name
		} else {
			envTypeName = def.Name
		}

		if strings.HasPrefix(def.Path, pathPrefix) {
			if def.Doc != "" {
				versions := []string{}
				removed := []string{}

				for _, version := range def.Added {
					if version == exclude && len(def.Removed) == 0 {
						continue
					}

					// Sanitize to minor version
					if strings.HasSuffix(version, ".0") {
						version = version[0 : len(version)-2]
					}

					versions = append(versions, "`"+version+"`")
				}

				for _, version := range def.Removed {
					// Sanitize to minor version
					if strings.HasSuffix(version, ".0") {
						version = version[0 : len(version)-2]
					}

					removed = append(removed, "`"+version+"`")
				}

				fmt.Println("###", typeName)
				if !ast.IsExported(def.Type) {
					fmt.Println("EV: <b>" + envName(envTypeName) + "</b><br />")
					fmt.Println("Type: `" + typeNameFn(def.Type) + "`<br />")
					if len(versions) > 0 {
						fmt.Println("Available since:", strings.Join(versions, ", "))
					}
					if len(removed) > 0 {
						fmt.Println("Removed in:", strings.Join(removed, ", "))
					}

					fmt.Println()
				}

				fmt.Println(docify(def.Doc))
				fmt.Println()
			}
		} else {
			continue
		}

		if ast.IsExported(def.Type) && !def.Seen {
			def.Seen = true
			listPath(data, def.Type+".", typeName, exclude)
		}
	}
}

// envName sanitizes full type name to env var name
func envName(in string) string {
	// prefix in
	in = "tyk.gw." + in
	// clear underscores
	in = strings.ReplaceAll(in, "_", "")
	// dots to underscores
	in = strings.ReplaceAll(in, ".", "_")
	// to upper
	return strings.ToUpper(in)
}

// typeName strips map/slice details from field type
func typeNameFn(in string) string {
	if !strings.Contains(in, "]") {
		return in
	}

	parts := strings.Split(in, "]")
	if ast.IsExported(parts[1]) {
		return parts[1]
	}

	// map[string]string types are kept
	return in
}

// docify sanitizes godoc with trailing punctuation
func docify(in string) string {
	if in == "" {
		return in
	}

	if strings.Contains(in, "Note:") {
		// This is not great, needs first party markdown in doc comment.
		in = strings.ReplaceAll(in, "Note:", "{{< note success >}}\n**Note**\n\n")
		in += "\n{{< /note >}}"
	}

	lastChar := in[len(in)-1:]
	lastCharRune := []rune(lastChar)

	if unicode.IsPunct(lastCharRune[0]) {
		return in
	}

	if strings.HasSuffix(in, "```") {
		return in
	}

	return in

	// Old docs don't add trailing punctuation:
	// return in + "."
}
