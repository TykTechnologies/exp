package stats

type (
	// SymbolReference is a O(1) structure that combines
	// the package being imported, the symbol being used and
	// which function is referencing that symbol.
	SymbolReference struct {
		Package string
		Import  string
		Symbol  string
		UsedBy  string
	}
)
