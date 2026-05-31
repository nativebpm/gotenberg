// Package gotenberg provides a client for the Gotenberg service.
// It offers a convenient API for document conversion using various engines.
package gotenberg

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/nativebpm/httpstream"
)

// downloadFrom represents the configuration for downloading files from external URLs.
type downloadFrom struct {
	URL     string            `json:"url"`
	Headers map[string]string `json:"extraHttpHeaders,omitempty"`
}

// Response represents a Gotenberg conversion response.
// It wraps the HTTP response and provides access to the Gotenberg trace header.
type Response struct {
	*http.Response
	GotenbergTrace string
}

// Request represents the base request builder carrying parameters and HTTP payload configurations.
type Request struct {
	HttpStream *httpstream.Client
	Req        *httpstream.Multipart
	Wh         map[string]string
	Meta       map[string]string
	Df         []downloadFrom
}

// Chromium represents a request builder specifically for Chromium-based PDF and screenshot conversions.
type Chromium struct {
	*Request
}

// LibreOffice represents a request builder specifically for Office document to PDF conversions.
type LibreOffice struct {
	*Request
}

// PDFEngines represents a request builder specifically for PDF engines actions (merge, split, flatten, bookmarks).
type PDFEngines struct {
	*Request
}

// Client is a Gotenberg HTTP client that wraps the base HTTP client
// with Gotenberg-specific functionality for document conversion.
type Client struct {
	HttpStream *httpstream.Client
}

// NewClient creates a new Gotenberg client with the given HTTP client and base URL.
// Returns an error if the base URL is invalid.
func NewClient(httpClient *http.Client, baseURL string) (*Client, error) {
	client, err := httpstream.NewClient(httpClient, baseURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		HttpStream: client,
	}, nil
}

// Use applies HTTP round-tripper middlewares to the client.
func (c *Client) Use(middleware func(http.RoundTripper) http.RoundTripper) *Client {
	c.HttpStream = c.HttpStream.Use(middleware)
	return c
}

// Chromium returns a Request builder configured for Chromium operations.
func (c *Client) Chromium() *Chromium {
	return &Chromium{
		Request: &Request{HttpStream: c.HttpStream},
	}
}

// LibreOffice returns a Request builder configured for LibreOffice operations.
func (c *Client) LibreOffice() *LibreOffice {
	return &LibreOffice{
		Request: &Request{HttpStream: c.HttpStream},
	}
}

// PDFEngines returns a Request builder configured for PDF Engines operations.
func (c *Client) PDFEngines() *PDFEngines {
	return &PDFEngines{
		Request: &Request{HttpStream: c.HttpStream},
	}
}

// Send executes the request and returns the response.
// It handles common fields like webhook headers, downloadFrom, and metadata.
func (r *Request) Send() (*Response, error) {
	// Dynamically marshal fields if they are set
	for _, item := range []struct {
		cond     bool
		isHeader bool
		key      string
		val      any
	}{
		{
			cond:     len(r.Wh) > 0,
			isHeader: true,
			key:      "Gotenberg-Webhook-Extra-Http-Headers",
			val:      r.Wh,
		},
		{
			cond:     len(r.Df) > 0,
			isHeader: false,
			key:      "downloadFrom",
			val:      r.Df,
		},
		{
			cond:     len(r.Meta) > 0,
			isHeader: false,
			key:      "metadata",
			val:      r.Meta,
		},
	} {
		if item.cond {
			b, err := json.Marshal(item.val)
			if err != nil {
				return nil, err
			}
			if item.isHeader {
				r.Req.Header(item.key, string(b))
			} else {
				r.Req.Param(item.key, string(b))
			}
		}
	}

	resp, err := r.Req.Send()
	if err != nil {
		return nil, err
	}

	return &Response{
		Response:       resp,
		GotenbergTrace: resp.Header.Get("Gotenberg-Trace"),
	}, nil
}

// Header adds an HTTP header to the request.
func (r *Request) Header(key, value string) *Request {
	r.Req.Header(key, value)
	return r
}

// Param adds a form parameter to the request.
func (r *Request) Param(key, value string) *Request {
	r.Req.Param(key, value)
	return r
}

// Bool adds a boolean form parameter to the request.
func (r *Request) Bool(fieldName string, value bool) *Request {
	r.Req.Bool(fieldName, value)
	return r
}

// Float adds a float64 form parameter to the request.
func (r *Request) Float(fieldName string, value float64) *Request {
	r.Req.Float(fieldName, value)
	return r
}

// file adds a file to the request with a custom field name.
func (r *Request) file(fieldName, filename string, content io.Reader) *Request {
	r.Req.File(fieldName, filename, content)
	return r
}

// File adds a file to the conversion request.
func (r *Request) File(filename string, content io.Reader) *Request {
	return r.file("files", filename, content)
}

// WebhookURL sets the webhook URL and HTTP method for successful operations.
func (r *Request) WebhookURL(url, method string) *Request {
	r.Req.Header("Gotenberg-Webhook-Url", url).
		Header("Gotenberg-Webhook-Method", method)
	return r
}

// WebhookErrorURL sets the webhook URL and HTTP method for failed operations.
func (r *Request) WebhookErrorURL(url, method string) *Request {
	r.Req.Header("Gotenberg-Webhook-Error-Url", url).
		Header("Gotenberg-Webhook-Error-Method", method)
	return r
}

// WebhookEventsURL sets the webhook events URL for structured JSON event callbacks.
func (r *Request) WebhookEventsURL(url string) *Request {
	r.Req.Header("Gotenberg-Webhook-Events-Url", url)
	return r
}

// WebhookHeader adds a custom header to be sent with webhook requests.
func (r *Request) WebhookHeader(key, value string) *Request {
	if r.Wh == nil {
		r.Wh = make(map[string]string)
	}
	r.Wh[key] = value
	return r
}

// DownloadFrom sets the download from configuration.
func (r *Request) DownloadFrom(url string, headers map[string]string) *Request {
	r.Df = append(r.Df, downloadFrom{URL: url, Headers: headers})
	return r
}

// OutputFilename sets the output filename.
func (r *Request) OutputFilename(filename string) *Request {
	r.Req.Header("Gotenberg-Output-Filename", filename)
	return r
}

// Trace sets the request trace identifier for debugging and monitoring.
func (r *Request) Trace(trace string) *Request {
	r.Req.Header("Gotenberg-Trace", trace)
	return r
}

// Timeout sets a timeout for the request.
func (r *Request) Timeout(duration time.Duration) *Request {
	r.Req.Timeout(duration)
	return r
}

// Metadata sets the metadata for the operation.
func (r *Request) Metadata(key, value string) *Request {
	if r.Meta == nil {
		r.Meta = make(map[string]string)
	}
	r.Meta[key] = value
	return r
}
