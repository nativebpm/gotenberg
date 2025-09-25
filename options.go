package gotenberg

// WithPaperSize sets paper size
func WithPaperSize(width, height float64) ConvOption {
	return func(c *convConfig) {
		c.PaperWidth = &width
		c.PaperHeight = &height
	}
}

// WithMargins sets margins
func WithMargins(top, right, bottom, left float64) ConvOption {
	return func(c *convConfig) {
		c.MarginTop = &top
		c.MarginRight = &right
		c.MarginBottom = &bottom
		c.MarginLeft = &left
	}
}

// WithSinglePage sets single page mode
func WithSinglePage(enabled bool) ConvOption {
	return func(c *convConfig) {
		c.SinglePage = &enabled
	}
}

// WithLandscape sets landscape orientation
func WithLandscape(enabled bool) ConvOption {
	return func(c *convConfig) {
		c.Landscape = &enabled
	}
}

// WithPrintBackground sets print background
func WithPrintBackground(enabled bool) ConvOption {
	return func(c *convConfig) {
		c.PrintBackground = &enabled
	}
}

// WithScale sets scale
func WithScale(scale float64) ConvOption {
	return func(c *convConfig) {
		c.Scale = &scale
	}
}

// WithOutputFilename sets output filename
func WithOutputFilename(filename string) ConvOption {
	return func(c *convConfig) {
		c.OutputFilename = &filename
	}
}

// WithWebhook sets webhook configuration
func WithWebhook(url, errorURL string) ConvOption {
	return func(c *convConfig) {
		c.WebhookURL = &url
		c.WebhookErrorURL = &errorURL
	}
}

// WithWebhookMethods sets webhook HTTP methods
func WithWebhookMethods(method, errorMethod string) ConvOption {
	return func(c *convConfig) {
		c.WebhookMethod = &method
		c.WebhookErrorMethod = &errorMethod
	}
}

// WithWebhookExtraHeaders sets extra headers for webhook
func WithWebhookExtraHeaders(headers map[string]string) ConvOption {
	return func(c *convConfig) {
		c.WebhookExtraHeaders = headers
	}
}

// WithHTMLAdditionalFiles adds additional files for HTML conversion
func WithHTMLAdditionalFiles(files map[string][]byte) ConvOption {
	return func(c *convConfig) {
		c.AdditionalFiles = files
	}
}

// WithHTMLHeader sets header HTML for HTML conversion
func WithHTMLHeader(headerHTML []byte) ConvOption {
	return func(c *convConfig) {
		c.HeaderHTML = headerHTML
	}
}

// WithHTMLFooter sets footer HTML for HTML conversion
func WithHTMLFooter(footerHTML []byte) ConvOption {
	return func(c *convConfig) {
		c.FooterHTML = footerHTML
	}
}

// WithPaperSizeA4 returns WithPaperSizeA4 paper size option
func WithPaperSizeA4() ConvOption {
	return WithPaperSize(PaperSizeA4[0], PaperSizeA4[1])
}

// WithPaperSizeLetter returns WithPaperSizeLetter paper size option
func WithPaperSizeLetter() ConvOption {
	return WithPaperSize(PaperSizeLetter[0], PaperSizeLetter[1])
}
