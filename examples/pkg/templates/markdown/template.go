package markdown

import (
	"embed"
)

//go:embed template.html content.md
var FS embed.FS
