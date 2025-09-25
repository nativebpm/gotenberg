package main

import (
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

	resp, err := client.ConvertURLToPDF("https://example.com",
		gotenberg.WithPrintBackground(true),
		gotenberg.WithOutputFilename("example.pdf"),
		gotenberg.WithPaperSizeLetter(),
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("./example.pdf", resp.PDF, 0644); err != nil {
		log.Fatal(err)
	}

	slog.Info("URL to PDF conversion completed",
		"pdf_size", len(resp.PDF),
		"trace", resp.Trace)
}
