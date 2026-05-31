package gotenberg

import (
	"context"
	"io"
	"strconv"
	"time"
)

// Convert creates a request to convert Office documents to PDF.
func (r *LibreOffice) Convert(ctx context.Context) *LibreOffice {
	r.Req = r.HttpStream.Multipart(ctx, "/forms/libreoffice/convert")
	return r
}

// Send executes the conversion request and returns the response.
func (r *LibreOffice) Send() (*Response, error) {
	return r.Request.Send()
}

// Header adds an HTTP header to the conversion request.
func (r *LibreOffice) Header(key, value string) *LibreOffice {
	r.Request.Header(key, value)
	return r
}

// Param adds a form parameter to the conversion request.
func (r *LibreOffice) Param(key, value string) *LibreOffice {
	r.Request.Param(key, value)
	return r
}

// Bool adds a boolean form parameter to the conversion request.
func (r *LibreOffice) Bool(fieldName string, value bool) *LibreOffice {
	r.Request.Bool(fieldName, value)
	return r
}

// Float adds a float64 form parameter to the conversion request.
func (r *LibreOffice) Float(fieldName string, value float64) *LibreOffice {
	r.Request.Float(fieldName, value)
	return r
}

// File adds a file to the conversion request.
func (r *LibreOffice) File(filename string, content io.Reader) *LibreOffice {
	r.Request.File(filename, content)
	return r
}

// WebhookURL sets the webhook URL and HTTP method for successful conversions.
func (r *LibreOffice) WebhookURL(url, method string) *LibreOffice {
	r.Request.WebhookURL(url, method)
	return r
}

// WebhookErrorURL sets the webhook URL and HTTP method for failed conversions.
func (r *LibreOffice) WebhookErrorURL(url, method string) *LibreOffice {
	r.Request.WebhookErrorURL(url, method)
	return r
}

// WebhookEventsURL sets the webhook events URL for structured JSON event callbacks.
func (r *LibreOffice) WebhookEventsURL(url string) *LibreOffice {
	r.Request.WebhookEventsURL(url)
	return r
}

// WebhookHeader adds a custom header to be sent with webhook requests.
func (r *LibreOffice) WebhookHeader(key, value string) *LibreOffice {
	r.Request.WebhookHeader(key, value)
	return r
}

// DownloadFrom sets the downloadFrom parameter for downloading files from URLs.
func (r *LibreOffice) DownloadFrom(url string, headers map[string]string) *LibreOffice {
	r.Request.DownloadFrom(url, headers)
	return r
}

// OutputFilename sets the output filename for the generated PDF.
func (r *LibreOffice) OutputFilename(filename string) *LibreOffice {
	r.Request.OutputFilename(filename)
	return r
}

// Trace sets the request trace identifier for debugging and monitoring.
func (r *LibreOffice) Trace(trace string) *LibreOffice {
	r.Request.Trace(trace)
	return r
}

// Timeout sets a timeout for the request.
func (r *LibreOffice) Timeout(duration time.Duration) *LibreOffice {
	r.Request.Timeout(duration)
	return r
}

// Metadata sets the metadata for the operation.
func (r *LibreOffice) Metadata(key, value string) *LibreOffice {
	r.Request.Metadata(key, value)
	return r
}

// Password sets the password for opening the source file.
func (r *LibreOffice) Password(password string) *LibreOffice {
	return r.Param("password", password)
}

// Landscape sets the paper orientation to landscape.
func (r *LibreOffice) Landscape(landscape bool) *LibreOffice {
	return r.Bool("landscape", landscape)
}

// NativePageRanges sets the page ranges to print.
func (r *LibreOffice) NativePageRanges(ranges string) *LibreOffice {
	return r.Param("nativePageRanges", ranges)
}

// UpdateIndexes specifies whether to update the indexes before conversion.
func (r *LibreOffice) UpdateIndexes(update bool) *LibreOffice {
	return r.Bool("updateIndexes", update)
}

// ExportFormFields specifies whether form fields are exported as widgets.
func (r *LibreOffice) ExportFormFields(export bool) *LibreOffice {
	return r.Bool("exportFormFields", export)
}

// AllowDuplicateFieldNames specifies whether multiple form fields can have the same name.
func (r *LibreOffice) AllowDuplicateFieldNames(allow bool) *LibreOffice {
	return r.Bool("allowDuplicateFieldNames", allow)
}

// ExportBookmarks specifies if bookmarks are exported to PDF.
func (r *LibreOffice) ExportBookmarks(export bool) *LibreOffice {
	return r.Bool("exportBookmarks", export)
}

// ExportBookmarksToPdfDestination specifies bookmarks export to PDF destination.
func (r *LibreOffice) ExportBookmarksToPdfDestination(export bool) *LibreOffice {
	return r.Bool("exportBookmarksToPdfDestination", export)
}

// ExportPlaceholders exports placeholders fields visual markings only.
func (r *LibreOffice) ExportPlaceholders(export bool) *LibreOffice {
	return r.Bool("exportPlaceholders", export)
}

// ExportNotes specifies if notes are exported to PDF.
func (r *LibreOffice) ExportNotes(export bool) *LibreOffice {
	return r.Bool("exportNotes", export)
}

// ExportNotesPages specifies if notes pages are exported to PDF.
func (r *LibreOffice) ExportNotesPages(export bool) *LibreOffice {
	return r.Bool("exportNotesPages", export)
}

// ExportOnlyNotesPages specifies if only notes pages are exported.
func (r *LibreOffice) ExportOnlyNotesPages(export bool) *LibreOffice {
	return r.Bool("exportOnlyNotesPages", export)
}

// ExportNotesInMargin specifies if notes in margin are exported.
func (r *LibreOffice) ExportNotesInMargin(export bool) *LibreOffice {
	return r.Bool("exportNotesInMargin", export)
}

// ConvertOooTargetToPdfTarget converts OOo target to PDF target.
func (r *LibreOffice) ConvertOooTargetToPdfTarget(convert bool) *LibreOffice {
	return r.Bool("convertOooTargetToPdfTarget", convert)
}

// ExportLinksRelativeFsys exports relative filesystem links.
func (r *LibreOffice) ExportLinksRelativeFsys(export bool) *LibreOffice {
	return r.Bool("exportLinksRelativeFsys", export)
}

// ExportHiddenSlides exports hidden slides for Impress.
func (r *LibreOffice) ExportHiddenSlides(export bool) *LibreOffice {
	return r.Bool("exportHiddenSlides", export)
}

// SkipEmptyPages suppresses automatically inserted empty pages.
func (r *LibreOffice) SkipEmptyPages(skip bool) *LibreOffice {
	return r.Bool("skipEmptyPages", skip)
}

// AddOriginalDocumentAsStream adds original document as stream.
func (r *LibreOffice) AddOriginalDocumentAsStream(add bool) *LibreOffice {
	return r.Bool("addOriginalDocumentAsStream", add)
}

// SinglePageSheets puts every sheet on exactly one page.
func (r *LibreOffice) SinglePageSheets(single bool) *LibreOffice {
	return r.Bool("singlePageSheets", single)
}

// LosslessImageCompression specifies lossless compression for images.
func (r *LibreOffice) LosslessImageCompression(lossless bool) *LibreOffice {
	return r.Bool("losslessImageCompression", lossless)
}

// Quality sets the JPG export quality.
func (r *LibreOffice) Quality(quality int) *LibreOffice {
	return r.Param("quality", strconv.Itoa(quality))
}

// ReduceImageResolution reduces image resolution.
func (r *LibreOffice) ReduceImageResolution(reduce bool) *LibreOffice {
	return r.Bool("reduceImageResolution", reduce)
}

// MaxImageResolution sets the max image resolution in DPI.
func (r *LibreOffice) MaxImageResolution(resolution int) *LibreOffice {
	return r.Param("maxImageResolution", strconv.Itoa(resolution))
}

// Merge merges the resulting PDFs alphanumerically.
func (r *LibreOffice) Merge(merge bool) *LibreOffice {
	return r.Bool("merge", merge)
}

// SplitMode sets the split mode.
func (r *LibreOffice) SplitMode(mode string) *LibreOffice {
	return r.Param("splitMode", mode)
}

// SplitSpan sets the split span.
func (r *LibreOffice) SplitSpan(span string) *LibreOffice {
	return r.Param("splitSpan", span)
}

// SplitUnify specifies whether to unify split pages.
func (r *LibreOffice) SplitUnify(unify bool) *LibreOffice {
	return r.Bool("splitUnify", unify)
}

// PDFA converts to PDF/A format.
func (r *LibreOffice) PDFA(pdfa string) *LibreOffice {
	return r.Param("pdfa", pdfa)
}

// PDFUA enables PDF for Universal Access.
func (r *LibreOffice) PDFUA(pdfua bool) *LibreOffice {
	return r.Bool("pdfua", pdfua)
}

// NativeWatermarkText sets the text for LibreOffice's built-in watermark.
func (r *LibreOffice) NativeWatermarkText(text string) *LibreOffice {
	return r.Param("nativeWatermarkText", text)
}

// NativeWatermarkColor sets the color of the native watermark.
func (r *LibreOffice) NativeWatermarkColor(color string) *LibreOffice {
	return r.Param("nativeWatermarkColor", color)
}

// NativeWatermarkFontHeight sets the font height of the native watermark in points.
func (r *LibreOffice) NativeWatermarkFontHeight(height int) *LibreOffice {
	return r.Param("nativeWatermarkFontHeight", strconv.Itoa(height))
}

// NativeWatermarkRotateAngle sets the rotation angle of the native watermark in degrees.
func (r *LibreOffice) NativeWatermarkRotateAngle(angle int) *LibreOffice {
	return r.Param("nativeWatermarkRotateAngle", strconv.Itoa(angle))
}

// NativeWatermarkFontName sets the font name for the native watermark.
func (r *LibreOffice) NativeWatermarkFontName(name string) *LibreOffice {
	return r.Param("nativeWatermarkFontName", name)
}

// NativeTiledWatermarkText sets a tiled watermark text using LibreOffice's built-in rendering.
func (r *LibreOffice) NativeTiledWatermarkText(text string) *LibreOffice {
	return r.Param("nativeTiledWatermarkText", text)
}

// InitialView sets the initial view when the PDF is opened.
func (r *LibreOffice) InitialView(view string) *LibreOffice {
	return r.Param("initialView", view)
}

// InitialPage sets the initial page number displayed.
func (r *LibreOffice) InitialPage(page int) *LibreOffice {
	return r.Param("initialPage", strconv.Itoa(page))
}

// Magnification sets the magnification mode when the PDF is opened.
func (r *LibreOffice) Magnification(mag string) *LibreOffice {
	return r.Param("magnification", mag)
}

// Zoom sets the zoom percentage when the PDF is opened.
func (r *LibreOffice) Zoom(zoom int) *LibreOffice {
	return r.Param("zoom", strconv.Itoa(zoom))
}

// PageLayout sets the page layout when the PDF is opened.
func (r *LibreOffice) PageLayout(layout string) *LibreOffice {
	return r.Param("pageLayout", layout)
}

// FirstPageOnLeft sets whether the first page is shown on the left.
func (r *LibreOffice) FirstPageOnLeft(v bool) *LibreOffice {
	return r.Bool("firstPageOnLeft", v)
}

// ResizeWindowToInitialPage sets whether to resize the viewer window.
func (r *LibreOffice) ResizeWindowToInitialPage(v bool) *LibreOffice {
	return r.Bool("resizeWindowToInitialPage", v)
}

// CenterWindow sets whether to center the viewer window.
func (r *LibreOffice) CenterWindow(v bool) *LibreOffice {
	return r.Bool("centerWindow", v)
}

// OpenInFullScreenMode sets whether to open the PDF in full-screen mode.
func (r *LibreOffice) OpenInFullScreenMode(v bool) *LibreOffice {
	return r.Bool("openInFullScreenMode", v)
}

// DisplayPDFDocumentTitle sets whether to display the document title.
func (r *LibreOffice) DisplayPDFDocumentTitle(v bool) *LibreOffice {
	return r.Bool("displayPDFDocumentTitle", v)
}

// HideViewerMenubar sets whether to hide the viewer's menu bar.
func (r *LibreOffice) HideViewerMenubar(v bool) *LibreOffice {
	return r.Bool("hideViewerMenubar", v)
}

// HideViewerToolbar sets whether to hide the viewer's toolbar.
func (r *LibreOffice) HideViewerToolbar(v bool) *LibreOffice {
	return r.Bool("hideViewerToolbar", v)
}

// HideViewerWindowControls sets whether to hide the viewer's window controls.
func (r *LibreOffice) HideViewerWindowControls(v bool) *LibreOffice {
	return r.Bool("hideViewerWindowControls", v)
}

// UseTransitionEffects sets whether to use slide transition effects in Impress presentations.
func (r *LibreOffice) UseTransitionEffects(v bool) *LibreOffice {
	return r.Bool("useTransitionEffects", v)
}

// OpenBookmarkLevels sets the number of bookmark levels to expand.
func (r *LibreOffice) OpenBookmarkLevels(levels int) *LibreOffice {
	return r.Param("openBookmarkLevels", strconv.Itoa(levels))
}

// EmbedsMetadata sets per-file metadata for embedded files as a JSON object.
func (r *LibreOffice) EmbedsMetadata(metadataJSON string) *LibreOffice {
	return r.Param("embedsMetadata", metadataJSON)
}

// Flatten flattens the resulting PDF.
func (r *LibreOffice) Flatten(flatten bool) *LibreOffice {
	return r.Bool("flatten", flatten)
}
