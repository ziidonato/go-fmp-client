package go_fmp

import (
	"fmt"
)

// BulkPriceTargetSummaryParams represents the parameters for the Bulk Price Target Summary API
type BulkPriceTargetSummaryParams struct {
	Date string `json:"date"` // Required: date (e.g., "2024-10-22")
}

// BulkPriceTargetSummaryResponse represents the response from the Bulk Price Target Summary API
type BulkPriceTargetSummaryResponse struct {
	Symbol               string  `json:"symbol"`
	Current              float64 `json:"current"`
	Target               float64 `json:"target"`
	Low                  float64 `json:"low"`
	High                 float64 `json:"high"`
	YearlyLow            float64 `json:"yearlyLow"`
	YearlyHigh           float64 `json:"yearlyHigh"`
	NumberOfAnalysts     int     `json:"numberOfAnalysts"`
}

// BulkPriceTargetSummary retrieves analyst price targets for multiple stocks
func (c *Client) BulkPriceTargetSummary(params BulkPriceTargetSummaryParams) ([]BulkPriceTargetSummaryResponse, error) {
	urlParams := map[string]string{
		"date": params.Date,
	}

	var result []BulkPriceTargetSummaryResponse
	err := c.doRequest(c.BaseURL+"/bulk-price-target-summary", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to get bulk price target summary: %w", err)
	}

	return result, nil
}
