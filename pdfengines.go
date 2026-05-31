package gotenberg

import (
	"context"
	"io"
	"strconv"
	"time"
)

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

// Watermark creates a request to apply a watermark behind page content.
func (r *PDFEngines) Watermark(ctx context.Context) *PDFEngines {
	r.Req = r.HttpStream.Multipart(ctx, "/forms/pdfengines/watermark")
	return r
}

// Stamp creates a request to apply a stamp on top of page content.
func (r *PDFEngines) Stamp(ctx context.Context) *PDFEngines {
	r.Req = r.HttpStream.Multipart(ctx, "/forms/pdfengines/stamp")
	return r
}

// Rotate creates a request to rotate PDF pages by 90°, 180°, or 270°.
func (r *PDFEngines) Rotate(ctx context.Context) *PDFEngines {
	r.Req = r.HttpStream.Multipart(ctx, "/forms/pdfengines/rotate")
	return r
}

// BookmarksRead creates a request to read the bookmark outline from PDF files as JSON.
func (r *PDFEngines) BookmarksRead(ctx context.Context) *PDFEngines {
	r.Req = r.HttpStream.Multipart(ctx, "/forms/pdfengines/bookmarks/read")
	return r
}

// BookmarksWrite creates a request to write bookmarks to PDF files.
func (r *PDFEngines) BookmarksWrite(ctx context.Context) *PDFEngines {
	r.Req = r.HttpStream.Multipart(ctx, "/forms/pdfengines/bookmarks/write")
	return r
}

// Send executes the request and returns the response.
func (r *PDFEngines) Send() (*Response, error) {
	return r.Request.Send()
}

// Header adds an HTTP header to the request.
func (r *PDFEngines) Header(key, value string) *PDFEngines {
	r.Request.Header(key, value)
	return r
}

// Param adds a form parameter to the request.
func (r *PDFEngines) Param(key, value string) *PDFEngines {
	r.Request.Param(key, value)
	return r
}

// Bool adds a boolean form parameter to the request.
func (r *PDFEngines) Bool(fieldName string, value bool) *PDFEngines {
	r.Request.Bool(fieldName, value)
	return r
}

// Float adds a float64 form parameter to the request.
func (r *PDFEngines) Float(fieldName string, value float64) *PDFEngines {
	r.Request.Float(fieldName, value)
	return r
}

// File adds a file to the request.
func (r *PDFEngines) File(filename string, content io.Reader) *PDFEngines {
	r.Request.File(filename, content)
	return r
}

// WebhookURL sets the webhook URL and HTTP method for successful operations.
func (r *PDFEngines) WebhookURL(url, method string) *PDFEngines {
	r.Request.WebhookURL(url, method)
	return r
}

// WebhookErrorURL sets the webhook URL and HTTP method for failed operations.
func (r *PDFEngines) WebhookErrorURL(url, method string) *PDFEngines {
	r.Request.WebhookErrorURL(url, method)
	return r
}

// WebhookEventsURL sets the webhook events URL for structured JSON event callbacks.
func (r *PDFEngines) WebhookEventsURL(url string) *PDFEngines {
	r.Request.WebhookEventsURL(url)
	return r
}

// WebhookHeader adds a custom header to be sent with webhook requests.
func (r *PDFEngines) WebhookHeader(key, value string) *PDFEngines {
	r.Request.WebhookHeader(key, value)
	return r
}

// DownloadFrom sets the downloadFrom parameter for downloading files from URLs.
func (r *PDFEngines) DownloadFrom(url string, headers map[string]string) *PDFEngines {
	r.Request.DownloadFrom(url, headers)
	return r
}

// OutputFilename sets the output filename.
func (r *PDFEngines) OutputFilename(filename string) *PDFEngines {
	r.Request.OutputFilename(filename)
	return r
}

// Trace sets the request trace identifier for debugging and monitoring.
func (r *PDFEngines) Trace(trace string) *PDFEngines {
	r.Request.Trace(trace)
	return r
}

// Timeout sets a timeout for the request.
func (r *PDFEngines) Timeout(duration time.Duration) *PDFEngines {
	r.Request.Timeout(duration)
	return r
}

// Metadata sets the metadata for the PDF.
func (r *PDFEngines) Metadata(key, value string) *PDFEngines {
	r.Request.Metadata(key, value)
	return r
}

// PDFA converts to PDF/A format.
func (r *PDFEngines) PDFA(pdfa string) *PDFEngines {
	return r.Param("pdfa", pdfa)
}

// PDFUA enables PDF for Universal Access.
func (r *PDFEngines) PDFUA(pdfua bool) *PDFEngines {
	return r.Bool("pdfua", pdfua)
}

// SplitMode sets the split mode.
func (r *PDFEngines) SplitMode(mode string) *PDFEngines {
	return r.Param("splitMode", mode)
}

// SplitSpan sets the split span.
func (r *PDFEngines) SplitSpan(span string) *PDFEngines {
	return r.Param("splitSpan", span)
}

// SplitUnify specifies whether to unify split pages.
func (r *PDFEngines) SplitUnify(unify bool) *PDFEngines {
	return r.Bool("splitUnify", unify)
}

// FlattenPDF sets the flatten flag.
func (r *PDFEngines) FlattenPDF(flatten bool) *PDFEngines {
	return r.Bool("flatten", flatten)
}

// RotateAngle sets the rotation angle for pages.
func (r *PDFEngines) RotateAngle(angle int) *PDFEngines {
	return r.Param("rotateAngle", strconv.Itoa(angle))
}

// RotatePages sets the page selection for rotation.
func (r *PDFEngines) RotatePages(pages string) *PDFEngines {
	return r.Param("rotatePages", pages)
}

// Bookmarks sets the bookmarks JSON.
func (r *PDFEngines) Bookmarks(json string) *PDFEngines {
	return r.Param("bookmarks", json)
}

// AutoIndexBookmarks extracts and reindexes existing bookmarks from input files during merge.
func (r *PDFEngines) AutoIndexBookmarks(v bool) *PDFEngines {
	return r.Bool("autoIndexBookmarks", v)
}

// WatermarkFile adds a watermark source file (image or PDF) to the request.
func (r *PDFEngines) WatermarkFile(filename string, content io.Reader) *PDFEngines {
	r.Request.file("watermark", filename, content)
	return r
}

// StampFile adds a stamp source file (image or PDF) to the request.
func (r *PDFEngines) StampFile(filename string, content io.Reader) *PDFEngines {
	r.Request.file("stamp", filename, content)
	return r
}

// EmbedsMetadata sets per-file metadata.
func (r *PDFEngines) EmbedsMetadata(metadataJSON string) *PDFEngines {
	return r.Param("embedsMetadata", metadataJSON)
}
