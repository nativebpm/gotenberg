package gotenberg

func WithPaperSize(width, height float64) ConvOption {
	return func(c *convConfig) {
		c.PaperWidth = &width
		c.PaperHeight = &height
	}
}

func WithMargins(top, right, bottom, left float64) ConvOption {
	return func(c *convConfig) {
		c.MarginTop = &top
		c.MarginRight = &right
		c.MarginBottom = &bottom
		c.MarginLeft = &left
	}
}

func WithSinglePage(enabled bool) ConvOption {
	return func(c *convConfig) {
		c.SinglePage = &enabled
	}
}

func WithLandscape(enabled bool) ConvOption {
	return func(c *convConfig) {
		c.Landscape = &enabled
	}
}

func WithPrintBackground(enabled bool) ConvOption {
	return func(c *convConfig) {
		c.PrintBackground = &enabled
	}
}

func WithScale(scale float64) ConvOption {
	return func(c *convConfig) {
		c.Scale = &scale
	}
}

func WithOutputFilename(filename string) ConvOption {
	return func(c *convConfig) {
		c.OutputFilename = &filename
	}
}

func WithWebhook(url, errorURL string) ConvOption {
	return func(c *convConfig) {
		c.WebhookURL = &url
		c.WebhookErrorURL = &errorURL
	}
}

func WithWebhookMethods(method, errorMethod string) ConvOption {
	return func(c *convConfig) {
		c.WebhookMethod = &method
		c.WebhookErrorMethod = &errorMethod
	}
}

func WithWebhookExtraHeaders(headers map[string]string) ConvOption {
	return func(c *convConfig) {
		if headers == nil {
			c.WebhookExtraHeaders = nil
			return
		}
		// Make a shallow copy to avoid retaining a reference to a caller-owned map
		copied := make(map[string]string, len(headers))
		for k, v := range headers {
			copied[k] = v
		}
		c.WebhookExtraHeaders = copied
	}
}

func WithHTMLAdditionalFiles(files map[string][]byte) ConvOption {
	return func(c *convConfig) {
		c.AdditionalFiles = files
	}
}

func WithHTMLHeader(headerHTML []byte) ConvOption {
	return func(c *convConfig) {
		c.HeaderHTML = headerHTML
	}
}

func WithHTMLFooter(footerHTML []byte) ConvOption {
	return func(c *convConfig) {
		c.FooterHTML = footerHTML
	}
}

func WithPaperSizeA4() ConvOption {
	return WithPaperSize(PaperSizeA4[0], PaperSizeA4[1])
}

func WithPaperSizeLetter() ConvOption {
	return WithPaperSize(PaperSizeLetter[0], PaperSizeLetter[1])
}
