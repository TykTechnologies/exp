package model


// ExtractOptions contains options for extraction
type ExtractOptions struct {
	IncludeFunctions  bool
	IncludeTests      bool
	IncludeUnexported bool
	IgnoreFiles       []string
	IncludeInternal   bool
}
