package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// PeersBulkResponse represents the response from the Stock Peers Bulk API
type PeersBulkResponse struct {
	Symbol string `json:"symbol"`
	Peers  string `json:"peers"`
}

// GetPeersBulk retrieves peer companies for all stocks in the database
func (c *Client) GetPeersBulk() ([]PeersBulkResponse, error) {
	// Build the URL
	baseURL := c.BaseURL + "/peers-bulk"
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %w", err)
	}

	// Add API key if available
	if c.APIKey != "" {
		q := u.Query()
		q.Set("apikey", c.APIKey)
		u.RawQuery = q.Encode()
	}

	// Create the request
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set headers
	req.Header.Set("User-Agent", "fmp-go-client")
	req.Header.Set("Accept", "application/json")

	// Make the request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	// Parse the response
	var result []PeersBulkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
