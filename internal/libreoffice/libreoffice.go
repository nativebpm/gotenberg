// Package libreoffice provides a client for the Gotenberg LibreOffice service.
// It offers a convenient API for converting Office documents to PDF documents.
package libreoffice

import (
	"context"
	"io"
	"strconv"
	"time"

	"github.com/nativebpm/gotenberg/internal/gotenberg"
	"github.com/nativebpm/connectors/streamhttp"
)

// LibreOffice represents a Gotenberg conversion request builder.
type LibreOffice struct {
	*gotenberg.Gotenberg
}

func NewLibreOffice(client *streamhttp.Client) *LibreOffice {
	return &LibreOffice{
		Gotenberg: gotenberg.NewGotenberg(client),
	}
}

// Convert creates a request to convert Office documents to PDF.
// The files parameter should contain the Office documents to be converted.
func (r *LibreOffice) Convert(ctx context.Context) *LibreOffice {
	r.Req = r.Client.Multipart(ctx, "/forms/libreoffice/convert")
	return r
}

// Send executes the conversion request and returns the response.
// Returns an error if the request fails or the conversion cannot be completed.
func (r *LibreOffice) Send() (*gotenberg.Response, error) {
	return r.Gotenberg.Send()
}

// Header adds a header to the conversion request.
func (r *LibreOffice) Header(key, value string) *LibreOffice {
	r.Gotenberg.Header(key, value)
	return r
}

// Param adds a form parameter to the conversion request.
func (r *LibreOffice) Param(key, value string) *LibreOffice {
	r.Gotenberg.Param(key, value)
	return r
}

// Bool adds a boolean form parameter to the conversion request.
func (r *LibreOffice) Bool(fieldName string, value bool) *LibreOffice {
	r.Gotenberg.Bool(fieldName, value)
	return r
}

// Float adds a float64 form parameter to the conversion request.
func (r *LibreOffice) Float(fieldName string, value float64) *LibreOffice {
	r.Gotenberg.Float(fieldName, value)
	return r
}

// File adds a file to the conversion request.
func (r *LibreOffice) File(filename string, content io.Reader) *LibreOffice {
	r.Gotenberg.File("files", filename, content)
	return r
}

// WebhookURL sets the webhook URL and HTTP method for successful conversions.
func (r *LibreOffice) WebhookURL(url, method string) *LibreOffice {
	r.Gotenberg.WebhookURL(url, method)
	return r
}

// WebhookErrorURL sets the webhook URL and HTTP method for failed conversions.
func (r *LibreOffice) WebhookErrorURL(url, method string) *LibreOffice {
	r.Gotenberg.WebhookErrorURL(url, method)
	return r
}

// WebhookHeader adds a custom header to be sent with webhook requests.
// Multiple headers can be added by calling this method multiple times.
func (r *LibreOffice) WebhookHeader(key, value string) *LibreOffice {
	r.Gotenberg.WebhookHeader(key, value)
	return r
}

// DownloadFrom sets the downloadFrom parameter for downloading files from URLs.
// The data should be a slice of DownloadItem representing the download configuration.
func (r *LibreOffice) DownloadFrom(url string, headers map[string]string) *LibreOffice {
	r.Gotenberg.DownloadFrom(url, headers)
	return r
}

// OutputFilename sets the output filename for the generated PDF.
func (r *LibreOffice) OutputFilename(filename string) *LibreOffice {
	r.Gotenberg.OutputFilename(filename)
	return r
}

// Trace sets the request trace identifier for debugging and monitoring.
// If not set, Gotenberg will assign a unique UUID trace.
func (r *LibreOffice) Trace(trace string) *LibreOffice {
	r.Gotenberg.Trace(trace)
	return r
}

// Timeout sets a timeout for the request.
func (r *LibreOffice) Timeout(duration time.Duration) *LibreOffice {
	r.Gotenberg.Timeout(duration)
	return r
}

// Password sets the password for opening the source file.
func (r *LibreOffice) Password(password string) *LibreOffice {
	r.Gotenberg.Param("password", password)
	return r
}

// Landscape sets the paper orientation to landscape.
func (r *LibreOffice) Landscape(landscape bool) *LibreOffice {
	r.Gotenberg.Bool("landscape", landscape)
	return r
}

// NativePageRanges sets the page ranges to print.
func (r *LibreOffice) NativePageRanges(ranges string) *LibreOffice {
	r.Gotenberg.Param("nativePageRanges", ranges)
	return r
}

// UpdateIndexes specifies whether to update the indexes before conversion.
func (r *LibreOffice) UpdateIndexes(update bool) *LibreOffice {
	r.Gotenberg.Bool("updateIndexes", update)
	return r
}

// ExportFormFields specifies whether form fields are exported as widgets.
func (r *LibreOffice) ExportFormFields(export bool) *LibreOffice {
	r.Gotenberg.Bool("exportFormFields", export)
	return r
}

// AllowDuplicateFieldNames specifies whether multiple form fields can have the same name.
func (r *LibreOffice) AllowDuplicateFieldNames(allow bool) *LibreOffice {
	r.Gotenberg.Bool("allowDuplicateFieldNames", allow)
	return r
}

// ExportBookmarks specifies if bookmarks are exported to PDF.
func (r *LibreOffice) ExportBookmarks(export bool) *LibreOffice {
	r.Gotenberg.Bool("exportBookmarks", export)
	return r
}

// ExportBookmarksToPdfDestination specifies bookmarks export to PDF destination.
func (r *LibreOffice) ExportBookmarksToPdfDestination(export bool) *LibreOffice {
	r.Gotenberg.Bool("exportBookmarksToPdfDestination", export)
	return r
}

// ExportPlaceholders exports placeholders fields visual markings only.
func (r *LibreOffice) ExportPlaceholders(export bool) *LibreOffice {
	r.Gotenberg.Bool("exportPlaceholders", export)
	return r
}

// ExportNotes specifies if notes are exported to PDF.
func (r *LibreOffice) ExportNotes(export bool) *LibreOffice {
	r.Gotenberg.Bool("exportNotes", export)
	return r
}

// ExportNotesPages specifies if notes pages are exported to PDF.
func (r *LibreOffice) ExportNotesPages(export bool) *LibreOffice {
	r.Gotenberg.Bool("exportNotesPages", export)
	return r
}

// ExportOnlyNotesPages specifies if only notes pages are exported.
func (r *LibreOffice) ExportOnlyNotesPages(export bool) *LibreOffice {
	r.Gotenberg.Bool("exportOnlyNotesPages", export)
	return r
}

// ExportNotesInMargin specifies if notes in margin are exported.
func (r *LibreOffice) ExportNotesInMargin(export bool) *LibreOffice {
	r.Gotenberg.Bool("exportNotesInMargin", export)
	return r
}

// ConvertOooTargetToPdfTarget converts OOo target to PDF target.
func (r *LibreOffice) ConvertOooTargetToPdfTarget(convert bool) *LibreOffice {
	r.Gotenberg.Bool("convertOooTargetToPdfTarget", convert)
	return r
}

// ExportLinksRelativeFsys exports relative filesystem links.
func (r *LibreOffice) ExportLinksRelativeFsys(export bool) *LibreOffice {
	r.Gotenberg.Bool("exportLinksRelativeFsys", export)
	return r
}

// ExportHiddenSlides exports hidden slides for Impress.
func (r *LibreOffice) ExportHiddenSlides(export bool) *LibreOffice {
	r.Gotenberg.Bool("exportHiddenSlides", export)
	return r
}

// SkipEmptyPages suppresses automatically inserted empty pages.
func (r *LibreOffice) SkipEmptyPages(skip bool) *LibreOffice {
	r.Gotenberg.Bool("skipEmptyPages", skip)
	return r
}

// AddOriginalDocumentAsStream adds original document as stream.
func (r *LibreOffice) AddOriginalDocumentAsStream(add bool) *LibreOffice {
	r.Gotenberg.Bool("addOriginalDocumentAsStream", add)
	return r
}

// SinglePageSheets puts every sheet on exactly one page.
func (r *LibreOffice) SinglePageSheets(single bool) *LibreOffice {
	r.Gotenberg.Bool("singlePageSheets", single)
	return r
}

// LosslessImageCompression specifies lossless compression for images.
func (r *LibreOffice) LosslessImageCompression(lossless bool) *LibreOffice {
	r.Gotenberg.Bool("losslessImageCompression", lossless)
	return r
}

// Quality sets the JPG export quality.
func (r *LibreOffice) Quality(quality int) *LibreOffice {
	r.Gotenberg.Param("quality", strconv.Itoa(quality))
	return r
}

// ReduceImageResolution reduces image resolution.
func (r *LibreOffice) ReduceImageResolution(reduce bool) *LibreOffice {
	r.Gotenberg.Bool("reduceImageResolution", reduce)
	return r
}

// MaxImageResolution sets the max image resolution in DPI.
func (r *LibreOffice) MaxImageResolution(resolution int) *LibreOffice {
	r.Gotenberg.Param("maxImageResolution", strconv.Itoa(resolution))
	return r
}

// Merge merges the resulting PDFs alphanumerically.
func (r *LibreOffice) Merge(merge bool) *LibreOffice {
	r.Gotenberg.Bool("merge", merge)
	return r
}

// SplitMode sets the split mode.
func (r *LibreOffice) SplitMode(mode string) *LibreOffice {
	r.Gotenberg.Param("splitMode", mode)
	return r
}

// SplitSpan sets the split span.
func (r *LibreOffice) SplitSpan(span string) *LibreOffice {
	r.Gotenberg.Param("splitSpan", span)
	return r
}

// SplitUnify specifies whether to unify split pages.
func (r *LibreOffice) SplitUnify(unify bool) *LibreOffice {
	r.Gotenberg.Bool("splitUnify", unify)
	return r
}

// PDFA converts to PDF/A format.
func (r *LibreOffice) PDFA(pdfa string) *LibreOffice {
	r.Gotenberg.Param("pdfa", pdfa)
	return r
}

// PDFUA enables PDF for Universal Access.
func (r *LibreOffice) PDFUA(pdfua bool) *LibreOffice {
	r.Gotenberg.Bool("pdfua", pdfua)
	return r
}

// Metadata sets the metadata for the PDF.
func (r *LibreOffice) Metadata(key, value string) *LibreOffice {
	r.Gotenberg.Metadata(key, value)
	return r
}

// Flatten flattens the resulting PDF.
func (r *LibreOffice) Flatten(flatten bool) *LibreOffice {
	r.Gotenberg.Bool("flatten", flatten)
	return r
}
