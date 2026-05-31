package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/nativebpm/gotenberg"
)

func main() {
	client, err := gotenberg.NewClient(http.Client{}, "http://localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	html := `<html><body><h1>Request Tracing Example</h1><p>This PDF was generated with a custom trace ID.</p></body></html>`

	// Set a custom trace ID for this request
	resp, err := client.Chromium().
		ConvertHTML(context.Background(), strings.NewReader(html)).
		Trace("my-custom-trace-12345").
		PaperSizeA6().
		Margins(0.5, 0.5, 0.5, 0.5).
		Send()
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	f, _ := os.Create("trace-example.pdf")
	io.Copy(f, resp.Body)
	f.Close()

	log.Printf("PDF generated with trace ID: %s", resp.GotenbergTrace)
}
