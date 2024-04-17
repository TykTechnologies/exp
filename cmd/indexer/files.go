package main

import "embed"

// This file is local and should not be copied without modification.

var (
	//go:embed index.tpl
	embeddedFiles embed.FS
)
