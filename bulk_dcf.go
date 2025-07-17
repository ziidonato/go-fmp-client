package go_fmp

import (
	"fmt"
	"time"
)

// BulkDCFParams represents the parameters for the Bulk DCF API
type BulkDCFParams struct {
	Date string `json:"date"` // Required: date (e.g., "2024-10-22")
}

// BulkDCFResponse represents the response from the Bulk DCF API
type BulkDCFResponse struct {
	Date       time.Time `json:"date"`
	Symbol     string  `json:"symbol"`
	DCF        float64 `json:"dcf"`
	Price      float64 `json:"price"`
}

// BulkDCF retrieves discounted cash flow valuations for multiple stocks
func (c *Client) BulkDCF(params BulkDCFParams) ([]BulkDCFResponse, error) {
	urlParams := map[string]string{
		"date": params.Date,
	}

	var result []BulkDCFResponse
	err := c.doRequest(c.BaseURL+"/bulk-dcf", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to get bulk DCF: %w", err)
	}

	return result, nil
}
