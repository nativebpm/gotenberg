package gotenberg

type convConfig struct {
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

	OutputFilename *string

	WebhookURL          *string
	WebhookErrorURL     *string
	WebhookMethod       *string
	WebhookErrorMethod  *string
	WebhookExtraHeaders map[string]string

	AdditionalFiles map[string][]byte

	HeaderHTML []byte
	FooterHTML []byte
}

type ConvOption func(*convConfig)
