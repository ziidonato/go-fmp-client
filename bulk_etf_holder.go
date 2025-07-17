package go_fmp

import (
	"fmt"
	"time"
)

// BulkETFHolderParams represents the parameters for the Bulk ETF Holder API
type BulkETFHolderParams struct {
	Date string `json:"date"` // Required: date (e.g., "2024-10-22")
}

// BulkETFHolderResponse represents the response from the Bulk ETF Holder API
type BulkETFHolderResponse struct {
	Symbol           string    `json:"symbol"`
	Asset            string    `json:"asset"`
	Name             string    `json:"name"`
	SharesHeld       int64     `json:"sharesHeld"`
	WeightPercentage float64   `json:"weightPercentage"`
	MarketValue      float64   `json:"marketValue"`
	Updated          string    `json:"updated"`
	LastUpdated      time.Time `json:"lastUpdated"`
}

// BulkETFHolder retrieves ETF holdings data for multiple ETFs
func (c *Client) BulkETFHolder(params BulkETFHolderParams) ([]BulkETFHolderResponse, error) {
	urlParams := map[string]string{
		"date": params.Date,
	}

	var result []BulkETFHolderResponse
	err := c.doRequest(c.BaseURL+"/bulk-etf-holder", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to get bulk ETF holder: %w", err)
	}

	return result, nil
}
