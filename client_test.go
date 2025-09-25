package gotenberg

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// mockGotenbergServer creates a mock Gotenberg server for testing
func mockGotenbergServer(t *testing.T, expectedPath string, response []byte, statusCode int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != expectedPath {
			t.Errorf("Expected path %s, got %s", expectedPath, r.URL.Path)
		}

		if r.Method != "POST" {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		// Check Content-Type
		if !strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data") {
			t.Errorf("Expected multipart/form-data content type, got %s", r.Header.Get("Content-Type"))
		}

		// Set response headers
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=test.pdf")
		w.Header().Set("Gotenberg-Trace", "test-trace-id")

		w.WriteHeader(statusCode)
		w.Write(response)
	}))
}

func TestNewClient(t *testing.T) {
	tests := []struct {
		name       string
		httpClient *http.Client
		baseURL    string
		expected   *Client
	}{
		{
			name:       "with custom http client",
			httpClient: &http.Client{},
			baseURL:    "http://localhost:3000",
			expected: &Client{
				httpClient: &http.Client{},
				baseURL:    "http://localhost:3000",
			},
		},
		{
			name:       "with nil http client",
			httpClient: nil,
			baseURL:    "http://localhost:3000",
			expected: &Client{
				httpClient: http.DefaultClient,
				baseURL:    "http://localhost:3000",
			},
		},
		{
			name:       "with trailing slash in URL",
			httpClient: nil,
			baseURL:    "http://localhost:3000/",
			expected: &Client{
				httpClient: http.DefaultClient,
				baseURL:    "http://localhost:3000",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewClient(tt.httpClient, tt.baseURL)

			if got.baseURL != tt.expected.baseURL {
				t.Errorf("NewClient() baseURL = %v, expected %v", got.baseURL, tt.expected.baseURL)
			}

			// For nil input, check if default client was set
			if tt.httpClient == nil && got.httpClient != http.DefaultClient {
				t.Errorf("NewClient() should use DefaultClient when nil passed")
			}

			// For non-nil input, verify the same instance was used
			if tt.httpClient != nil && got.httpClient != tt.httpClient {
				t.Errorf("NewClient() should use provided client")
			}
		})
	}
}

func TestClient_ConvertURLToPDF(t *testing.T) {
	pdfContent := []byte("%PDF-1.4 test content")
	server := mockGotenbergServer(t, "/forms/chromium/convert/url", pdfContent, http.StatusOK)
	defer server.Close()

	client := NewClient(nil, server.URL)

	tests := []struct {
		name    string
		url     string
		options []ConvOption
		wantErr bool
	}{
		{
			name:    "successful conversion without options",
			url:     "https://example.com",
			options: nil,
			wantErr: false,
		},
		{
			name: "successful conversion with options",
			url:  "https://example.com",
			options: []ConvOption{
				WithPaperSize(8.5, 11),
				WithMargins(1, 1, 1, 1),
				WithLandscape(true),
				WithOutputFilename("test.pdf"),
			},
			wantErr: false,
		},
		{
			name: "successful conversion with webhook",
			url:  "https://example.com",
			options: []ConvOption{
				WithWebhook("https://webhook.example.com", "https://webhook-error.example.com"),
				WithWebhookMethods("POST", "PUT"),
			},
			wantErr: false,
		},
		{
			name:    "empty URL should fail",
			url:     "",
			options: nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.ConvertURLToPDF(tt.url, tt.options...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertURLToPDF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && resp == nil {
				t.Errorf("ConvertURLToPDF() response is nil")
				return
			}

			if !tt.wantErr {
				if !bytes.Equal(resp.PDF, pdfContent) {
					t.Errorf("ConvertURLToPDF() PDF content mismatch")
				}
				if resp.ContentType != "application/pdf" {
					t.Errorf("ConvertURLToPDF() ContentType = %v, expected application/pdf", resp.ContentType)
				}
				if resp.Trace != "test-trace-id" {
					t.Errorf("ConvertURLToPDF() Trace = %v, expected test-trace-id", resp.Trace)
				}
			}
		})
	}
}

func TestClient_ConvertHTMLToPDF(t *testing.T) {
	pdfContent := []byte("%PDF-1.4 test content")
	server := mockGotenbergServer(t, "/forms/chromium/convert/html", pdfContent, http.StatusOK)
	defer server.Close()

	client := NewClient(nil, server.URL)
	indexHTML := []byte("<html><body><h1>Test</h1></body></html>")

	tests := []struct {
		name      string
		indexHTML []byte
		options   []ConvOption
		wantErr   bool
	}{
		{
			name:      "successful conversion without options",
			indexHTML: indexHTML,
			options:   nil,
			wantErr:   false,
		},
		{
			name:      "successful conversion with options",
			indexHTML: indexHTML,
			options: []ConvOption{
				WithPaperSize(8.5, 11),
				WithMargins(1, 1, 1, 1),
				WithLandscape(true),
				WithOutputFilename("test.pdf"),
				WithHTMLAdditionalFiles(map[string][]byte{
					"style.css": []byte("body { font-family: Arial; }"),
				}),
				WithHTMLHeader([]byte("<html><body>Header</body></html>")),
				WithHTMLFooter([]byte("<html><body>Footer</body></html>")),
			},
			wantErr: false,
		},
		{
			name:      "empty HTML should fail",
			indexHTML: []byte{},
			options:   nil,
			wantErr:   true,
		},
		{
			name:      "nil HTML should fail",
			indexHTML: nil,
			options:   nil,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.ConvertHTMLToPDF(tt.indexHTML, tt.options...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertHTMLToPDF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && resp == nil {
				t.Errorf("ConvertHTMLToPDF() response is nil")
				return
			}

			if !tt.wantErr {
				if !bytes.Equal(resp.PDF, pdfContent) {
					t.Errorf("ConvertHTMLToPDF() PDF content mismatch")
				}
			}
		})
	}
}

func TestClient_ConvertMarkdownToPDF(t *testing.T) {
	pdfContent := []byte("%PDF-1.4 test content")
	server := mockGotenbergServer(t, "/forms/chromium/convert/markdown", pdfContent, http.StatusOK)
	defer server.Close()

	client := NewClient(nil, server.URL)
	indexHTML := []byte("<html><body>{{ .markdown }}</body></html>")
	markdownFiles := map[string][]byte{
		"test.md": []byte("# Test Markdown\n\nThis is a test."),
	}

	tests := []struct {
		name          string
		indexHTML     []byte
		markdownFiles map[string][]byte
		options       []ConvOption
		wantErr       bool
	}{
		{
			name:          "successful conversion without options",
			indexHTML:     indexHTML,
			markdownFiles: markdownFiles,
			options:       nil,
			wantErr:       false,
		},
		{
			name:          "successful conversion with options",
			indexHTML:     indexHTML,
			markdownFiles: markdownFiles,
			options: []ConvOption{
				WithPaperSize(8.5, 11),
				WithMargins(1, 1, 1, 1),
				WithLandscape(true),
				WithOutputFilename("test.pdf"),
				WithHTMLAdditionalFiles(map[string][]byte{
					"style.css": []byte("body { font-family: Arial; }"),
				}),
			},
			wantErr: false,
		},
		{
			name:          "empty HTML should fail",
			indexHTML:     []byte{},
			markdownFiles: markdownFiles,
			options:       nil,
			wantErr:       true,
		},
		{
			name:          "empty markdown files should fail",
			indexHTML:     indexHTML,
			markdownFiles: map[string][]byte{},
			options:       nil,
			wantErr:       true,
		},
		{
			name:          "nil markdown files should fail",
			indexHTML:     indexHTML,
			markdownFiles: nil,
			options:       nil,
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.ConvertMarkdownToPDF(tt.indexHTML, tt.markdownFiles, tt.options...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertMarkdownToPDF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && resp == nil {
				t.Errorf("ConvertMarkdownToPDF() response is nil")
				return
			}

			if !tt.wantErr {
				if !bytes.Equal(resp.PDF, pdfContent) {
					t.Errorf("ConvertMarkdownToPDF() PDF content mismatch")
				}
			}
		})
	}
}

func TestClient_ErrorResponse(t *testing.T) {
	errorMsg := "Bad Request"
	server := mockGotenbergServer(t, "/forms/chromium/convert/url", []byte(errorMsg), http.StatusBadRequest)
	defer server.Close()

	client := NewClient(nil, server.URL)

	resp, err := client.ConvertURLToPDF("https://example.com")

	if err == nil {
		t.Errorf("Expected error, got nil")
		return
	}

	if resp != nil {
		t.Errorf("Expected nil response, got %v", resp)
	}

	gotErr, ok := err.(*GotenbergError)
	if !ok {
		t.Errorf("Expected GotenbergError, got %T", err)
		return
	}

	if gotErr.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, gotErr.StatusCode)
	}

	if gotErr.Message != errorMsg {
		t.Errorf("Expected error message %q, got %q", errorMsg, gotErr.Message)
	}
}

func TestClient_WebhookResponse(t *testing.T) {
	// Webhook response returns 204 No Content
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Gotenberg-Trace", "webhook-trace-id")
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := NewClient(nil, server.URL)

	resp, err := client.ConvertURLToPDF("https://example.com",
		WithWebhook("https://webhook.example.com", "https://webhook-error.example.com"))

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
		return
	}

	if resp == nil {
		t.Errorf("Expected response, got nil")
		return
	}

	if resp.PDF != nil {
		t.Errorf("Expected nil PDF for webhook response, got %v", resp.PDF)
	}

	if resp.Trace != "webhook-trace-id" {
		t.Errorf("Expected trace 'webhook-trace-id', got %q", resp.Trace)
	}
}

func TestPaperSizes(t *testing.T) {
	// Test that paper sizes are defined
	if PaperSizeA4[0] != 8.27 || PaperSizeA4[1] != 11.7 {
		t.Errorf("PaperSizeA4 incorrect: got %v, expected [8.27, 11.7]", PaperSizeA4)
	}

	if PaperSizeLetter[0] != 8.5 || PaperSizeLetter[1] != 11 {
		t.Errorf("PaperSizeLetter incorrect: got %v, expected [8.5, 11]", PaperSizeLetter)
	}
}
