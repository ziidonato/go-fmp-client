package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// DCFBulkResponse represents the response from the DCF Bulk API
type DCFBulkResponse struct {
	Symbol     string `json:"symbol"`
	Date       string `json:"date"`
	DCF        string `json:"dcf"`
	StockPrice string `json:"Stock Price"`
}

// GetDCFBulk retrieves discounted cash flow valuations for multiple symbols
func (c *Client) GetDCFBulk() ([]DCFBulkResponse, error) {
	// Build the URL
	baseURL := "https://financialmodelingprep.com/stable/dcf-bulk"
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
	var result []DCFBulkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
