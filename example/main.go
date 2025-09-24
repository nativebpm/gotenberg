package main

import (
	"bytes"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/nativebpm/gotenberg-client"
	"github.com/nativebpm/gotenberg-client/model"
	"github.com/nativebpm/gotenberg-client/templates"
)

func main() {
	// Example of using new gotenberg package with functional options

	gotenbergURL := `http://localhost:3000` // or os.Getenv("GOTENBERG_URL")

	// Create HTTP client
	hc := &http.Client{
		Timeout: 90 * time.Second,
	}

	// Create Gotenberg client
	client := gotenberg.NewClient(hc, gotenbergURL)

	data := model.DemoData

	// Generate HTML from template
	htmlDoc := makeHtmlDemo(data)

	// Example 1: Convert HTML to PDF (new functional options approach)
	err := convertHTMLToPDFExample(client, htmlDoc)
	if err != nil {
		slog.Error("Failed to convert HTML to PDF", "error", err)
		return
	}

	// Example 2: Convert URL to PDF with minimal options
	err = convertURLToPDFExample(client)
	if err != nil {
		slog.Error("Failed to convert URL to PDF", "error", err)
		return
	}

	// Example 3: Convert with webhook (async)
	err = convertHTMLToPDFWithWebhookExample(client, htmlDoc)
	if err != nil {
		slog.Error("Failed to convert HTML to PDF with webhook", "error", err)
		return
	}

	// Example 4: Minimal conversion without any options
	err = convertHTMLMinimalExample(client, htmlDoc)
	if err != nil {
		slog.Error("Failed minimal HTML conversion", "error", err)
		return
	}

	slog.Info("All conversions completed successfully")
}

// convertHTMLToPDFExample demonstrates HTML to PDF conversion with options
func convertHTMLToPDFExample(client *gotenberg.Client, htmlDoc *bytes.Buffer) error {
	slog.Info("Converting HTML to PDF with options...")

	resp, err := client.ConvertHTMLToPDF(htmlDoc.Bytes(),
		gotenberg.WithHTMLPrintBackground(true),
		gotenberg.WithHTMLLandscape(false),
		gotenberg.WithHTMLScale(1.0),
		gotenberg.WithHTMLOutputFilename("invoice.pdf"),
		gotenberg.A4HTML(),                            // Use predefined A4 paper size
		gotenberg.WithHTMLMargins(1.0, 1.0, 1.0, 1.0), // 1 inch margins
	)
	if err != nil {
		return err
	}

	// Save PDF
	err = os.WriteFile("./invoice_new.pdf", resp.PDF, 0644)
	if err != nil {
		return err
	}

	slog.Info("HTML to PDF conversion completed",
		"pdf_size", len(resp.PDF),
		"content_type", resp.ContentType,
		"trace", resp.Trace)

	return nil
}

// convertURLToPDFExample demonstrates URL to PDF conversion
func convertURLToPDFExample(client *gotenberg.Client) error {
	slog.Info("Converting URL to PDF...")

	resp, err := client.ConvertURLToPDF("https://example.com",
		gotenberg.WithPrintBackground(true),
		gotenberg.WithOutputFilename("example.pdf"),
		gotenberg.Letter(), // Use predefined Letter paper size
	)
	if err != nil {
		return err
	}

	// Save PDF
	err = os.WriteFile("./example.pdf", resp.PDF, 0644)
	if err != nil {
		return err
	}

	slog.Info("URL to PDF conversion completed",
		"pdf_size", len(resp.PDF),
		"trace", resp.Trace)

	return nil
}

// convertHTMLToPDFWithWebhookExample demonstrates async conversion with webhook
func convertHTMLToPDFWithWebhookExample(client *gotenberg.Client, htmlDoc *bytes.Buffer) error {
	slog.Info("Converting HTML to PDF with webhook (async)...")

	resp, err := client.ConvertHTMLToPDF(htmlDoc.Bytes(),
		gotenberg.WithHTMLPrintBackground(true),
		gotenberg.WithHTMLOutputFilename("invoice_async.pdf"),
		gotenberg.WithHTMLWebhook(
			"https://your-webhook-url.com/success",
			"https://your-webhook-url.com/error",
		),
		gotenberg.WithHTMLWebhookMethods("POST", "POST"),
		gotenberg.WithHTMLWebhookExtraHeaders(map[string]string{
			"Authorization":   "Bearer your-token",
			"X-Custom-Header": "custom-value",
		}),
	)
	if err != nil {
		return err
	}

	// When using webhook, PDF is not returned immediately (resp.PDF will be nil)
	slog.Info("Async HTML to PDF conversion started",
		"trace", resp.Trace,
		"pdf_returned", resp.PDF != nil)

	return nil
}

// convertHTMLMinimalExample demonstrates minimal conversion without any options
func convertHTMLMinimalExample(client *gotenberg.Client, htmlDoc *bytes.Buffer) error {
	slog.Info("Converting HTML to PDF (minimal, no options)...")

	// This is the beauty of functional options - you can call without any options!
	resp, err := client.ConvertHTMLToPDF(htmlDoc.Bytes())
	if err != nil {
		return err
	}

	// Save PDF
	err = os.WriteFile("./invoice_minimal.pdf", resp.PDF, 0644)
	if err != nil {
		return err
	}

	slog.Info("Minimal HTML to PDF conversion completed",
		"pdf_size", len(resp.PDF),
		"trace", resp.Trace)

	return nil
}

func makeHtmlDemo(data model.InvoiceData) *bytes.Buffer {
	buf := bytes.NewBuffer(nil)
	templates.Demo.Execute(buf, data)
	return buf
}
