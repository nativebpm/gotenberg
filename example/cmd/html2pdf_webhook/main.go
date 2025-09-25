package main

import (
	"bytes"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/nativebpm/gotenberg-client"
	"github.com/nativebpm/gotenberg-client/example/model"
	"github.com/nativebpm/gotenberg-client/example/pkg/image"
	"github.com/nativebpm/gotenberg-client/example/pkg/templates/invoice"
)

func main() {
	gotenbergURL := `http://localhost:3000`

	httpClient := &http.Client{
		Timeout: 90 * time.Second,
	}

	client := gotenberg.NewClient(httpClient, gotenbergURL)

	data := model.InvoiceData
	html := bytes.NewBuffer(nil)
	invoice.Template.Execute(html, data)

	logo := image.LogoPNG()
	files := map[string][]byte{"logo.png": logo}

	resp, err := client.ConvertHTMLToPDF(html.Bytes(),
		gotenberg.WithPrintBackground(true),
		gotenberg.WithOutputFilename("invoice_async.pdf"),
		gotenberg.WithWebhook(
			"https://your-webhook-url.com/success",
			"https://your-webhook-url.com/error",
		),
		gotenberg.WithWebhookMethods("POST", "POST"),
		gotenberg.WithWebhookExtraHeaders(map[string]string{
			"Authorization":   "Bearer your-token",
			"X-Custom-Header": "custom-value",
		}),
		gotenberg.WithHTMLAdditionalFiles(files),
	)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Async HTML to PDF conversion started",
		"trace", resp.Trace,
		"pdf_returned", resp.PDF != nil)
}
