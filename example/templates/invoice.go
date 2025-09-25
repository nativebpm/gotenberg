package templates

import (
	"embed"
	"html/template"
)

//go:embed invoice.html
var invoiceTemplatesFS embed.FS

var InvoiceTemplate = template.Must(template.ParseFS(invoiceTemplatesFS, "invoice.html"))
