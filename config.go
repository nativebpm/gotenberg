package gotenberg

// convConfig contains common configuration options for all PDF conversion types
type convConfig struct {
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

	// Additional files (images, CSS, fonts, etc.)
	AdditionalFiles map[string][]byte

	// Header and Footer
	HeaderHTML []byte
	FooterHTML []byte
}

// ConvOption represents a functional option for URL to PDF conversion
type ConvOption func(*convConfig)
