package gotenberg

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"
)

// Client represents a client for working with Gotenberg API
type Client struct {
	httpClient *http.Client
	baseURL    string
}

// NewClient creates a new Gotenberg client
func NewClient(httpClient *http.Client, baseURL string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	// Remove trailing slash from URL
	baseURL = strings.TrimSuffix(baseURL, "/")

	return &Client{
		httpClient: httpClient,
		baseURL:    baseURL,
	}
}

// URLToPDFOption represents a functional option for URL to PDF conversion
type URLToPDFOption func(*urlToPDFConfig)

// urlToPDFConfig internal configuration for URL to PDF conversion
type urlToPDFConfig struct {
	// Page Properties
	SinglePage              *bool
	PaperWidth              *float64
	PaperHeight             *float64
	MarginTop               *float64
	MarginBottom            *float64
	MarginLeft              *float64
	MarginRight             *float64
	PreferCSSPageSize       *bool
	GenerateDocumentOutline *bool
	GenerateTaggedPDF       *bool
	PrintBackground         *bool
	OmitBackground          *bool
	Landscape               *bool
	Scale                   *float64
	NativePageRanges        *string

	// Output options
	OutputFilename *string

	// Webhook options
	WebhookURL          *string
	WebhookErrorURL     *string
	WebhookMethod       *string
	WebhookErrorMethod  *string
	WebhookExtraHeaders map[string]string
}

// HTMLToPDFOption represents a functional option for HTML to PDF conversion
type HTMLToPDFOption func(*htmlToPDFConfig)

// htmlToPDFConfig internal configuration for HTML to PDF conversion
type htmlToPDFConfig struct {
	// Additional files (images, CSS, fonts, etc.)
	AdditionalFiles map[string][]byte

	// Header and Footer
	HeaderHTML []byte
	FooterHTML []byte

	// Page Properties (same as for URL)
	SinglePage              *bool
	PaperWidth              *float64
	PaperHeight             *float64
	MarginTop               *float64
	MarginBottom            *float64
	MarginLeft              *float64
	MarginRight             *float64
	PreferCSSPageSize       *bool
	GenerateDocumentOutline *bool
	GenerateTaggedPDF       *bool
	PrintBackground         *bool
	OmitBackground          *bool
	Landscape               *bool
	Scale                   *float64
	NativePageRanges        *string

	// Output options
	OutputFilename *string

	// Webhook options
	WebhookURL          *string
	WebhookErrorURL     *string
	WebhookMethod       *string
	WebhookErrorMethod  *string
	WebhookExtraHeaders map[string]string
}

// MarkdownToPDFOption represents a functional option for Markdown to PDF conversion
type MarkdownToPDFOption func(*markdownToPDFConfig)

// markdownToPDFConfig internal configuration for Markdown to PDF conversion
type markdownToPDFConfig struct {
	// Additional files
	AdditionalFiles map[string][]byte

	// Header and Footer
	HeaderHTML []byte
	FooterHTML []byte

	// Page Properties (same as for URL)
	SinglePage              *bool
	PaperWidth              *float64
	PaperHeight             *float64
	MarginTop               *float64
	MarginBottom            *float64
	MarginLeft              *float64
	MarginRight             *float64
	PreferCSSPageSize       *bool
	GenerateDocumentOutline *bool
	GenerateTaggedPDF       *bool
	PrintBackground         *bool
	OmitBackground          *bool
	Landscape               *bool
	Scale                   *float64
	NativePageRanges        *string

	// Output options
	OutputFilename *string

	// Webhook options
	WebhookURL          *string
	WebhookErrorURL     *string
	WebhookMethod       *string
	WebhookErrorMethod  *string
	WebhookExtraHeaders map[string]string
}

// ConvertURLToPDF converts URL to PDF
func (c *Client) ConvertURLToPDF(url string, opts ...URLToPDFOption) (*PDFResponse, error) {
	if url == "" {
		return nil, fmt.Errorf("URL is required")
	}

	// Apply all options
	config := &urlToPDFConfig{}
	for _, opt := range opts {
		opt(config)
	}

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Add URL
	if err := writer.WriteField("url", url); err != nil {
		return nil, fmt.Errorf("failed to write url field: %w", err)
	}

	// Add page properties
	if err := c.addPageProperties(writer, pageProperties{
		SinglePage:              config.SinglePage,
		PaperWidth:              config.PaperWidth,
		PaperHeight:             config.PaperHeight,
		MarginTop:               config.MarginTop,
		MarginBottom:            config.MarginBottom,
		MarginLeft:              config.MarginLeft,
		MarginRight:             config.MarginRight,
		PreferCSSPageSize:       config.PreferCSSPageSize,
		GenerateDocumentOutline: config.GenerateDocumentOutline,
		GenerateTaggedPDF:       config.GenerateTaggedPDF,
		PrintBackground:         config.PrintBackground,
		OmitBackground:          config.OmitBackground,
		Landscape:               config.Landscape,
		Scale:                   config.Scale,
		NativePageRanges:        config.NativePageRanges,
	}); err != nil {
		return nil, fmt.Errorf("failed to add page properties: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close writer: %w", err)
	}

	req, err := http.NewRequest("POST", c.baseURL+"/forms/chromium/convert/url", &buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.ContentLength = int64(buf.Len())

	// Add webhook headers if specified
	c.addWebhookHeaders(req, webhookOptions{
		URL:          config.WebhookURL,
		ErrorURL:     config.WebhookErrorURL,
		Method:       config.WebhookMethod,
		ErrorMethod:  config.WebhookErrorMethod,
		ExtraHeaders: config.WebhookExtraHeaders,
	})

	// Add header for filename
	if config.OutputFilename != nil {
		req.Header.Set("Gotenberg-Output-Filename", *config.OutputFilename)
	}

	return c.doRequest(req)
}

// ConvertHTMLToPDF converts HTML to PDF
func (c *Client) ConvertHTMLToPDF(indexHTML []byte, opts ...HTMLToPDFOption) (*PDFResponse, error) {
	if len(indexHTML) == 0 {
		return nil, fmt.Errorf("indexHTML is required")
	}

	// Apply all options
	config := &htmlToPDFConfig{}
	for _, opt := range opts {
		opt(config)
	}

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Add index.html
	if err := c.addFileField(writer, "files", "index.html", indexHTML); err != nil {
		return nil, fmt.Errorf("failed to add index.html: %w", err)
	}

	// Add additional files
	for filename, content := range config.AdditionalFiles {
		if err := c.addFileField(writer, "files", filename, content); err != nil {
			return nil, fmt.Errorf("failed to add file %s: %w", filename, err)
		}
	}

	// Add header.html if provided
	if len(config.HeaderHTML) > 0 {
		if err := c.addFileField(writer, "files", "header.html", config.HeaderHTML); err != nil {
			return nil, fmt.Errorf("failed to add header.html: %w", err)
		}
	}

	// Add footer.html if provided
	if len(config.FooterHTML) > 0 {
		if err := c.addFileField(writer, "files", "footer.html", config.FooterHTML); err != nil {
			return nil, fmt.Errorf("failed to add footer.html: %w", err)
		}
	}

	// Add page properties
	if err := c.addPageProperties(writer, pageProperties{
		SinglePage:              config.SinglePage,
		PaperWidth:              config.PaperWidth,
		PaperHeight:             config.PaperHeight,
		MarginTop:               config.MarginTop,
		MarginBottom:            config.MarginBottom,
		MarginLeft:              config.MarginLeft,
		MarginRight:             config.MarginRight,
		PreferCSSPageSize:       config.PreferCSSPageSize,
		GenerateDocumentOutline: config.GenerateDocumentOutline,
		GenerateTaggedPDF:       config.GenerateTaggedPDF,
		PrintBackground:         config.PrintBackground,
		OmitBackground:          config.OmitBackground,
		Landscape:               config.Landscape,
		Scale:                   config.Scale,
		NativePageRanges:        config.NativePageRanges,
	}); err != nil {
		return nil, fmt.Errorf("failed to add page properties: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close writer: %w", err)
	}

	req, err := http.NewRequest("POST", c.baseURL+"/forms/chromium/convert/html", &buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.ContentLength = int64(buf.Len())

	// Add webhook headers if specified
	c.addWebhookHeaders(req, webhookOptions{
		URL:          config.WebhookURL,
		ErrorURL:     config.WebhookErrorURL,
		Method:       config.WebhookMethod,
		ErrorMethod:  config.WebhookErrorMethod,
		ExtraHeaders: config.WebhookExtraHeaders,
	})

	// Add header for filename
	if config.OutputFilename != nil {
		req.Header.Set("Gotenberg-Output-Filename", *config.OutputFilename)
	}

	return c.doRequest(req)
}

// ConvertMarkdownToPDF converts Markdown to PDF
func (c *Client) ConvertMarkdownToPDF(indexHTML []byte, markdownFiles map[string][]byte, opts ...MarkdownToPDFOption) (*PDFResponse, error) {
	if len(indexHTML) == 0 {
		return nil, fmt.Errorf("indexHTML is required")
	}
	if len(markdownFiles) == 0 {
		return nil, fmt.Errorf("at least one markdown file is required")
	}

	// Apply all options
	config := &markdownToPDFConfig{}
	for _, opt := range opts {
		opt(config)
	}

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Add index.html
	if err := c.addFileField(writer, "files", "index.html", indexHTML); err != nil {
		return nil, fmt.Errorf("failed to add index.html: %w", err)
	}

	// Add markdown files
	for filename, content := range markdownFiles {
		if err := c.addFileField(writer, "files", filename, content); err != nil {
			return nil, fmt.Errorf("failed to add markdown file %s: %w", filename, err)
		}
	}

	// Add additional files
	for filename, content := range config.AdditionalFiles {
		if err := c.addFileField(writer, "files", filename, content); err != nil {
			return nil, fmt.Errorf("failed to add file %s: %w", filename, err)
		}
	}

	// Add header.html if provided
	if len(config.HeaderHTML) > 0 {
		if err := c.addFileField(writer, "files", "header.html", config.HeaderHTML); err != nil {
			return nil, fmt.Errorf("failed to add header.html: %w", err)
		}
	}

	// Add footer.html if provided
	if len(config.FooterHTML) > 0 {
		if err := c.addFileField(writer, "files", "footer.html", config.FooterHTML); err != nil {
			return nil, fmt.Errorf("failed to add footer.html: %w", err)
		}
	}

	// Add page properties
	if err := c.addPageProperties(writer, pageProperties{
		SinglePage:              config.SinglePage,
		PaperWidth:              config.PaperWidth,
		PaperHeight:             config.PaperHeight,
		MarginTop:               config.MarginTop,
		MarginBottom:            config.MarginBottom,
		MarginLeft:              config.MarginLeft,
		MarginRight:             config.MarginRight,
		PreferCSSPageSize:       config.PreferCSSPageSize,
		GenerateDocumentOutline: config.GenerateDocumentOutline,
		GenerateTaggedPDF:       config.GenerateTaggedPDF,
		PrintBackground:         config.PrintBackground,
		OmitBackground:          config.OmitBackground,
		Landscape:               config.Landscape,
		Scale:                   config.Scale,
		NativePageRanges:        config.NativePageRanges,
	}); err != nil {
		return nil, fmt.Errorf("failed to add page properties: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close writer: %w", err)
	}

	req, err := http.NewRequest("POST", c.baseURL+"/forms/chromium/convert/markdown", &buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.ContentLength = int64(buf.Len())

	// Add webhook headers if specified
	c.addWebhookHeaders(req, webhookOptions{
		URL:          config.WebhookURL,
		ErrorURL:     config.WebhookErrorURL,
		Method:       config.WebhookMethod,
		ErrorMethod:  config.WebhookErrorMethod,
		ExtraHeaders: config.WebhookExtraHeaders,
	})

	// Add header for filename
	if config.OutputFilename != nil {
		req.Header.Set("Gotenberg-Output-Filename", *config.OutputFilename)
	}

	return c.doRequest(req)
}
