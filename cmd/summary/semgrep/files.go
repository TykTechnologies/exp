package semgrep

import "embed"

// This file is local and should not be copied without modification.

var (
	//go:embed report*.tpl
	files embed.FS
)
