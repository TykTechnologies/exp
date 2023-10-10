package model

import (
	"fmt"
	"go/ast"
	"path"
	"sort"
	"strings"

	"golang.org/x/exp/slices"
)

type (
	Definition struct {
		Package string
		Doc     StringSet
		Imports StringSet
		Types   DeclarationList
		Consts  DeclarationList
		Vars    DeclarationList
		Funcs   DeclarationList
	}
)

type StringSet map[string][]string

type (
	DeclarationKind string

	Declaration struct {
		Kind DeclarationKind
		File string

		SelfContained bool

		Imports    []string            `json:",omitempty"`
		References map[string][]string `json:",omitempty"`

		Name     string   `json:",omitempty"`
		Names    []string `json:",omitempty"`
		Receiver string   `json:",omitempty"`

		Arguments []string `json:",omitempty"`
		Returns   []string `json:",omitempty"`

		Signature string `json:",omitempty"`
		Source    string
	}
)

const (
	StructKind  DeclarationKind = "struct"
	ImportKind                  = "import"
	ConstKind                   = "const"
	TypeKind                    = "type"
	FuncKind                    = "func"
	VarKind                     = "var"
	CommentKind                 = "comment"
)

func (d DeclarationKind) String() string {
	return string(d)
}

func (d *Definition) getImports(decl *Declaration) []string {
	return d.Imports.Get(decl.File)
}

func (d *Definition) Order() []*Declaration {
	count := len(d.Types) + len(d.Funcs) + len(d.Vars) + len(d.Consts)
	result := make([]*Declaration, 0, count)

	result = append(result, d.Types...)
	result = append(result, d.Funcs...)
	result = append(result, d.Vars...)
	result = append(result, d.Consts...)
	return result
}

func (d *Definition) Sort() {
	d.Types.Sort()
	d.Vars.Sort()
	d.Consts.Sort()
	d.Funcs.Sort()
}

func (d *Definition) Fill() {
	for _, decl := range d.Order() {
		decl.Imports = d.getImports(decl)
	}
}

func (d *Declaration) Keys() []string {
	trimPath := "*."
	if d.Name != "" {
		return []string{
			strings.Trim(d.Receiver+"."+d.Name, trimPath),
		}
	}
	if len(d.Names) != 0 {
		result := make([]string, len(d.Names))
		for k, v := range d.Names {
			result[k] = strings.Trim(d.Receiver+"."+v, trimPath)
		}
	}
	return nil
}

func (i *StringSet) Add(key, lit string) {
	data := *i
	if data == nil {
		data = make(StringSet)
	}
	if set, ok := data[key]; ok {
		if slices.Contains(set, lit) {
			return
		}
		data[key] = append(set, lit)
		return
	}
	data[key] = []string{lit}
	*i = data
}

func (i StringSet) Get(key string) []string {
	val, _ := i[key]
	if val != nil {
		sort.Strings(val)
	}
	return val
}

func (i StringSet) All() []string {
	result := []string{}
	for _, set := range i {
		result = append(result, set...)
	}
	return result
}

// Map returns a map with the short package name as the key
// and the full import path as the value.
func (i StringSet) Map() map[string]string {
	result := map[string]string{}
	imports := i.All()

	for _, imported := range imports {
		var short, long string

		// aliased package
		if strings.Contains(imported, " ") {
			line := strings.Split(imported, " ")
			short, long = line[0], strings.Trim(line[1], `"`)
		} else {
			long = strings.Trim(imported, `"`)
			short = path.Base(long)
		}

		val, ok := result[short]
		if ok && val != long {
			fmt.Printf("WARN: Import path conflict: %s\n%s (prev) != %s (new)\n", short, val, long)
		}

		result[short] = long
	}

	return result
}

type DeclarationList []*Declaration

func (p *DeclarationList) Append(in ...*Declaration) {
	*p = append(*p, in...)
}

func (p DeclarationList) FindKind(kind DeclarationKind) (result []*Declaration) {
	for _, decl := range p {
		if decl.Kind == kind {
			result = append(result, decl)
		}
	}
	return
}

func (p *DeclarationList) Sort() {
	sort.Slice(*p, func(i, j int) bool {
		a, b := (*p)[i], (*p)[j]
		if a.Kind != b.Kind {
			indexOf := map[DeclarationKind]int{
				CommentKind: 0,
				ImportKind:  1,
				ConstKind:   2,
				StructKind:  3,
				TypeKind:    4,
				VarKind:     5,
				FuncKind:    6,
			}
			return indexOf[a.Kind] < indexOf[b.Kind]
		}
		ae, be := ast.IsExported(a.Name), ast.IsExported(b.Name)
		if ae != be {
			return ae
		}

		if a.Receiver != b.Receiver {
			if a.Receiver == "" {
				return true
			}
			return a.Receiver < b.Receiver
		}

		if a.Signature != b.Signature {
			return a.Signature < b.Signature
		}

		return a.Name < b.Name
	})
}
