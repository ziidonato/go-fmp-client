package go_fmp

import (
	"encoding/json"
	"fmt"
)

// CommoditiesIntervalChartParams represents the parameters for the interval-based commodities chart APIs
type CommoditiesIntervalChartParams struct {
	Symbol string  `json:"symbol"` // Required: Commodity symbol (e.g., "GCUSD")
	From   *string `json:"from"`   // Optional: Start date (e.g., "2024-01-01")
	To     *string `json:"to"`     // Optional: End date (e.g., "2024-03-01")
}

// CommoditiesIntervalChartResponse represents the response from the interval-based commodities chart APIs
type CommoditiesIntervalChartResponse struct {
	Date   string  `json:"date"`
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
}

// commoditiesIntervalChart is a helper function to avoid code duplication for interval-based charts
func (c *Client) commoditiesIntervalChart(interval string, params CommoditiesIntervalChartParams) ([]CommoditiesIntervalChartResponse, error) {
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

	return doRequest[[]CommoditiesIntervalChartResponse](c, fmt.Sprintf("https://financialmodelingprep.com/stable/historical-chart/%s", interval)
}

// CommoditiesChart1Min retrieves 1-minute interval data for commodities
func (c *Client) CommoditiesChart1Min(params CommoditiesIntervalChartParams) ([]CommoditiesIntervalChartResponse, error) {
	return c.commoditiesIntervalChart("1min", params)
}

// CommoditiesChart5Min retrieves 5-minute interval data for commodities
func (c *Client) CommoditiesChart5Min(params CommoditiesIntervalChartParams) ([]CommoditiesIntervalChartResponse, error) {
	return c.commoditiesIntervalChart("5min", params)
}

// CommoditiesChart1Hour retrieves 1-hour interval data for commodities
func (c *Client) CommoditiesChart1Hour(params CommoditiesIntervalChartParams) ([]CommoditiesIntervalChartResponse, error) {
	return c.commoditiesIntervalChart("1hour", params)
}
