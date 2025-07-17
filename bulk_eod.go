package go_fmp

import (
	"fmt"
	"time"
)

// BulkEODParams represents the parameters for the Bulk EOD API
type BulkEODParams struct {
	Date string `json:"date"` // Required: date (e.g., "2024-10-22")
}

// BulkEODResponse represents the response from the Bulk EOD API
type BulkEODResponse struct {
	Symbol               string    `json:"symbol"`
	Date                 time.Time `json:"date"`
	Open                 float64   `json:"open"`
	High                 float64   `json:"high"`
	Low                  float64   `json:"low"`
	Close                float64   `json:"close"`
	AdjClose             float64   `json:"adjClose"`
	Volume               int64     `json:"volume"`
	UnadjustedVolume     int64     `json:"unadjustedVolume"`
	Change               float64   `json:"change"`
	ChangePercent        float64   `json:"changePercent"`
	VWAP                 float64   `json:"vwap"`
	Label                string    `json:"label"`
	ChangeOverTime       float64   `json:"changeOverTime"`
}

// BulkEOD retrieves end-of-day pricing data for multiple stocks
func (c *Client) BulkEOD(params BulkEODParams) ([]BulkEODResponse, error) {
	urlParams := map[string]string{
		"date": params.Date,
	}

	var result []BulkEODResponse
	err := c.doRequest(c.BaseURL+"/bulk-eod", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to get bulk EOD: %w", err)
	}

	return result, nil
}
