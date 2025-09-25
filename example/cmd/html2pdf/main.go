package main

import (
	"bytes"
	"context"
	"io"
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

	resp, err := client.ConvertHTMLToPDF(context.Background(), html.Bytes(),
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

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("unexpected status code: %d", resp.StatusCode)
	}

	pdfData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("./invoice_new.pdf", pdfData, 0644); err != nil {
		log.Fatal(err)
	}

	slog.Info("HTML to PDF conversion completed",
		"pdf_size", resp.ContentLength,
		"content_type", resp.Header.Get("Content-Type"),
		"trace", resp.Header.Get("Gotenberg-Trace"))
}
