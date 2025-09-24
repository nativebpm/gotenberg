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

/*
Бенчмарки для анализа аллокации памяти в пакете Gotenberg:

1. BenchmarkNewClient - Проверяет оверхед создания клиента (важно если клиент не переиспользуется)
2. BenchmarkConvertURLToPDF_* - Основная функциональность с разным количеством опций
3. BenchmarkConvertHTMLToPDF_* - HTML конвертация с файлами (много аллокаций для multipart)
4. BenchmarkConvertMarkdownToPDF_* - Markdown конвертация
5. BenchmarkOptionCreation - Стоимость functional options (замыкания + аллокации в хип)
6. BenchmarkConfigApplication - Применение опций к конфигурации
7. BenchmarkMultipartFormCreation - Самая дорогая операция (multipart/form-data)
8. BenchmarkUtilityFunctions - Pointer helpers (Bool/Float64/String)
9. BenchmarkLargePayload - Поведение с большими данными

Запуск: go test -bench=. -benchmem -benchtime=5s ./pkg/gotenberg/
Анализ: go test -bench=BenchmarkConvertHTMLToPDF_WithFiles -benchmem -memprofile=mem.prof
*/

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

// BenchmarkNewClient benchmarks client creation - важно для понимания оверхеда создания клиента
// Полезно если клиент создается часто, а не переиспользуется
func BenchmarkNewClient(b *testing.B) {
	httpClient := &http.Client{}
	baseURL := "http://localhost:3000"

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = NewClient(httpClient, baseURL)
	}
}

// BenchmarkConvertURLToPDF_NoOptions benchmarks URL to PDF conversion without options
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

// BenchmarkConvertURLToPDF_WithOptions benchmarks URL to PDF conversion with options
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

// BenchmarkConvertURLToPDF_ManyOptions benchmarks URL to PDF conversion with many options
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

// BenchmarkConvertHTMLToPDF_NoOptions benchmarks HTML to PDF conversion without options
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

// BenchmarkConvertHTMLToPDF_WithOptions benchmarks HTML to PDF conversion with options
func BenchmarkConvertHTMLToPDF_WithOptions(b *testing.B) {
	server := mockGotenbergServerForBench()
	defer server.Close()

	client := NewClient(nil, server.URL)
	html := []byte("<html><body><h1>Test</h1></body></html>")

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		resp, err := client.ConvertHTMLToPDF(html,
			WithHTMLPaperSize(8.5, 11),
			WithHTMLMargins(1, 1, 1, 1),
			WithHTMLLandscape(false),
			WithHTMLPrintBackground(true),
			WithHTMLOutputFilename("test.pdf"),
		)
		if err != nil {
			b.Fatalf("ConvertHTMLToPDF failed: %v", err)
		}
		_ = resp
	}
}

// BenchmarkConvertHTMLToPDF_WithFiles benchmarks HTML to PDF conversion with additional files
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
			WithAdditionalFiles(map[string][]byte{
				"style.css": css,
				"logo.png":  image,
			}),
			WithHeader([]byte("<html><body>Header</body></html>")),
			WithFooter([]byte("<html><body>Footer</body></html>")),
		)
		if err != nil {
			b.Fatalf("ConvertHTMLToPDF failed: %v", err)
		}
		_ = resp
	}
}

// BenchmarkConvertMarkdownToPDF_NoOptions benchmarks Markdown to PDF conversion without options
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

// BenchmarkOptionCreation - критично для functional options pattern
// Показывает стоимость создания замыканий и аллокации в хип
func BenchmarkOptionCreation(b *testing.B) {
	b.Run("WithPaperSize", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = WithPaperSize(8.5, 11) // 2 аллокации Float64 + замыкание
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
			_ = WithWebhookExtraHeaders(headers) // Может быть дорого из-за map
		}
	})
}

// BenchmarkConfigApplication benchmarks applying functional options to config
func BenchmarkConfigApplication(b *testing.B) {
	options := []URLToPDFOption{
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
		config := &urlToPDFConfig{}
		for _, opt := range options {
			opt(config)
		}
		_ = config
	}
}

// BenchmarkMultipartFormCreation benchmarks multipart form creation and writing
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

// BenchmarkUtilityFunctions - важно для оптимизации вспомогательных функций
// Каждый вызов Bool/String/Float64 создает аллокацию в хип
func BenchmarkUtilityFunctions(b *testing.B) {
	b.Run("Bool", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = Bool(true) // 1 аллокация - escape to heap
		}
	})

	b.Run("Float64", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = Float64(1.5) // 1 аллокация - escape to heap
		}
	})
}

// BenchmarkLargePayload benchmarks performance with large HTML payloads
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
