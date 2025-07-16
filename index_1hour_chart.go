package go_fmp

import (
	"encoding/json"
	"fmt"
)

// Index1HourChartResponse represents the response from the 1-hour interval index price API
type Index1HourChartResponse struct {
	Date   string  `json:"date"`
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`

// GetIndex1HourChart retrieves 1-hour interval intraday data for stock indexes
func (c *Client) GetIndex1HourChart(symbol, from, to string) ([]Index1HourChartResponse, error) {
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

	url := "https://financialmodelingprep.com/stable/historical-chart/1hour"

	return doRequest[[]Index1HourChartResponse](c, url, params)
