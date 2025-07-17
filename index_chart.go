package go_fmp

import (
	"fmt"
	"time"
)

// IndexChartParams represents the parameters for the Index Chart API
type IndexChartParams struct {
	Symbol string  `json:"symbol"` // Required: Index symbol (e.g., "^GSPC")
	From   *string `json:"from"`   // Optional: Start date (YYYY-MM-DD format)
	To     *string `json:"to"`     // Optional: End date (YYYY-MM-DD format)
}

// IndexChartResponse represents the response from the Index Chart API
type IndexChartResponse struct {
	Date   time.Time  `json:"date"`
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
}

type Index5MinChartResponse = IndexChartResponse

type Index1HourChartResponse = IndexChartResponse

func (c *Client) GetIndex1MinChart(symbol, from, to string) ([]IndexChartResponse, error) {
	return c.getIndexChartInterval(symbol, from, to, "1min")
}

func (c *Client) GetIndex5MinChart(symbol, from, to string) ([]Index5MinChartResponse, error) {
	return c.getIndexChartInterval(symbol, from, to, "5min")
}

func (c *Client) GetIndex1HourChart(symbol, from, to string) ([]Index1HourChartResponse, error) {
	return c.getIndexChartInterval(symbol, from, to, "1hour")
}

func (c *Client) getIndexChartInterval(symbol, from, to, interval string) ([]IndexChartResponse, error) {
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
	url := c.BaseURL + "/historical-chart/" + interval
	var result []IndexChartResponse
	err := c.doRequest(url, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
