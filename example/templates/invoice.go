package templates

import (
	"embed"
	"html/template"
)

//go:embed invoice.html
var templatesFS embed.FS

var InvoiceTemplate = template.Must(template.ParseFS(templatesFS, "invoice.html"))
