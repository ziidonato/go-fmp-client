package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// EarningsSurprisesBulkParams represents the parameters for the Earnings Surprises Bulk API
type EarningsSurprisesBulkParams struct {
	Year string `json:"year"` // Required: year (e.g., "2024")
}

// EarningsSurprisesBulkResponse represents the response from the Earnings Surprises Bulk API
type EarningsSurprisesBulkResponse struct {
	Symbol       string `json:"symbol"`
	Date         string `json:"date"`
	EPSActual    string `json:"epsActual"`
	EPSEstimated string `json:"epsEstimated"`
	LastUpdated  string `json:"lastUpdated"`
}

// GetEarningsSurprisesBulk retrieves bulk data on annual earnings surprises
func (c *Client) GetEarningsSurprisesBulk(params EarningsSurprisesBulkParams) ([]EarningsSurprisesBulkResponse, error) {
	// Validate required parameters
	if params.Year == "" {
		return nil, fmt.Errorf("year parameter is required")
	}

	// Build the URL
	baseURL := c.BaseURL + "/earnings-surprises-bulk"
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %w", err)
	}

	// Add query parameters
	q := u.Query()
	q.Set("year", params.Year)
	u.RawQuery = q.Encode()

	// Add API key if available
	if c.APIKey != "" {
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
	var result []EarningsSurprisesBulkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
