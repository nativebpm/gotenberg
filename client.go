package gotenberg

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
}

func NewClient(httpClient *http.Client, baseURL string) *Client {
	u, err := url.Parse(strings.TrimSuffix(baseURL, "/"))
	if err != nil {
		panic(fmt.Sprintf("invalid base URL: %s", err))
	}

	return &Client{
		httpClient: httpClient,
		baseURL:    u,
	}
}

// ConvertURLToPDF converts URL to PDF
func (c Client) ConvertURLToPDF(ctx context.Context, url string, opts ...ConvOption) (*http.Response, error) {
	if url == "" {
		return nil, fmt.Errorf("URL is required")
	}

	// Apply all options
	config := &convConfig{}
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

	req, err := http.NewRequestWithContext(ctx, "POST", c.baseURL.JoinPath("/forms/chromium/convert/url").String(), &buf)
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

	return c.httpClient.Do(req)
}

// ConvertHTMLToPDF converts HTML to PDF
func (c Client) ConvertHTMLToPDF(ctx context.Context, indexHTML []byte, opts ...ConvOption) (*http.Response, error) {
	if len(indexHTML) == 0 {
		return nil, fmt.Errorf("indexHTML is required")
	}

	// Apply all options
	config := &convConfig{}
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

	req, err := http.NewRequestWithContext(ctx, "POST", c.baseURL.JoinPath("/forms/chromium/convert/html").String(), &buf)
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

	return c.httpClient.Do(req)
}

// ConvertMarkdownToPDF converts Markdown to PDF
func (c Client) ConvertMarkdownToPDF(ctx context.Context, indexHTML []byte, markdownFiles map[string][]byte, opts ...ConvOption) (*http.Response, error) {
	if len(indexHTML) == 0 {
		return nil, fmt.Errorf("indexHTML is required")
	}
	if len(markdownFiles) == 0 {
		return nil, fmt.Errorf("at least one markdown file is required")
	}

	// Apply all options
	config := &convConfig{}
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

	req, err := http.NewRequestWithContext(ctx, "POST", c.baseURL.JoinPath("/forms/chromium/convert/markdown").String(), &buf)
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

	return c.httpClient.Do(req)
}
