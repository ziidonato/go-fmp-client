package go_fmp

import (
	"fmt"
	"time"
)

// ForexChartParams represents the parameters for the Forex Chart API
type ForexChartParams struct {
	Symbol   string `json:"symbol"`   // Required: Forex pair symbol (e.g., "EURUSD")
	Interval string `json:"interval"` // Required: Time interval (e.g., "1min", "5min", "1hour")
	From     string `json:"from"`     // Optional: Start date (YYYY-MM-DD format)
	To       string `json:"to"`       // Optional: End date (YYYY-MM-DD format)
}

// ForexChartResponse represents the response from the Forex Chart API
type ForexChartResponse struct {
	Date   time.Time  `json:"date"`
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
}

// ForexChart1Min retrieves real-time, 1-minute intraday forex data for currency pairs
func (c *Client) ForexChart1Min(params ForexChartParams) ([]ForexChartResponse, error) {
	return c.getForexChartInterval(params, "1min")
}

// ForexChart5Min retrieves real-time, 5-minute intraday forex data for currency pairs
func (c *Client) ForexChart5Min(params ForexChartParams) ([]ForexChartResponse, error) {
	return c.getForexChartInterval(params, "5min")
}

// ForexChart1Hour retrieves real-time, 1-hour intraday forex data for currency pairs
func (c *Client) ForexChart1Hour(params ForexChartParams) ([]ForexChartResponse, error) {
	return c.getForexChartInterval(params, "1hour")
}

func (c *Client) getForexChartInterval(params ForexChartParams, interval string) ([]ForexChartResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}
	urlParams := map[string]string{
		"symbol": params.Symbol,
	}
	if params.From != "" {
		urlParams["from"] = params.From
	}
	if params.To != "" {
		urlParams["to"] = params.To
	}
	var result []ForexChartResponse
	url := c.BaseURL + "/historical-chart/" + interval
	err := c.doRequest(url, urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	return result, nil
}
