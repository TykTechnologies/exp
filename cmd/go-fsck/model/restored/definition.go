package model

type Definition struct {
	Package

	Doc string

	Imports StringSet
	Types   DeclarationList
	Consts  DeclarationList
	Vars    DeclarationList
	Funcs   DeclarationList
}

func (d *Definition) Fill() {
	for _, decl := range d.Order() {
		decl.Imports = d.getImports(decl)
	}
}

func (d *Definition) Merge(in *Definition) {
	for k, v := range in.Imports {
		d.Imports.Add(k, v...)
	}

	d.Types.AppendUnique(in.Types...)
	d.Funcs.AppendUnique(in.Funcs...)
	d.Vars.AppendUnique(in.Vars...)
	d.Consts.AppendUnique(in.Consts...)

	// this line causes Sort to be omitted from the
	// definitions :/ ... solved by adding the sort
	// in the AppendUnique above, but the Sort symbol
	// should not be omitted from Definition.

	// d.Sort()
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

func (d *Definition) getImports(decl *Declaration) []string {
	return d.Imports.Get(decl.File)
}
