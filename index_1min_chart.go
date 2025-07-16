package go_fmp

import (
	"encoding/json"
	"fmt"
)

// Index1MinChartResponse represents the response from the 1-minute interval index price API
type Index1MinChartResponse struct {
	Date   string  `json:"date"`
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`

// GetIndex1MinChart retrieves 1-minute interval intraday data for stock indexes
func (c *Client) GetIndex1MinChart(symbol, from, to string) ([]Index1MinChartResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	params := map[string]string{
		"symbol": symbol,
	}

	if from != "" {
		params["from"] = from
	}
	if to != "" {
		params["to"] = to
	}

	url := "https://financialmodelingprep.com/stable/historical-chart/1min"

	return doRequest[[]Index1MinChartResponse](c, url, params)
