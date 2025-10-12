// Package common provides shared types and constants for Gotenberg client modules.
package gotenberg

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/nativebpm/connectors/httpstream"
)

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

// Gotenberg provides common functionality for building Gotenberg requests.
// It wraps the underlying multipart request and provides shared methods.
type Gotenberg struct {
	HttpStream *httpstream.Client
	Req        *httpstream.Multipart
	Wh         map[string]string
	Meta       map[string]string
	Df         []downloadFrom
}

// NewGotenberg creates a new RequestBuilder with the given client.
func NewGotenberg(httpStream *httpstream.Client) *Gotenberg {
	return &Gotenberg{
		HttpStream: httpStream,
	}
}

// Timeout sets a timeout for the request.
func (r *Gotenberg) Timeout(duration time.Duration) *Gotenberg {
	r.Req.Timeout(duration)
	return r
}

// Send executes the request and returns the response.
// It handles common fields like webhook headers, downloadFrom, and metadata.
func (r *Gotenberg) Send() (*Response, error) {
	if len(r.Wh) > 0 {
		webhookHeaders, err := json.Marshal(r.Wh)
		if err != nil {
			return nil, err
		}
		r.Req.Header("Gotenberg-Webhook-Extra-Http-Headers", string(webhookHeaders))
	}

	if len(r.Df) > 0 {
		downloadFrom, err := json.Marshal(r.Df)
		if err != nil {
			return nil, err
		}
		r.Req.Param("downloadFrom", string(downloadFrom))
	}

	if len(r.Meta) > 0 {
		metadata, err := json.Marshal(r.Meta)
		if err != nil {
			return nil, err
		}
		r.Req.Param("metadata", string(metadata))
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

// Header adds a header to the request.
func (r *Gotenberg) Header(key, value string) *Gotenberg {
	r.Req.Header(key, value)
	return r
}

// Param adds a form parameter to the request.
func (r *Gotenberg) Param(key, value string) *Gotenberg {
	r.Req.Param(key, value)
	return r
}

// Bool adds a boolean form parameter to the request.
func (r *Gotenberg) Bool(fieldName string, value bool) *Gotenberg {
	r.Req.Bool(fieldName, value)
	return r
}

// Float adds a float64 form parameter to the request.
func (r *Gotenberg) Float(fieldName string, value float64) *Gotenberg {
	r.Req.Float(fieldName, value)
	return r
}

// File adds a file to the request.
func (r *Gotenberg) File(fieldName, filename string, content io.Reader) *Gotenberg {
	r.Req.File(fieldName, filename, content)
	return r
}

// WebhookURL sets the webhook URL and HTTP method for successful operations.
func (r *Gotenberg) WebhookURL(url, method string) *Gotenberg {
	r.Req.Header("Gotenberg-Webhook-Url", url).
		Header("Gotenberg-Webhook-Method", method)
	return r
}

// WebhookErrorURL sets the webhook URL and HTTP method for failed operations.
func (r *Gotenberg) WebhookErrorURL(url, method string) *Gotenberg {
	r.Req.Header("Gotenberg-Webhook-Error-Url", url).
		Header("Gotenberg-Webhook-Error-Method", method)
	return r
}

// WebhookHeader adds a custom header to be sent with webhook requests.
// Multiple headers can be added by calling this method multiple times.
func (r *Gotenberg) WebhookHeader(key, value string) *Gotenberg {
	if r.Wh == nil {
		r.Wh = make(map[string]string)
	}
	r.Wh[key] = value
	return r
}

// DownloadFrom sets the download from configuration.
func (r *Gotenberg) DownloadFrom(url string, headers map[string]string) *Gotenberg {
	r.Df = append(r.Df, downloadFrom{URL: url, Headers: headers})
	return r
}

// OutputFilename sets the output filename.
func (r *Gotenberg) OutputFilename(filename string) *Gotenberg {
	r.Req.Header("Gotenberg-Output-Filename", filename)
	return r
}

// Trace sets the request trace identifier for debugging and monitoring.
// If not set, Gotenberg will assign a unique UUID trace.
func (r *Gotenberg) Trace(trace string) *Gotenberg {
	r.Req.Header("Gotenberg-Trace", trace)
	return r
}

// Metadata sets the metadata for the operation.
func (r *Gotenberg) Metadata(key, value string) *Gotenberg {
	if r.Meta == nil {
		r.Meta = make(map[string]string)
	}
	r.Meta[key] = value
	return r
}
