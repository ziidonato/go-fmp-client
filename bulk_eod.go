package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// EODBulkParams represents the parameters for the EOD Bulk API
type EODBulkParams struct {
	Date string `json:"date"` // Required: date (e.g., "2024-10-22")
}

// EODBulkResponse represents the response from the EOD Bulk API
type EODBulkResponse struct {
	Symbol   string `json:"symbol"`
	Date     string `json:"date"`
	Open     string `json:"open"`
	Low      string `json:"low"`
	High     string `json:"high"`
	Close    string `json:"close"`
	AdjClose string `json:"adjClose"`
	Volume   string `json:"volume"`
}

// GetEODBulk retrieves end-of-day stock price data for multiple symbols in bulk
func (c *Client) GetEODBulk(params EODBulkParams) ([]EODBulkResponse, error) {
	// Validate required parameters
	if params.Date == "" {
		return nil, fmt.Errorf("date parameter is required")
	}

	// Build the URL
	baseURL := "https://financialmodelingprep.com/stable/eod-bulk"
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %w", err)
	}

	// Add query parameters
	q := u.Query()
	q.Set("date", params.Date)
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
	var result []EODBulkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
