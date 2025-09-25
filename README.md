# Gotenberg Go Client

A comprehensive Go client for [Gotenberg](https://gotenberg.dev/) API with functional options pattern for clean and flexible usage.

## Features

- **Functional Options Pattern**: Clean API without mandatory option structs
- **URL to PDF**: Convert web pages to PDF
- **HTML to PDF**: Convert HTML documents to PDF with additional assets
- **Markdown to PDF**: Convert Markdown files to PDF
- **Webhook Support**: Asynchronous processing with webhooks
- **Comprehensive Options**: Support for all Gotenberg configuration options
- **Type Safety**: Strongly typed options with pointer helpers
- **Easy Configuration**: Predefined paper sizes and helper functions

## Installation

```bash
go get your-module/pkg/gotenberg
```

## Quick Start

### Basic Usage

```go
package main

import (
    "net/http"
    "time"
    "your-module/pkg/gotenberg"
)

func main() {
    // Create client
    client := gotenberg.NewClient(&http.Client{
        Timeout: 30 * time.Second,
    }, "http://localhost:3000")

    // Convert URL to PDF (minimal)
    resp, err := client.ConvertURLToPDF("https://example.com")
    if err != nil {
        panic(err)
    }
    
    // Save PDF
    os.WriteFile("example.pdf", resp.PDF, 0644)
}
```

### URL to PDF with Options

```go
resp, err := client.ConvertURLToPDF("https://example.com",
    gotenberg.WithPaperSize(8.5, 11),           // Letter size
    gotenberg.WithMargins(1, 1, 1, 1),          // 1 inch margins
    gotenberg.WithLandscape(false),             // Portrait
    gotenberg.WithPrintBackground(true),        // Include background
    gotenberg.WithOutputFilename("page.pdf"),   // Custom filename
)
```

### HTML to PDF with Assets

```go
indexHTML := []byte(`<html><body><h1>Hello World</h1></body></html>`)

resp, err := client.ConvertHTMLToPDF(indexHTML,
    gotenberg.A4(),                         // Predefined A4 size
    gotenberg.WithHTMLMargins(0.5, 0.5, 0.5, 0.5),
    gotenberg.WithAdditionalFiles(map[string][]byte{
        "style.css": []byte("body { font-family: Arial; }"),
        "logo.png":  logoBytes,
    }),
    gotenberg.WithHeader([]byte("<html><body>Header</body></html>")),
    gotenberg.WithFooter([]byte("<html><body>Footer</body></html>")),
)
```

### Markdown to PDF

```go
indexHTML := []byte(`<html><body>{{ .markdown }}</body></html>`)
markdownFiles := map[string][]byte{
    "content.md": []byte("# Hello\n\nThis is **markdown**."),
}

resp, err := client.ConvertMarkdownToPDF(indexHTML, markdownFiles,
    gotenberg.WithMarkdownPaperSize(8.27, 11.7), // A4
    gotenberg.WithMarkdownLandscape(false),
    gotenberg.WithMarkdownOutputFilename("doc.pdf"),
)
```

### Webhook (Async Processing)

```go
resp, err := client.ConvertHTMLToPDF(htmlContent,
    gotenberg.WithHTMLWebhook(
        "https://your-app.com/webhook/success",
        "https://your-app.com/webhook/error",
    ),
    gotenberg.WithHTMLWebhookMethods("POST", "POST"),
    gotenberg.WithHTMLWebhookExtraHeaders(map[string]string{
        "Authorization": "Bearer your-token",
        "X-Request-ID":  "req-123",
    }),
)

// resp.PDF will be nil for webhook requests
// PDF will be sent to your webhook URL
```

## API Reference

### Client Creation

```go
func NewClient(httpClient *http.Client, baseURL string) *Client
```

If `httpClient` is nil, `http.DefaultClient` is used.

### Conversion Methods

#### URL to PDF
```go
func (c *Client) ConvertURLToPDF(url string, opts ...URLToPDFOption) (*PDFResponse, error)
```

#### HTML to PDF
```go
func (c *Client) ConvertHTMLToPDF(indexHTML []byte, opts ...ConfigOption) (*PDFResponse, error)
```

#### Markdown to PDF
```go
func (c *Client) ConvertMarkdownToPDF(indexHTML []byte, markdownFiles map[string][]byte, opts ...ConfigOption) (*PDFResponse, error)
```

### Response Structure

```go
type PDFResponse struct {
    PDF                []byte  // PDF content (nil for webhook responses)
    ContentType        string  // Response content type
    ContentLength      int64   // Content length
    ContentDisposition string  // Content disposition header
    Trace              string  // Gotenberg trace ID
}
```

### Error Handling

```go
type GotenbergError struct {
    StatusCode int    // HTTP status code
    Message    string // Error message from Gotenberg
    Trace      string // Gotenberg trace ID
}
```

## Functional Options

### URL to PDF Options

- `WithPaperSize(width, height float64)` - Set custom paper size
- `WithMargins(top, right, bottom, left float64)` - Set margins
- `WithSinglePage(bool)` - Single page mode
- `WithLandscape(bool)` - Landscape orientation
- `WithPrintBackground(bool)` - Include background graphics
- `WithScale(float64)` - Page scale factor
- `WithOutputFilename(string)` - Custom output filename
- `WithWebhook(url, errorURL string)` - Webhook configuration
- `WithWebhookMethods(method, errorMethod string)` - Webhook HTTP methods
- `WithWebhookExtraHeaders(map[string]string)` - Additional webhook headers

### HTML to PDF Options

Same as URL options plus:
- `WithAdditionalFiles(map[string][]byte)` - Additional assets (CSS, images, fonts)
- `WithHeader([]byte)` - Header HTML content
- `WithFooter([]byte)` - Footer HTML content

Use `WithHTML*` prefixed functions for HTML-specific options:
- `WithHTMLPaperSize()`, `WithHTMLMargins()`, etc.

### Markdown to PDF Options

Same as HTML options with `WithMarkdown*` prefixed functions:
- `WithMarkdownPaperSize()`, `WithMarkdownMargins()`, etc.

### Predefined Helpers

#### Paper Sizes
```go
A4()            // A4 for URL conversion
A4()        // A4 for HTML conversion  
A4Markdown()    // A4 for Markdown conversion
Letter()        // Letter for URL conversion
Letter()    // Letter for HTML conversion
LetterMarkdown() // Letter for Markdown conversion
```

#### Paper Size Constants
```go
PaperSizeA4      = [2]float64{8.27, 11.7}  // A4
PaperSizeLetter  = [2]float64{8.5, 11}     // Letter
PaperSizeLegal   = [2]float64{8.5, 14}     // Legal
// ... and more
```

#### Utility Functions
```go
Bool(v bool) *bool           // Create bool pointer
String(v string) *string     // Create string pointer  
Float64(v float64) *float64  // Create float64 pointer
Int(v int) *int              // Create int pointer
```

## Examples

### Complete Example with Error Handling

```go
package main

import (
    "fmt"
    "net/http"
    "os"
    "time"
    "your-module/pkg/gotenberg"
)

func main() {
    client := gotenberg.NewClient(&http.Client{
        Timeout: 60 * time.Second,
    }, "http://localhost:3000")

    html := []byte(`
        <!DOCTYPE html>
        <html>
        <head>
            <meta charset="UTF-8">
            <title>Invoice</title>
            <link rel="stylesheet" href="style.css">
        </head>
        <body>
            <h1>Invoice #12345</h1>
            <p>Thank you for your business!</p>
        </body>
        </html>
    `)

    css := []byte(`
        body { 
            font-family: Arial, sans-serif; 
            margin: 40px;
        }
        h1 { 
            color: #333; 
            border-bottom: 2px solid #007cba;
        }
    `)

    resp, err := client.ConvertHTMLToPDF(html,
        gotenberg.A4(),
        gotenberg.WithHTMLMargins(1, 1, 1, 1),
        gotenberg.WithHTMLPrintBackground(true),
        gotenberg.WithHTMLOutputFilename("invoice-12345.pdf"),
        gotenberg.WithAdditionalFiles(map[string][]byte{
            "style.css": css,
        }),
    )

    if err != nil {
        if gotErr, ok := err.(*gotenberg.GotenbergError); ok {
            fmt.Printf("Gotenberg error %d: %s (trace: %s)\n", 
                gotErr.StatusCode, gotErr.Message, gotErr.Trace)
        } else {
            fmt.Printf("Error: %v\n", err)
        }
        return
    }

    if err := os.WriteFile("invoice.pdf", resp.PDF, 0644); err != nil {
        fmt.Printf("Failed to save PDF: %v\n", err)
        return
    }

    fmt.Printf("PDF created successfully! Size: %d bytes, Trace: %s\n", 
        len(resp.PDF), resp.Trace)
}
```

## Testing

Run the tests:

```bash
go test ./pkg/gotenberg/...
```

The package includes comprehensive tests covering all functionality with mock Gotenberg server responses.

## Requirements

- Go 1.19+
- Running Gotenberg instance (see [Gotenberg installation](https://gotenberg.dev/docs/getting-started/installation))

## License

This project is licensed under the MIT License.