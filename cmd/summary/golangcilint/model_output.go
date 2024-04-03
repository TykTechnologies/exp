package golangcilint

type Summary struct {
	Files []*File
}

type File struct {
	Name   string
	Errors []*Error
}

type Error struct {
	FromLinter string
	Text       []string
	Count      int
}
