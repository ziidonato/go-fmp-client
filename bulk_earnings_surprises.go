package go_fmp

import (
	"fmt"
	"time"
)

// BulkEarningsSurprisesParams represents the parameters for the Bulk Earnings Surprises API
type BulkEarningsSurprisesParams struct {
	Date string `json:"date"` // Required: date (e.g., "2024-10-22")
}

// BulkEarningsSurprisesResponse represents the response from the Bulk Earnings Surprises API
type BulkEarningsSurprisesResponse struct {
	Symbol              string    `json:"symbol"`
	Date                time.Time `json:"date"`
	EstimatedEPS        float64   `json:"estimatedEPS"`
	ActualEPS           float64   `json:"actualEPS"`
	LastUpdated         time.Time `json:"lastUpdated"`
	EPSSurprise         float64   `json:"epsSurprise"`
	EstimatedRevenue    float64   `json:"estimatedRevenue"`
	ActualRevenue       float64   `json:"actualRevenue"`
	RevenueSurprise     float64   `json:"revenueSurprise"`
}

// BulkEarningsSurprises retrieves earnings surprise data for multiple stocks
func (c *Client) BulkEarningsSurprises(params BulkEarningsSurprisesParams) ([]BulkEarningsSurprisesResponse, error) {
	urlParams := map[string]string{
		"date": params.Date,
	}

	var result []BulkEarningsSurprisesResponse
	err := c.doRequest(c.BaseURL+"/bulk-earnings-surprises", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to get bulk earnings surprises: %w", err)
	}

	return result, nil
}
