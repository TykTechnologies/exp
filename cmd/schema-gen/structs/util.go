package structs

import (
	"strings"
)

func TrimSpace(in interface{ Text() string }) string {
	return strings.TrimSpace(in.Text())
}
