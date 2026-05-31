package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/nativebpm/gotenberg"
)

func main() {
	// Start a mock HTML server that responds slowly
	go startMockHTMLServer()

	// Wait a bit for server to start
	time.Sleep(100 * time.Millisecond)

	// Create client pointing to real Gotenberg server
	client, err := gotenberg.NewClient(http.Client{}, "http://localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Making request to Gotenberg server with 10 second timeout...")
	log.Println("Gotenberg will try to fetch HTML from slow server at http://localhost:3001/html")

	// Convert web page from URL to PDF with 10 second timeout
	// The HTML server responds slowly, causing timeout
	start := time.Now()
	resp, err := client.Chromium().
		ConvertURL(context.Background(), "http://host.docker.internal:3001/html").
		Timeout(10 * time.Second).
		Trace("timeout-example-request-1").
		Send()

	elapsed := time.Since(start)

	if err != nil {
		log.Printf("Request failed after %v: %v", elapsed, err)
		log.Println("This demonstrates timeout functionality - Gotenberg timed out waiting for HTML content")
		return
	}

	defer resp.Body.Close()

	file, err := os.Create("timeout-example.pdf")
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer file.Close()

	n, err := file.ReadFrom(resp.Body)
	if err != nil {
		log.Fatalf("Failed to write PDF: %v", err)
	}

	log.Printf("Unexpected success after %v, received %d bytes", elapsed, n)
}

func startMockHTMLServer() {
	http.HandleFunc("/html", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Mock HTML server received request, simulating slow response...")

		// Simulate slow HTML server response (15 seconds)
		time.Sleep(15 * time.Second)

		// Return a simple HTML page
		htmlContent := `<!DOCTYPE html>
<html>
<head>
    <title>Test Page</title>
</head>
<body>
    <h1>Hello World</h1>
    <p>This is a test page that loads slowly.</p>
</body>
</html>`

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(htmlContent))

		log.Println("Mock HTML server sent response after 15 seconds")
	})

	log.Println("Starting mock HTML server on :3001")
	log.Fatal(http.ListenAndServe(":3001", nil))
}
