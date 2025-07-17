package go_fmp

import (
	"time"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// UpgradesDowngradesConsensusBulkResponse represents the response from the Upgrades Downgrades Consensus Bulk API
type UpgradesDowngradesConsensusBulkResponse struct {
	Symbol     string `json:"symbol"`
	StrongBuy  string `json:"strongBuy"`
	Buy        string `json:"buy"`
	Hold       string `json:"hold"`
	Sell       string `json:"sell"`
	StrongSell string `json:"strongSell"`
	Consensus GradeConsensus `json:"consensus"`
}

// UpgradesDowngradesConsensusBulk retrieves analyst ratings across all symbols
func (c *Client) UpgradesDowngradesConsensusBulk() ([]UpgradesDowngradesConsensusBulkResponse, error) {
	// Build the URL
	baseURL := c.BaseURL + "/upgrades-downgrades-consensus-bulk"
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
	var result []UpgradesDowngradesConsensusBulkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
