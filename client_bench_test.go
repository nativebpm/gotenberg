package gotenberg

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// mockGotenbergServerForBench creates a lightweight mock server for benchmarking
func mockGotenbergServerForBench() *httptest.Server {
	// Minimal PDF response for benchmarking
	pdfContent := []byte("%PDF-1.4\n1 0 obj\n<< /Type /Catalog /Pages 2 0 R >>\nendobj\n2 0 obj\n<< /Type /Pages /Kids [3 0 R] /Count 1 >>\nendobj\n3 0 obj\n<< /Type /Page /Parent 2 0 R /Resources << /Font << /F1 4 0 R >> >> /MediaBox [0 0 612 792] /Contents 5 0 R >>\nendobj\n4 0 obj\n<< /Type /Font /Subtype /Type1 /BaseFont /Times-Roman >>\nendobj\n5 0 obj\n<< /Length 44 >>\nstream\nBT\n/F1 12 Tf\n72 720 Td\n(Hello World) Tj\nET\nendstream\nendobj\nxref\n0 6\n0000000000 65535 f \n0000000010 00000 n \n0000000079 00000 n \n0000000173 00000 n \n0000000301 00000 n \n0000000380 00000 n \ntrailer\n<< /Size 6 /Root 1 0 R >>\nstartxref\n492\n%%EOF")

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Consume request body to avoid connection issues
		io.Copy(io.Discard, r.Body)
		r.Body.Close()

		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=test.pdf")
		w.Header().Set("Gotenberg-Trace", "bench-trace-id")
		w.WriteHeader(http.StatusOK)
		w.Write(pdfContent)
	}))
}

// BenchmarkNewClient measures client creation overhead.
// Important if clients are created frequently instead of being reused.
// Should show minimal allocations since Client is a simple struct.
func BenchmarkNewClient(b *testing.B) {
	httpClient := &http.Client{}
	baseURL := "http://localhost:3000"

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = NewClient(httpClient, baseURL)
	}
}

// BenchmarkConvertURLToPDF_NoOptions measures URL to PDF conversion without any options.
// This is the baseline performance with minimal allocations from multipart form creation.
func BenchmarkConvertURLToPDF_NoOptions(b *testing.B) {
	server := mockGotenbergServerForBench()
	defer server.Close()

	client := NewClient(nil, server.URL)
	url := "https://example.com"

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		resp, err := client.ConvertURLToPDF(url)
		if err != nil {
			b.Fatalf("ConvertURLToPDF failed: %v", err)
		}
		_ = resp
	}
}

// BenchmarkConvertURLToPDF_WithOptions measures URL to PDF conversion with common options.
// Tests functional options pattern overhead and multipart form creation efficiency.
func BenchmarkConvertURLToPDF_WithOptions(b *testing.B) {
	server := mockGotenbergServerForBench()
	defer server.Close()

	client := NewClient(nil, server.URL)
	url := "https://example.com"

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		resp, err := client.ConvertURLToPDF(url,
			WithPaperSize(8.5, 11),
			WithMargins(1, 1, 1, 1),
			WithLandscape(false),
			WithPrintBackground(true),
			WithOutputFilename("test.pdf"),
		)
		if err != nil {
			b.Fatalf("ConvertURLToPDF failed: %v", err)
		}
		_ = resp
	}
}

// BenchmarkConvertURLToPDF_ManyOptions measures URL to PDF conversion with maximum options.
// Tests worst-case scenario with webhooks, headers, and all configuration options.
func BenchmarkConvertURLToPDF_ManyOptions(b *testing.B) {
	server := mockGotenbergServerForBench()
	defer server.Close()

	client := NewClient(nil, server.URL)
	url := "https://example.com"

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		resp, err := client.ConvertURLToPDF(url,
			WithPaperSize(8.5, 11),
			WithMargins(1, 1, 1, 1),
			WithSinglePage(false),
			WithLandscape(false),
			WithPrintBackground(true),
			WithScale(1.0),
			WithOutputFilename("test.pdf"),
			WithWebhook("https://webhook.example.com", "https://webhook-error.example.com"),
			WithWebhookMethods("POST", "PUT"),
			WithWebhookExtraHeaders(map[string]string{
				"Authorization":   "Bearer token",
				"X-Custom-Header": "custom-value",
			}),
		)
		if err != nil {
			b.Fatalf("ConvertURLToPDF failed: %v", err)
		}
		_ = resp
	}
}

// BenchmarkConvertHTMLToPDF_NoOptions measures HTML to PDF conversion without options.
// Includes file upload overhead but minimal configuration.
func BenchmarkConvertHTMLToPDF_NoOptions(b *testing.B) {
	server := mockGotenbergServerForBench()
	defer server.Close()

	client := NewClient(nil, server.URL)
	html := []byte("<html><body><h1>Test</h1></body></html>")

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		resp, err := client.ConvertHTMLToPDF(html)
		if err != nil {
			b.Fatalf("ConvertHTMLToPDF failed: %v", err)
		}
		_ = resp
	}
}

// BenchmarkConvertHTMLToPDF_WithOptions measures HTML to PDF conversion with configuration.
// Tests combination of file uploads and functional options with optimized formatters.
func BenchmarkConvertHTMLToPDF_WithOptions(b *testing.B) {
	server := mockGotenbergServerForBench()
	defer server.Close()

	client := NewClient(nil, server.URL)
	html := []byte("<html><body><h1>Test</h1></body></html>")

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		resp, err := client.ConvertHTMLToPDF(html,
			WithPaperSize(8.5, 11),
			WithMargins(1, 1, 1, 1),
			WithLandscape(false),
			WithPrintBackground(true),
			WithOutputFilename("test.pdf"),
		)
		if err != nil {
			b.Fatalf("ConvertHTMLToPDF failed: %v", err)
		}
		_ = resp
	}
}

// BenchmarkConvertHTMLToPDF_WithFiles measures HTML conversion with multiple additional files.
// Tests multipart form efficiency with CSS, images, headers, and footers.
func BenchmarkConvertHTMLToPDF_WithFiles(b *testing.B) {
	server := mockGotenbergServerForBench()
	defer server.Close()

	client := NewClient(nil, server.URL)
	html := []byte("<html><head><link rel='stylesheet' href='style.css'></head><body><h1>Test</h1></body></html>")
	css := []byte("body { font-family: Arial; margin: 20px; }")
	image := []byte("fake-image-data-for-benchmark")

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		resp, err := client.ConvertHTMLToPDF(html,
			WithHTMLAdditionalFiles(map[string][]byte{
				"style.css": css,
				"logo.png":  image,
			}),
			WithHTMLHeader([]byte("<html><body>Header</body></html>")),
			WithHTMLFooter([]byte("<html><body>Footer</body></html>")),
		)
		if err != nil {
			b.Fatalf("ConvertHTMLToPDF failed: %v", err)
		}
		_ = resp
	}
}

// BenchmarkConvertMarkdownToPDF_NoOptions measures Markdown to PDF conversion baseline.
// Tests template HTML + Markdown files upload performance.
func BenchmarkConvertMarkdownToPDF_NoOptions(b *testing.B) {
	server := mockGotenbergServerForBench()
	defer server.Close()

	client := NewClient(nil, server.URL)
	html := []byte("<html><body>{{ .markdown }}</body></html>")
	markdown := map[string][]byte{
		"test.md": []byte("# Test\n\nThis is a **test** markdown."),
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		resp, err := client.ConvertMarkdownToPDF(html, markdown)
		if err != nil {
			b.Fatalf("ConvertMarkdownToPDF failed: %v", err)
		}
		_ = resp
	}
}

// BenchmarkOptionCreation measures functional option creation cost.
// Critical for understanding closure creation and heap allocation overhead.
func BenchmarkOptionCreation(b *testing.B) {
	b.Run("WithPaperSize", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = WithPaperSize(8.5, 11) // Allocation behavior depends on WithPaperSize implementation
		}
	})

	b.Run("WithWebhookExtraHeaders", func(b *testing.B) {
		headers := map[string]string{
			"Authorization":   "Bearer token",
			"X-Custom-Header": "custom-value",
		}
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = WithWebhookExtraHeaders(headers) // May be expensive due to map copying
		}
	})
}

// BenchmarkConfigApplication measures the cost of applying multiple options to config.
// Tests functional option pattern performance with realistic option combinations.
func BenchmarkConfigApplication(b *testing.B) {
	options := []ConvOption{
		WithPaperSize(8.5, 11),
		WithMargins(1, 1, 1, 1),
		WithSinglePage(false),
		WithLandscape(false),
		WithPrintBackground(true),
		WithScale(1.0),
		WithOutputFilename("test.pdf"),
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		config := &convConfig{}
		for _, opt := range options {
			opt(config)
		}
		_ = config
	}
}

// BenchmarkMultipartFormCreation measures raw multipart form creation performance.
// Uses standard bytes.Buffer allocation for each iteration to measure baseline overhead.
func BenchmarkMultipartFormCreation(b *testing.B) {
	url := "https://example.com"

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)

		// Simulate adding URL field
		writer.WriteField("url", url)

		// Simulate adding some form fields
		writer.WriteField("paperWidth", "8.5")
		writer.WriteField("paperHeight", "11")
		writer.WriteField("marginTop", "1")
		writer.WriteField("marginBottom", "1")
		writer.WriteField("marginLeft", "1")
		writer.WriteField("marginRight", "1")
		writer.WriteField("landscape", "false")
		writer.WriteField("printBackground", "true")

		writer.Close()
		_ = buf.Bytes()
	}
}

// BenchmarkLargePayload measures performance with large HTML content.
// Tests memory allocation patterns and performance with substantial data.
func BenchmarkLargePayload(b *testing.B) {
	server := mockGotenbergServerForBench()
	defer server.Close()

	client := NewClient(nil, server.URL)

	// Create large HTML content
	var htmlBuilder strings.Builder
	htmlBuilder.WriteString("<html><body>")
	for i := 0; i < 1000; i++ {
		htmlBuilder.WriteString("<p>This is paragraph ")
		htmlBuilder.WriteString(string(rune('0' + (i % 10))))
		htmlBuilder.WriteString(" with some content to make the HTML larger.</p>")
	}
	htmlBuilder.WriteString("</body></html>")
	largeHTML := []byte(htmlBuilder.String())

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		resp, err := client.ConvertHTMLToPDF(largeHTML)
		if err != nil {
			b.Fatalf("ConvertHTMLToPDF failed: %v", err)
		}
		_ = resp
	}
}
