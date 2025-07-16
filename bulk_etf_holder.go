package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// ETFHolderBulkParams represents the parameters for the ETF Holder Bulk API
type ETFHolderBulkParams struct {
	Part string `json:"part"` // Required: part number (e.g., "1")
}

// ETFHolderBulkResponse represents the response from the ETF Holder Bulk API
type ETFHolderBulkResponse struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	SharesNumber     string `json:"sharesNumber"`
	Asset            string `json:"asset"`
	WeightPercentage string `json:"weightPercentage"`
	CUSIP            string `json:"cusip"`
	ISIN             string `json:"isin"`
	MarketValue      string `json:"marketValue"`
	LastUpdated      string `json:"lastUpdated"`
}

// ETFHolderBulk retrieves detailed information about assets and shares held by ETFs
func (c *Client) ETFHolderBulk(params ETFHolderBulkParams) ([]ETFHolderBulkResponse, error) {
	// Validate required parameters
	if params.Part == "" {
		return nil, fmt.Errorf("part parameter is required")
	}

	// Build the URL
	baseURL := c.BaseURL + "/etf-holder-bulk"
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %w", err)
	}

	// Add query parameters
	q := u.Query()
	q.Set("part", params.Part)
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
	var result []ETFHolderBulkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
