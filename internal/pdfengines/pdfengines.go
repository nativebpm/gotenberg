// Package pdfengines provides a client for the Gotenberg PDF Engines service.
// It offers a convenient API for converting PDFs to PDF/A & PDF/UA, reading/writing metadata, merging, splitting, and flattening PDFs.
package pdfengines

import (
	"context"
	"io"
	"time"

	"github.com/nativebpm/gotenberg/internal/gotenberg"
	"github.com/nativebpm/connectors/httpstream"
)

// PDFEngines represents a Gotenberg conversion request builder.
// It wraps the underlying multipart request and provides PDF Engines-specific methods.
type PDFEngines struct {
	*gotenberg.Gotenberg
}

func NewPDFEngines(client *httpstream.Client) *PDFEngines {
	return &PDFEngines{
		Gotenberg: gotenberg.NewGotenberg(client),
	}
}

// Convert creates a request to convert PDFs to PDF/A & PDF/UA.
func (r *PDFEngines) Convert(ctx context.Context) *PDFEngines {
	r.Req = r.HttpStream.Multipart(ctx, "/forms/pdfengines/convert")
	return r
}

// MetadataRead creates a request to read metadata from PDFs.
func (r *PDFEngines) MetadataRead(ctx context.Context) *PDFEngines {
	r.Req = r.HttpStream.Multipart(ctx, "/forms/pdfengines/metadata/read")
	return r
}

// MetadataWrite creates a request to write metadata to PDFs.
func (r *PDFEngines) MetadataWrite(ctx context.Context) *PDFEngines {
	r.Req = r.HttpStream.Multipart(ctx, "/forms/pdfengines/metadata/write")
	return r
}

// Merge creates a request to merge PDFs.
func (r *PDFEngines) Merge(ctx context.Context) *PDFEngines {
	r.Req = r.HttpStream.Multipart(ctx, "/forms/pdfengines/merge")
	return r
}

// Split creates a request to split PDFs.
func (r *PDFEngines) Split(ctx context.Context) *PDFEngines {
	r.Req = r.HttpStream.Multipart(ctx, "/forms/pdfengines/split")
	return r
}

// Flatten creates a request to flatten PDFs.
func (r *PDFEngines) Flatten(ctx context.Context) *PDFEngines {
	r.Req = r.HttpStream.Multipart(ctx, "/forms/pdfengines/flatten")
	return r
}

// Send executes the request and returns the response.
// Returns an error if the request fails.
func (r *PDFEngines) Send() (*gotenberg.Response, error) {
	return r.Gotenberg.Send()
}

// Header adds a header to the request.
func (r *PDFEngines) Header(key, value string) *PDFEngines {
	r.Gotenberg.Header(key, value)
	return r
}

// Param adds a form parameter to the request.
func (r *PDFEngines) Param(key, value string) *PDFEngines {
	r.Gotenberg.Param(key, value)
	return r
}

// Bool adds a boolean form parameter to the request.
func (r *PDFEngines) Bool(fieldName string, value bool) *PDFEngines {
	r.Gotenberg.Bool(fieldName, value)
	return r
}

// File adds a file to the request.
func (r *PDFEngines) File(filename string, content io.Reader) *PDFEngines {
	r.Gotenberg.File("files", filename, content)
	return r
}

// WebhookURL sets the webhook URL and HTTP method for successful operations.
func (r *PDFEngines) WebhookURL(url, method string) *PDFEngines {
	r.Gotenberg.WebhookURL(url, method)
	return r
}

// WebhookErrorURL sets the webhook URL and HTTP method for failed operations.
func (r *PDFEngines) WebhookErrorURL(url, method string) *PDFEngines {
	r.Gotenberg.WebhookErrorURL(url, method)
	return r
}

// WebhookHeader adds a custom header to be sent with webhook requests.
// Multiple headers can be added by calling this method multiple times.
func (r *PDFEngines) WebhookHeader(key, value string) *PDFEngines {
	r.Gotenberg.WebhookHeader(key, value)
	return r
}

// DownloadFrom sets the downloadFrom parameter for downloading files from URLs.
// The data should be a slice of DownloadItem representing the download configuration.
func (r *PDFEngines) DownloadFrom(url string, headers map[string]string) *PDFEngines {
	r.Gotenberg.DownloadFrom(url, headers)
	return r
}

// OutputFilename sets the output filename.
func (r *PDFEngines) OutputFilename(filename string) *PDFEngines {
	r.Gotenberg.OutputFilename(filename)
	return r
}

// Trace sets the request trace identifier for debugging and monitoring.
// If not set, Gotenberg will assign a unique UUID trace.
func (r *PDFEngines) Trace(trace string) *PDFEngines {
	r.Gotenberg.Trace(trace)
	return r
}

// Timeout sets a timeout for the request.
func (r *PDFEngines) Timeout(duration time.Duration) *PDFEngines {
	r.Gotenberg.Timeout(duration)
	return r
}

// PDFA converts to PDF/A format.
func (r *PDFEngines) PDFA(pdfa string) *PDFEngines {
	r.Gotenberg.Param("pdfa", pdfa)
	return r
}

// PDFUA enables PDF for Universal Access.
func (r *PDFEngines) PDFUA(pdfua bool) *PDFEngines {
	r.Gotenberg.Bool("pdfua", pdfua)
	return r
}

// Metadata sets the metadata for the PDF.
func (r *PDFEngines) Metadata(key, value string) *PDFEngines {
	r.Gotenberg.Metadata(key, value)
	return r
}

// SplitMode sets the split mode.
func (r *PDFEngines) SplitMode(mode string) *PDFEngines {
	r.Gotenberg.Param("splitMode", mode)
	return r
}

// SplitSpan sets the split span.
func (r *PDFEngines) SplitSpan(span string) *PDFEngines {
	r.Gotenberg.Param("splitSpan", span)
	return r
}

// SplitUnify specifies whether to unify split pages.
func (r *PDFEngines) SplitUnify(unify bool) *PDFEngines {
	r.Gotenberg.Bool("splitUnify", unify)
	return r
}

// Flatten sets the flatten flag.
func (r *PDFEngines) FlattenPDF(flatten bool) *PDFEngines {
	r.Gotenberg.Bool("flatten", flatten)
	return r
}
