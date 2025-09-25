package main

import (
	"bytes"
	"log"
	"log/slog"
	"net/http"
	"os"
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
		gotenberg.WithLandscape(false),
		gotenberg.WithScale(1.0),
		gotenberg.WithOutputFilename("invoice.pdf"),
		gotenberg.WithPaperSizeA4(),
		gotenberg.WithMargins(1.0, 1.0, 1.0, 1.0),
		gotenberg.WithHTMLAdditionalFiles(files),
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("./invoice_new.pdf", resp.PDF, 0644); err != nil {
		log.Fatal(err)
	}

	slog.Info("HTML to PDF conversion completed",
		"pdf_size", len(resp.PDF),
		"content_type", resp.ContentType,
		"trace", resp.Trace)
}
