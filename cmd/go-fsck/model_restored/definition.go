package model

type (
	Definition struct {
		Package string

		Imports Imports
		Types   DeclarationList
		Consts  DeclarationList
		Vars    DeclarationList
		Funcs   DeclarationList
	}
)

func (d *Definition) getImports(decl *Declaration) []string {
	return d.Imports.Get(decl.File)
}

func (d *Definition) Fill() {
	all := append([]*Declaration{}, d.Types...)
	all = append(all, d.Funcs...)
	all = append(all, d.Vars...)
	all = append(all, d.Consts...)

	for _, decl := range all {
		decl.Imports = d.getImports(decl)
	}
}
