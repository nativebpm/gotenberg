package gotenberg

import (
	"context"
	"encoding/json"
	"io"

	"github.com/nativebpm/httpstream"
)

// HealthResponse represents the response from the health check endpoint.
type HealthResponse struct {
	Status  string         `json:"status"`
	Details map[string]any `json:"details"`
}

// getBytes performs a HTTP GET request to the specified path and returns the response body bytes.
func getBytes(ctx context.Context, stream *httpstream.Client, path string) ([]byte, error) {
	resp, err := stream.Request(ctx, httpstream.GET, path).Send()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

// GetHealth performs a health check on the Gotenberg service.
// It returns the overall status and details about each module (e.g., Chromium, LibreOffice).
func (c *Client) GetHealth(ctx context.Context) (*HealthResponse, error) {
	resp, err := c.HttpStream.Request(ctx, httpstream.GET, "/health").Send()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var healthResp HealthResponse
	if err := json.NewDecoder(resp.Body).Decode(&healthResp); err != nil {
		return nil, err
	}

	return &healthResp, nil
}

// GetVersion returns the Gotenberg service version.
func (c *Client) GetVersion(ctx context.Context) (string, error) {
	b, err := getBytes(ctx, c.HttpStream, "/version")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// GetMetrics returns the Prometheus metrics from the Gotenberg service.
func (c *Client) GetMetrics(ctx context.Context) (string, error) {
	b, err := getBytes(ctx, c.HttpStream, "/prometheus/metrics")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
