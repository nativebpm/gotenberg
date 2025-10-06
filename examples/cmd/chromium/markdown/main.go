package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/nativebpm/connectors/gotenberg"
	"github.com/nativebpm/connectors/gotenberg/examples/pkg/templates/markdown"
)

func main() {
	httpClient := http.Client{
		Timeout: 30 * time.Second,
	}

	client, err := gotenberg.NewClient(httpClient, "http://localhost:3000")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	indexHTML, err := markdown.FS.ReadFile("template.html")
	if err != nil {
		log.Fatalf("Failed to read template.html: %v", err)
	}

	markdownContent, err := markdown.FS.ReadFile("content.md")
	if err != nil {
		log.Fatalf("Failed to read content.md: %v", err)
	}

	ctx := context.Background()

	response, err := client.Chromium().
		ConvertMarkdown(ctx, bytes.NewReader(indexHTML)).
		File("content.md", bytes.NewReader(markdownContent)).
		PaperSizeA4().
		Landscape().
		Margins(1, 1, 1, 1).
		OutputFilename("markdown-example.pdf").
		Send()

	if err != nil {
		log.Fatalf("Failed to convert markdown: %v", err)
	}
	defer response.Body.Close()

	file, err := os.Create("markdown-example.pdf")
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer file.Close()

	_, err = file.ReadFrom(response.Body)
	if err != nil {
		log.Fatalf("Failed to write PDF: %v", err)
	}

	fmt.Printf("Markdown converted to PDF successfully!\n")
	fmt.Printf("Output file: markdown-example.pdf\n")
	fmt.Printf("Gotenberg trace: %s\n", response.GotenbergTrace)
}
