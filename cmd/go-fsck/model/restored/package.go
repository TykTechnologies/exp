package model

type (
	Package struct {
		// Package is the name of the package.
		Package string
		// ImportPath contains the import path (github...).
		ImportPath string
		// Path is sanitized to contain the relative location (folder).
		Path string
		// TestPackage is true if this is a test package.
		TestPackage bool
	}

	Definition struct {
		Package

		Doc     StringSet
		Imports StringSet
		Types   DeclarationList
		Consts  DeclarationList
		Vars    DeclarationList
		Funcs   DeclarationList
	}
)

func (d *Definition) ClearSource() {
	d.Types.ClearSource()
	d.Vars.ClearSource()
	d.Consts.ClearSource()
	d.Funcs.ClearSource()
}

func (d *Definition) Fill() {
	for _, decl := range d.Order() {
		decl.Imports = d.getImports(decl)
	}
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

func (p Package) Name() string {
	return p.Package
}

func (d *Definition) getImports(decl *Declaration) []string {
	return d.Imports.Get(decl.File)
}
