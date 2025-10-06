package invoice

import (
	"embed"
	"html/template"
)

//go:embed template.html
var fs embed.FS
var Template = template.Must(template.ParseFS(fs, "template.html"))
