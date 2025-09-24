package gotenberg

// Option functions for URL to PDF conversion

// WithPaperSize sets paper size
func WithPaperSize(width, height float64) URLToPDFOption {
	return func(c *urlToPDFConfig) {
		c.PaperWidth = Float64(width)
		c.PaperHeight = Float64(height)
	}
}

// WithMargins sets margins
func WithMargins(top, right, bottom, left float64) URLToPDFOption {
	return func(c *urlToPDFConfig) {
		c.MarginTop = Float64(top)
		c.MarginRight = Float64(right)
		c.MarginBottom = Float64(bottom)
		c.MarginLeft = Float64(left)
	}
}

// WithSinglePage sets single page mode
func WithSinglePage(enabled bool) URLToPDFOption {
	return func(c *urlToPDFConfig) {
		c.SinglePage = Bool(enabled)
	}
}

// WithLandscape sets landscape orientation
func WithLandscape(enabled bool) URLToPDFOption {
	return func(c *urlToPDFConfig) {
		c.Landscape = Bool(enabled)
	}
}

// WithPrintBackground sets print background
func WithPrintBackground(enabled bool) URLToPDFOption {
	return func(c *urlToPDFConfig) {
		c.PrintBackground = Bool(enabled)
	}
}

// WithScale sets scale
func WithScale(scale float64) URLToPDFOption {
	return func(c *urlToPDFConfig) {
		c.Scale = Float64(scale)
	}
}

// WithOutputFilename sets output filename
func WithOutputFilename(filename string) URLToPDFOption {
	return func(c *urlToPDFConfig) {
		c.OutputFilename = String(filename)
	}
}

// WithWebhook sets webhook configuration
func WithWebhook(url, errorURL string) URLToPDFOption {
	return func(c *urlToPDFConfig) {
		c.WebhookURL = String(url)
		c.WebhookErrorURL = String(errorURL)
	}
}

// WithWebhookMethods sets webhook HTTP methods
func WithWebhookMethods(method, errorMethod string) URLToPDFOption {
	return func(c *urlToPDFConfig) {
		c.WebhookMethod = String(method)
		c.WebhookErrorMethod = String(errorMethod)
	}
}

// WithWebhookExtraHeaders sets extra headers for webhook
func WithWebhookExtraHeaders(headers map[string]string) URLToPDFOption {
	return func(c *urlToPDFConfig) {
		c.WebhookExtraHeaders = headers
	}
}

// Option functions for HTML to PDF conversion

// WithHTMLPaperSize sets paper size for HTML conversion
func WithHTMLPaperSize(width, height float64) HTMLToPDFOption {
	return func(c *htmlToPDFConfig) {
		c.PaperWidth = Float64(width)
		c.PaperHeight = Float64(height)
	}
}

// WithHTMLMargins sets margins for HTML conversion
func WithHTMLMargins(top, right, bottom, left float64) HTMLToPDFOption {
	return func(c *htmlToPDFConfig) {
		c.MarginTop = Float64(top)
		c.MarginRight = Float64(right)
		c.MarginBottom = Float64(bottom)
		c.MarginLeft = Float64(left)
	}
}

// WithHTMLSinglePage sets single page mode for HTML conversion
func WithHTMLSinglePage(enabled bool) HTMLToPDFOption {
	return func(c *htmlToPDFConfig) {
		c.SinglePage = Bool(enabled)
	}
}

// WithHTMLLandscape sets landscape orientation for HTML conversion
func WithHTMLLandscape(enabled bool) HTMLToPDFOption {
	return func(c *htmlToPDFConfig) {
		c.Landscape = Bool(enabled)
	}
}

// WithHTMLPrintBackground sets print background for HTML conversion
func WithHTMLPrintBackground(enabled bool) HTMLToPDFOption {
	return func(c *htmlToPDFConfig) {
		c.PrintBackground = Bool(enabled)
	}
}

// WithHTMLScale sets scale for HTML conversion
func WithHTMLScale(scale float64) HTMLToPDFOption {
	return func(c *htmlToPDFConfig) {
		c.Scale = Float64(scale)
	}
}

// WithHTMLOutputFilename sets output filename for HTML conversion
func WithHTMLOutputFilename(filename string) HTMLToPDFOption {
	return func(c *htmlToPDFConfig) {
		c.OutputFilename = String(filename)
	}
}

// WithHTMLWebhook sets webhook configuration for HTML conversion
func WithHTMLWebhook(url, errorURL string) HTMLToPDFOption {
	return func(c *htmlToPDFConfig) {
		c.WebhookURL = String(url)
		c.WebhookErrorURL = String(errorURL)
	}
}

// WithHTMLWebhookMethods sets webhook HTTP methods for HTML conversion
func WithHTMLWebhookMethods(method, errorMethod string) HTMLToPDFOption {
	return func(c *htmlToPDFConfig) {
		c.WebhookMethod = String(method)
		c.WebhookErrorMethod = String(errorMethod)
	}
}

// WithHTMLWebhookExtraHeaders sets extra headers for webhook in HTML conversion
func WithHTMLWebhookExtraHeaders(headers map[string]string) HTMLToPDFOption {
	return func(c *htmlToPDFConfig) {
		c.WebhookExtraHeaders = headers
	}
}

// WithAdditionalFiles adds additional files for HTML conversion
func WithAdditionalFiles(files map[string][]byte) HTMLToPDFOption {
	return func(c *htmlToPDFConfig) {
		c.AdditionalFiles = files
	}
}

// WithHeader sets header HTML for HTML conversion
func WithHeader(headerHTML []byte) HTMLToPDFOption {
	return func(c *htmlToPDFConfig) {
		c.HeaderHTML = headerHTML
	}
}

// WithFooter sets footer HTML for HTML conversion
func WithFooter(footerHTML []byte) HTMLToPDFOption {
	return func(c *htmlToPDFConfig) {
		c.FooterHTML = footerHTML
	}
}

// Option functions for Markdown to PDF conversion

// WithMarkdownPaperSize sets paper size for Markdown conversion
func WithMarkdownPaperSize(width, height float64) MarkdownToPDFOption {
	return func(c *markdownToPDFConfig) {
		c.PaperWidth = Float64(width)
		c.PaperHeight = Float64(height)
	}
}

// WithMarkdownMargins sets margins for Markdown conversion
func WithMarkdownMargins(top, right, bottom, left float64) MarkdownToPDFOption {
	return func(c *markdownToPDFConfig) {
		c.MarginTop = Float64(top)
		c.MarginRight = Float64(right)
		c.MarginBottom = Float64(bottom)
		c.MarginLeft = Float64(left)
	}
}

// WithMarkdownSinglePage sets single page mode for Markdown conversion
func WithMarkdownSinglePage(enabled bool) MarkdownToPDFOption {
	return func(c *markdownToPDFConfig) {
		c.SinglePage = Bool(enabled)
	}
}

// WithMarkdownLandscape sets landscape orientation for Markdown conversion
func WithMarkdownLandscape(enabled bool) MarkdownToPDFOption {
	return func(c *markdownToPDFConfig) {
		c.Landscape = Bool(enabled)
	}
}

// WithMarkdownPrintBackground sets print background for Markdown conversion
func WithMarkdownPrintBackground(enabled bool) MarkdownToPDFOption {
	return func(c *markdownToPDFConfig) {
		c.PrintBackground = Bool(enabled)
	}
}

// WithMarkdownScale sets scale for Markdown conversion
func WithMarkdownScale(scale float64) MarkdownToPDFOption {
	return func(c *markdownToPDFConfig) {
		c.Scale = Float64(scale)
	}
}

// WithMarkdownOutputFilename sets output filename for Markdown conversion
func WithMarkdownOutputFilename(filename string) MarkdownToPDFOption {
	return func(c *markdownToPDFConfig) {
		c.OutputFilename = String(filename)
	}
}

// WithMarkdownWebhook sets webhook configuration for Markdown conversion
func WithMarkdownWebhook(url, errorURL string) MarkdownToPDFOption {
	return func(c *markdownToPDFConfig) {
		c.WebhookURL = String(url)
		c.WebhookErrorURL = String(errorURL)
	}
}

// WithMarkdownWebhookMethods sets webhook HTTP methods for Markdown conversion
func WithMarkdownWebhookMethods(method, errorMethod string) MarkdownToPDFOption {
	return func(c *markdownToPDFConfig) {
		c.WebhookMethod = String(method)
		c.WebhookErrorMethod = String(errorMethod)
	}
}

// WithMarkdownWebhookExtraHeaders sets extra headers for webhook in Markdown conversion
func WithMarkdownWebhookExtraHeaders(headers map[string]string) MarkdownToPDFOption {
	return func(c *markdownToPDFConfig) {
		c.WebhookExtraHeaders = headers
	}
}

// WithMarkdownAdditionalFiles adds additional files for Markdown conversion
func WithMarkdownAdditionalFiles(files map[string][]byte) MarkdownToPDFOption {
	return func(c *markdownToPDFConfig) {
		c.AdditionalFiles = files
	}
}

// WithMarkdownHeader sets header HTML for Markdown conversion
func WithMarkdownHeader(headerHTML []byte) MarkdownToPDFOption {
	return func(c *markdownToPDFConfig) {
		c.HeaderHTML = headerHTML
	}
}

// WithMarkdownFooter sets footer HTML for Markdown conversion
func WithMarkdownFooter(footerHTML []byte) MarkdownToPDFOption {
	return func(c *markdownToPDFConfig) {
		c.FooterHTML = footerHTML
	}
}
