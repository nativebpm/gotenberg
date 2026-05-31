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
	client, err := gotenberg.NewClient(http.Client{}, "http://localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	// Simple HTML content that will be treated as a document
	htmlContent := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>Sample Document</title>
</head>
<body>
    <h1>Sample Office Document</h1>
    <p>This is a simple HTML document that will be converted to PDF using LibreOffice.</p>
    <ul>
        <li>Item 1</li>
        <li>Item 2</li>
        <li>Item 3</li>
    </ul>
</body>
</html>`

	// Convert HTML to PDF using LibreOffice
	resp, err := client.LibreOffice().
		Convert(context.Background()).
		File("document.html", strings.NewReader(htmlContent)).
		Landscape(true).
		Send()

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Save the PDF to file
	out, err := os.Create("output.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err = out.ReadFrom(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PDF generated successfully: output.pdf")
}
