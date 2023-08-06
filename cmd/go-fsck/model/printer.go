package model

import (
	"go/ast"
	"go/printer"
	"go/token"
	"io"
)

func PrintSource(out io.Writer, fset *token.FileSet, file *ast.File, node any) error {
	val := &printer.CommentedNode{
		Node:     node,
		Comments: file.Comments,
	}
	return printer.Fprint(out, fset, val)
}
