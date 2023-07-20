package model

type (
	Definition struct {
		Package string

		Imports []string
		Types   DeclarationList
		Consts  DeclarationList
		Vars    DeclarationList
		Funcs   DeclarationList
	}
)
