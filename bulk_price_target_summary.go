package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// PriceTargetSummaryBulkResponse represents the response from the Price Target Summary Bulk API
type PriceTargetSummaryBulkResponse struct {
	Symbol                    string `json:"symbol"`
	LastMonthCount            string `json:"lastMonthCount"`
	LastMonthAvgPriceTarget   string `json:"lastMonthAvgPriceTarget"`
	LastQuarterCount          string `json:"lastQuarterCount"`
	LastQuarterAvgPriceTarget string `json:"lastQuarterAvgPriceTarget"`
	LastYearCount             string `json:"lastYearCount"`
	LastYearAvgPriceTarget    string `json:"lastYearAvgPriceTarget"`
	AllTimeCount              string `json:"allTimeCount"`
	AllTimeAvgPriceTarget     string `json:"allTimeAvgPriceTarget"`
	Publishers                string `json:"publishers"`
}

// GetPriceTargetSummaryBulk retrieves comprehensive overview of price targets for all listed symbols
func (c *Client) GetPriceTargetSummaryBulk() ([]PriceTargetSummaryBulkResponse, error) {
	// Build the URL
	baseURL := "https://financialmodelingprep.com/stable/price-target-summary-bulk"
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
	var result []PriceTargetSummaryBulkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
