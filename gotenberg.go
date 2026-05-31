// Package gotenberg provides a client for the Gotenberg service.
// It offers a convenient API for document conversion using various engines.
package gotenberg

import (
	"log/slog"
	"net/http"

	"github.com/nativebpm/gotenberg/internal/chromium"
	"github.com/nativebpm/gotenberg/internal/gotenberg"
	"github.com/nativebpm/gotenberg/internal/libreoffice"
	"github.com/nativebpm/gotenberg/internal/pdfengines"
	"github.com/nativebpm/httpclient"
)

// Re-export common types for easier access.
type Response = gotenberg.Response

// Client is a Gotenberg HTTP client that wraps the base HTTP client
// with Gotenberg-specific functionality for document conversion.
type Client struct {
	client *httpclient.HTTPClient
}

// NewClient creates a new Gotenberg client with the given HTTP client and base URL.
// Returns an error if the base URL is invalid.
func NewClient(httpClient http.Client, baseURL string) (*Client, error) {
	client, err := httpclient.NewClient(httpClient, baseURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) WithLogger(logger *slog.Logger) *Client {
	c.client = c.client.WithLogger(logger)
	return c
}

func (c *Client) Use(middleware httpclient.Middleware) *Client {
	c.client = c.client.Use(middleware)
	return c
}

func (c *Client) Chromium() *chromium.Chromium {
	return chromium.NewChromium(c.client)
}

func (c *Client) LibreOffice() *libreoffice.LibreOffice {
	return libreoffice.NewLibreOffice(c.client)
}

func (c *Client) PDFEngines() *pdfengines.PDFEngines {
	return pdfengines.NewPDFEngines(c.client)
}
