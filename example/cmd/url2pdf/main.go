package main

import (
	"context"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/nativebpm/gotenberg-client"
)

func main() {
	gotenbergURL := `http://localhost:3000`

	httpClient := &http.Client{
		Timeout: 90 * time.Second,
	}

	client := gotenberg.NewClient(httpClient, gotenbergURL)

	resp, err := client.ConvertURLToPDF(context.Background(), "https://example.com",
		gotenberg.WithPrintBackground(true),
		gotenberg.WithOutputFilename("example.pdf"),
		gotenberg.WithPaperSizeLetter(),
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

	if err := os.WriteFile("./example.pdf", pdfData, 0644); err != nil {
		log.Fatal(err)
	}

	slog.Info("URL to PDF conversion completed",
		"pdf_size", resp.ContentLength,
		"trace", resp.Header.Get("Gotenberg-Trace"))
}
