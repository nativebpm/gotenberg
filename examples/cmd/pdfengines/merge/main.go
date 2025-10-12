package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/nativebpm/gotenberg"
)

func main() {
	// Create a Gotenberg client
	client, err := gotenberg.NewClient(&http.Client{}, "http://localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	// Simple HTML content for first PDF
	html1 := `
<!DOCTYPE html>
<html>
<head><title>Page 1</title></head>
<body><h1>First Page</h1><p>This is the first PDF.</p></body>
</html>`

	// Simple HTML content for second PDF
	html2 := `
<!DOCTYPE html>
<html>
<head><title>Page 2</title></head>
<body><h1>Second Page</h1><p>This is the second PDF.</p></body>
</html>`

	// First, convert HTML to PDFs using Chromium
	pdf1Resp, err := client.Chromium().
		ConvertHTML(context.Background(), strings.NewReader(html1)).
		Send()
	if err != nil {
		log.Fatal(err)
	}
	defer pdf1Resp.Body.Close()

	pdf2Resp, err := client.Chromium().
		ConvertHTML(context.Background(), strings.NewReader(html2)).
		Send()
	if err != nil {
		log.Fatal(err)
	}
	defer pdf2Resp.Body.Close()

	// Read PDF data
	pdf1Data := make([]byte, 0)
	buf1 := make([]byte, 1024)
	for {
		n, err := pdf1Resp.Body.Read(buf1)
		if n > 0 {
			pdf1Data = append(pdf1Data, buf1[:n]...)
		}
		if err != nil {
			break
		}
	}

	pdf2Data := make([]byte, 0)
	buf2 := make([]byte, 1024)
	for {
		n, err := pdf2Resp.Body.Read(buf2)
		if n > 0 {
			pdf2Data = append(pdf2Data, buf2[:n]...)
		}
		if err != nil {
			break
		}
	}

	// Merge the PDFs
	mergedResp, err := client.PDFEngines().
		Merge(context.Background()).
		File("pdf1.pdf", strings.NewReader(string(pdf1Data))).
		File("pdf2.pdf", strings.NewReader(string(pdf2Data))).
		Send()

	if err != nil {
		log.Fatal(err)
	}
	defer mergedResp.Body.Close()

	// Save the merged PDF to file
	out, err := os.Create("merged.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err = out.ReadFrom(mergedResp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Merged PDF generated successfully: merged.pdf")
}
