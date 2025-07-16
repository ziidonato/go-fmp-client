package go_fmp

import (
	"encoding/json"
	"fmt"
)

// EODBulkParams represents the parameters for the EOD Bulk API
type EODBulkParams struct {
	Date string `json:"date"` // Required: date (e.g., "2024-10-22")

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

// GetEODBulk retrieves end-of-day stock price data for multiple symbols in bulk
func (c *Client) GetEODBulk(params EODBulkParams) ([]EODBulkResponse, error) {
	// Validate required parameters
	if params.Date == "" {
		return nil, fmt.Errorf("date parameter is required")
	}

	urlParams := map[string]string{
		"date": params.Date,
	}

	return doRequest[[]EODBulkResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
