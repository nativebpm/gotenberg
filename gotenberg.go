// Package gotenberg provides a client for the Gotenberg service.
// It offers a convenient API for document conversion using various engines.
package gotenberg

import (
	"net/http"

	"github.com/nativebpm/gotenberg/internal/chromium"
	"github.com/nativebpm/gotenberg/internal/gotenberg"
	"github.com/nativebpm/gotenberg/internal/libreoffice"
	"github.com/nativebpm/gotenberg/internal/pdfengines"
	"github.com/nativebpm/connectors/httpstream"
)

// Re-export common types for easier access.
type Response = gotenberg.Response

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

func (c *Client) Use(middleware func(http.RoundTripper) http.RoundTripper) *Client {
	c.HttpStream = c.HttpStream.Use(middleware)
	return c
}

func (c *Client) Chromium() *chromium.Chromium {
	return chromium.NewChromium(c.HttpStream)
}

func (c *Client) LibreOffice() *libreoffice.LibreOffice {
	return libreoffice.NewLibreOffice(c.HttpStream)
}

func (c *Client) PDFEngines() *pdfengines.PDFEngines {
	return pdfengines.NewPDFEngines(c.HttpStream)
}
