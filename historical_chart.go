package go_fmp

import (
	"fmt"
	"strconv"
)

// HistoricalChartParams represents the parameters for the Historical Chart APIs
type HistoricalChartParams struct {
	Symbol      string  `json:"symbol"`      // Required: Stock symbol (e.g., "AAPL")
	From        *string `json:"from"`        // Optional: Start date (e.g., "2024-01-01")
	To          *string `json:"to"`          // Optional: End date (e.g., "2024-03-01")
	NonAdjusted *bool   `json:"nonAdjusted"` // Optional: Whether to use non-adjusted data (e.g., false)
}

// HistoricalChartResponse represents the response from the Historical Chart APIs
type HistoricalChartResponse struct {
	Date   string  `json:"date"`
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
}

// HistoricalChart1Min retrieves stock data in 1-minute intervals
func (c *Client) HistoricalChart1Min(params HistoricalChartParams) ([]HistoricalChartResponse, error) {
	return c.HistoricalChart("1min", params)
}

// HistoricalChart5Min retrieves stock data in 5-minute intervals
func (c *Client) HistoricalChart5Min(params HistoricalChartParams) ([]HistoricalChartResponse, error) {
	return c.HistoricalChart("5min", params)
}

// HistoricalChart15Min retrieves stock data in 15-minute intervals
func (c *Client) HistoricalChart15Min(params HistoricalChartParams) ([]HistoricalChartResponse, error) {
	return c.HistoricalChart("15min", params)
}

// HistoricalChart30Min retrieves stock data in 30-minute intervals
func (c *Client) HistoricalChart30Min(params HistoricalChartParams) ([]HistoricalChartResponse, error) {
	return c.HistoricalChart("30min", params)
}

// HistoricalChart1Hour retrieves stock data in 1-hour intervals
func (c *Client) HistoricalChart1Hour(params HistoricalChartParams) ([]HistoricalChartResponse, error) {
	return c.HistoricalChart("1hour", params)
}

// HistoricalChart4Hour retrieves stock data in 4-hour intervals
func (c *Client) HistoricalChart4Hour(params HistoricalChartParams) ([]HistoricalChartResponse, error) {
	return c.HistoricalChart("4hour", params)
}

// HistoricalChart is a helper function to avoid code duplication for different intervals
func (c *Client) HistoricalChart(interval string, params HistoricalChartParams) ([]HistoricalChartResponse, error) {
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

	if params.NonAdjusted != nil {
		urlParams["nonadjusted"] = strconv.FormatBool(*params.NonAdjusted)
	}

	endpoint := fmt.Sprintf("%s/historical-chart/%s", c.BaseURL, interval)
	var result []HistoricalChartResponse
	err := c.doRequest(endpoint, urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
