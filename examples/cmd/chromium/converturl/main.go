package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/nativebpm/connectors/gotenberg"
)

func main() {
	client, err := gotenberg.NewClient(http.Client{}, "http://localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	// Convert web page from URL to PDF
	resp, err := client.
		Chromium().
		ConvertURL(context.Background(), "https://gotenberg.dev/docs/getting-started/introduction").
		PaperSizeA4().
		Margins(1, 1, 1, 1).
		PrintBackground().
		Timeout(30 * time.Second).
		Send()
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	f, err := os.Create("converturl-chromium.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("PDF generated from URL: converturl-chromium.pdf")
}
