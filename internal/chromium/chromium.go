// Package chromium provides a client for the Gotenberg Chromium service.
// It offers a convenient API for converting HTML, URLs, and Markdown to PDF documents and screenshots.
package chromium

import (
	"context"
	"io"
	"strconv"
	"time"

	"github.com/nativebpm/gotenberg/internal/gotenberg"
	"github.com/nativebpm/connectors/httpclient"
)

// Chromium represents a Gotenberg conversion request builder.
type Chromium struct {
	*gotenberg.Gotenberg
}

func NewChromium(client *httpclient.HTTPClient) *Chromium {
	return &Chromium{
		Gotenberg: gotenberg.NewGotenberg(client),
	}
}

// ConvertHTML creates a request to convert HTML content to PDF.
// The html parameter should contain the HTML content to be converted.
func (r *Chromium) ConvertHTML(ctx context.Context, html io.Reader) *Chromium {
	r.Req = r.Client.Multipart(ctx, "/forms/chromium/convert/html").File("files", "index.html", html)
	return r
}

// ConvertURL creates a request to convert a web page at the given URL to PDF.
func (r *Chromium) ConvertURL(ctx context.Context, url string) *Chromium {
	r.Req = r.Client.Multipart(ctx, "/forms/chromium/convert/url").Param("url", url)
	return r
}

// ConvertMarkdown creates a request to convert Markdown content to PDF.
func (r *Chromium) ConvertMarkdown(ctx context.Context, html io.Reader) *Chromium {
	r.Req = r.Client.Multipart(ctx, "/forms/chromium/convert/markdown").File("files", "index.html", html)
	return r
}

// ScreenshotURL creates a request to take a screenshot of a web page at the given URL.
func (r *Chromium) ScreenshotURL(ctx context.Context, url string) *Chromium {
	r.Req = r.Client.Multipart(ctx, "/forms/chromium/screenshot/url").Param("url", url)
	return r
}

// ScreenshotHTML creates a request to take a screenshot of HTML content.
func (r *Chromium) ScreenshotHTML(ctx context.Context, html io.Reader) *Chromium {
	r.Req = r.Client.Multipart(ctx, "/forms/chromium/screenshot/html").File("files", "index.html", html)
	return r
}

// ScreenshotMarkdown creates a request to take a screenshot of Markdown content.
func (r *Chromium) ScreenshotMarkdown(ctx context.Context, html io.Reader) *Chromium {
	r.Req = r.Client.Multipart(ctx, "/forms/chromium/screenshot/markdown").File("files", "index.html", html)
	return r
}

// Send executes the conversion request and returns the response.
// Returns an error if the request fails or the conversion cannot be completed.
func (r *Chromium) Send() (*gotenberg.Response, error) {
	return r.Gotenberg.Send()
}

// Header adds a header to the conversion request.
func (r *Chromium) Header(key, value string) *Chromium {
	r.Gotenberg.Header(key, value)
	return r
}

// Param adds a form parameter to the conversion request.
func (r *Chromium) Param(key, value string) *Chromium {
	r.Gotenberg.Param(key, value)
	return r
}

// Bool adds a boolean form parameter to the conversion request.
func (r *Chromium) Bool(fieldName string, value bool) *Chromium {
	r.Gotenberg.Bool(fieldName, value)
	return r
}

// Float adds a float64 form parameter to the conversion request.
func (r *Chromium) Float(fieldName string, value float64) *Chromium {
	r.Gotenberg.Float(fieldName, value)
	return r
}

// File adds a file to the conversion request.
func (r *Chromium) File(filename string, content io.Reader) *Chromium {
	r.Gotenberg.File("files", filename, content)
	return r
}

// Timeout sets a timeout for the request.
func (r *Chromium) Timeout(duration time.Duration) *Chromium {
	r.Gotenberg.Timeout(duration)
	return r
}

// ScreenshotWidth sets the device screen width in pixels.
func (r *Chromium) ScreenshotWidth(width int) *Chromium {
	return r.Param("width", strconv.Itoa(width))
}

// ScreenshotHeight sets the device screen height in pixels.
func (r *Chromium) ScreenshotHeight(height int) *Chromium {
	return r.Param("height", strconv.Itoa(height))
}

// ScreenshotClip defines whether to clip the screenshot according to the device dimensions.
func (r *Chromium) ScreenshotClip(clip bool) *Chromium {
	return r.Bool("clip", clip)
}

// ScreenshotFormat sets the image compression format.
func (r *Chromium) ScreenshotFormat(format string) *Chromium {
	return r.Param("format", format)
}

// ScreenshotQuality sets the compression quality from range 0 to 100 (jpeg only).
func (r *Chromium) ScreenshotQuality(quality int) *Chromium {
	return r.Param("quality", strconv.Itoa(quality))
}

// ScreenshotOmitBackground hides the default white background and allows generating screenshots with transparency.
func (r *Chromium) ScreenshotOmitBackground(omit bool) *Chromium {
	return r.Bool("omitBackground", omit)
}

// ScreenshotOptimizeForSpeed defines whether to optimize image encoding for speed, not for resulting size.
func (r *Chromium) ScreenshotOptimizeForSpeed(optimize bool) *Chromium {
	return r.Bool("optimizeForSpeed", optimize)
}

// WebhookURL sets the webhook URL and HTTP method for successful conversions.
func (r *Chromium) WebhookURL(url, method string) *Chromium {
	r.Gotenberg.WebhookURL(url, method)
	return r
}

// WebhookErrorURL sets the webhook URL and HTTP method for failed conversions.
func (r *Chromium) WebhookErrorURL(url, method string) *Chromium {
	r.Gotenberg.WebhookErrorURL(url, method)
	return r
}

// WebhookHeader adds a custom header to be sent with webhook requests.
// Multiple headers can be added by calling this method multiple times.
func (r *Chromium) WebhookHeader(key, value string) *Chromium {
	r.Gotenberg.WebhookHeader(key, value)
	return r
}

// DownloadFrom sets the downloadFrom parameter for downloading files from URLs.
// The data should be a slice of DownloadItem representing the download configuration.
func (r *Chromium) DownloadFrom(url string, headers map[string]string) *Chromium {
	r.Gotenberg.DownloadFrom(url, headers)
	return r
}

// OutputFilename sets the output filename for the generated PDF.
func (r *Chromium) OutputFilename(filename string) *Chromium {
	r.Gotenberg.OutputFilename(filename)
	return r
}

// Trace sets the request trace identifier for debugging and monitoring.
// If not set, Gotenberg will assign a unique UUID trace.
func (r *Chromium) Trace(trace string) *Chromium {
	r.Gotenberg.Trace(trace)
	return r
}

// PaperSize sets the paper size for the PDF using width and height in inches.
func (r *Chromium) PaperSize(width, height float64) *Chromium {
	return r.PaperWidth(width).PaperHeight(height)
}

// PaperSizeA4 sets the paper size to A4 format.
func (r *Chromium) PaperSizeA4() *Chromium {
	return r.PaperSize(8.27, 11.7)
}

// PaperSizeA6 sets the paper size to A6 format.
func (r *Chromium) PaperSizeA6() *Chromium {
	return r.PaperSize(4.13, 5.83)
}

// PaperSizeLetter sets the paper size to Letter format.
func (r *Chromium) PaperSizeLetter() *Chromium {
	return r.PaperSize(8.5, 11)
}

// Margins sets the page margins for the PDF in inches.
// Parameters are in order: top, right, bottom, left.
func (r *Chromium) Margins(top, right, bottom, left float64) *Chromium {
	return r.MarginTop(top).MarginRight(right).MarginBottom(bottom).MarginLeft(left)
}

// SinglePage sets whether to print the entire content in one single page.
func (r *Chromium) SinglePage() *Chromium {
	return r.Bool("singlePage", true)
}

// PaperWidth sets the paper width in inches.
func (r *Chromium) PaperWidth(value float64) *Chromium {
	return r.Float("paperWidth", value)
}

// PaperHeight sets the paper height in inches.
func (r *Chromium) PaperHeight(value float64) *Chromium {
	return r.Float("paperHeight", value)
}

// MarginTop sets the top margin in inches.
func (r *Chromium) MarginTop(value float64) *Chromium {
	return r.Float("marginTop", value)
}

// MarginBottom sets the bottom margin in inches.
func (r *Chromium) MarginBottom(value float64) *Chromium {
	return r.Float("marginBottom", value)
}

// MarginLeft sets the left margin in inches.
func (r *Chromium) MarginLeft(value float64) *Chromium {
	return r.Float("marginLeft", value)
}

// MarginRight sets the right margin in inches.
func (r *Chromium) MarginRight(value float64) *Chromium {
	return r.Float("marginRight", value)
}

// PreferCssPageSize sets whether to prefer page size as defined by CSS.
func (r *Chromium) PreferCssPageSize() *Chromium {
	return r.Bool("preferCssPageSize", true)
}

// GenerateDocumentOutline sets whether the document outline should be embedded into the PDF.
func (r *Chromium) GenerateDocumentOutline() *Chromium {
	return r.Bool("generateDocumentOutline", true)
}

// GenerateTaggedPdf sets whether to generate tagged (accessible) PDF.
func (r *Chromium) GenerateTaggedPdf() *Chromium {
	return r.Bool("generateTaggedPdf", true)
}

// PrintBackground sets whether to print the background graphics.
func (r *Chromium) PrintBackground() *Chromium {
	return r.Bool("printBackground", true)
}

// OmitBackground hides the default white background and allows generating PDFs with transparency.
func (r *Chromium) OmitBackground() *Chromium {
	return r.Bool("omitBackground", true)
}

// Landscape sets the paper orientation to landscape.
func (r *Chromium) Landscape() *Chromium {
	return r.Bool("landscape", true)
}

// Scale sets the scale of the page rendering.
func (r *Chromium) Scale(value float64) *Chromium {
	return r.Float("scale", value)
}

// NativePageRanges sets the page ranges to print, e.g., '1-5, 8, 11-13'.
func (r *Chromium) NativePageRanges(value string) *Chromium {
	return r.Param("nativePageRanges", value)
}

// WaitDelay sets the duration to wait when loading an HTML document before converting it into PDF.
func (r *Chromium) WaitDelay(value string) *Chromium {
	return r.Param("waitDelay", value)
}

// WaitForExpression sets the JavaScript expression to wait before converting an HTML document into PDF until it returns true.
func (r *Chromium) WaitForExpression(value string) *Chromium {
	return r.Param("waitForExpression", value)
}
