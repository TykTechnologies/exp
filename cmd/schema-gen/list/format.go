package list

import (
	"github.com/stoewer/go-strcase"
)

var format = struct {
	SnakeCase func(string) string
}{
	SnakeCase: strcase.SnakeCase,
}
