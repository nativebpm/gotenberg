package gotenberg

import (
	"context"
	"encoding/json"
	"io"

	"github.com/nativebpm/httpclient"
)

// HealthResponse represents the response from the health check endpoint.
type HealthResponse struct {
	Status  string         `json:"status"`
	Details map[string]any `json:"details"`
}

// GetHealth performs a health check on the Gotenberg service.
// It returns the overall status and details about each module (e.g., Chromium, LibreOffice).
func (c *Client) GetHealth(ctx context.Context) (*HealthResponse, error) {
	resp, err := c.client.Request(ctx, httpclient.GET, "/health").Send()
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

func (c *Client) GetVersion(ctx context.Context) (string, error) {
	resp, err := c.client.Request(ctx, httpclient.GET, "/version").Send()
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (c *Client) GetMetrics(ctx context.Context) (string, error) {
	resp, err := c.client.Request(ctx, httpclient.GET, "/prometheus/metrics").Send()
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
