package go_fmp

import (
	"encoding/json"
	"fmt"
)

// Index5MinChartResponse represents the response from the 5-minute interval index price API
type Index5MinChartResponse struct {
	Date   string  `json:"date"`
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
}

// GetIndex5MinChart retrieves 5-minute interval intraday price data for stock indexes
func (c *Client) GetIndex5MinChart(symbol, from, to string) ([]Index5MinChartResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	params := map[string]string{
		"symbol": symbol,
	}

	if from != " {
		params["from"] = from
	}
	if to != " {
		params["to"] = to
	}

	url := "https://financialmodelingprep.com/stable/historical-chart/5min"

	return doRequest[[]Index5MinChartResponse](c, url, params)
}
