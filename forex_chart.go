package go_fmp

import (
	"fmt"
	"time"
)

// ForexChart1MinParams represents the parameters for the 1-Minute Forex Chart API
type ForexChart1MinParams struct {
	Symbol string  `json:"symbol"` // Required: Forex symbol (e.g., "EURUSD")
	From   *string `json:"from"`   // Optional: Start date (e.g., "2024-01-01")
	To     *string `json:"to"`     // Optional: End date (e.g., "2024-03-01")
}

type ForexChart5MinParams = ForexChart1MinParams

type ForexChart1HourParams = ForexChart1MinParams

// ForexChart1MinResponse represents the response from the 1-Minute Forex Chart API
type ForexChart1MinResponse struct {
	Date   time.Time `json:"date"`
	Open   float64   `json:"open"`
	Low    float64   `json:"low"`
	High   float64   `json:"high"`
	Close  float64   `json:"close"`
	Volume int64     `json:"volume"`
}

type ForexChart5MinResponse = ForexChart1MinResponse

type ForexChart1HourResponse = ForexChart1MinResponse

// ForexChart1Min retrieves real-time, 1-minute intraday forex data for currency pairs
func (c *Client) ForexChart1Min(params ForexChart1MinParams) ([]ForexChart1MinResponse, error) {
	return c.getForexChartInterval(params, "1min")
}

// ForexChart5Min retrieves real-time, 5-minute intraday forex data for currency pairs
func (c *Client) ForexChart5Min(params ForexChart5MinParams) ([]ForexChart5MinResponse, error) {
	return c.getForexChartInterval(params, "5min")
}

// ForexChart1Hour retrieves real-time, 1-hour intraday forex data for currency pairs
func (c *Client) ForexChart1Hour(params ForexChart1HourParams) ([]ForexChart1HourResponse, error) {
	return c.getForexChartInterval(params, "1hour")
}

func (c *Client) getForexChartInterval(params ForexChart1MinParams, interval string) ([]ForexChart1MinResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}
	urlParams := map[string]string{
		"symbol": params.Symbol,
	}
	if params.From != nil {
		urlParams["from"] = *params.From
	}
	if params.To != nil {
		urlParams["to"] = *params.To
	}
	var result []ForexChart1MinResponse
	url := c.BaseURL + "/historical-chart/" + interval
	err := c.doRequest(url, urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	return result, nil
}
