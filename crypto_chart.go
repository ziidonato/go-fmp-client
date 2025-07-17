package go_fmp

import (
	"fmt"
	"time"
)

// CryptoChartParams represents the parameters for the Crypto Chart API
type CryptoChartParams struct {
	Symbol   string `json:"symbol"`   // Required: Crypto symbol (e.g., "BTCUSD")
	Interval string `json:"interval"` // Required: Time interval (e.g., "1min", "5min", "1hour")
	From     string `json:"from"`     // Optional: Start date (YYYY-MM-DD format)
	To       string `json:"to"`       // Optional: End date (YYYY-MM-DD format)
}

// CryptoChartResponse represents the response from the Crypto Chart API
type CryptoChartResponse struct {
	Date   time.Time  `json:"date"`
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
	Volume float64 `json:"volume"`
}

// CryptoChart1Min retrieves 1-minute interval cryptocurrency price data
func (c *Client) CryptoChart1Min(params CryptoChartParams) ([]CryptoChartResponse, error) {
	params.Interval = "1min"
	return c.getCryptoChartInterval(params)
}

// CryptoChart5Min retrieves 5-minute interval cryptocurrency price data
func (c *Client) CryptoChart5Min(params CryptoChartParams) ([]CryptoChartResponse, error) {
	params.Interval = "5min"
	return c.getCryptoChartInterval(params)
}

// CryptoChart1Hour retrieves 1-hour interval cryptocurrency price data
func (c *Client) CryptoChart1Hour(params CryptoChartParams) ([]CryptoChartResponse, error) {
	params.Interval = "1hour"
	return c.getCryptoChartInterval(params)
}

func (c *Client) getCryptoChartInterval(params CryptoChartParams) ([]CryptoChartResponse, error) {
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

	url := fmt.Sprintf("%s/historical-chart/%s/%s", c.BaseURL, params.Interval, params.Symbol)
	var result []CryptoChartResponse
	err := c.doRequest(url, urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}