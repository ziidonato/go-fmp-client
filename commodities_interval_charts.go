package go_fmp

import (
	"fmt"
	"time"
)

// CommoditiesIntervalChartsParams represents the parameters for the Commodities Interval Charts API
type CommoditiesIntervalChartsParams struct {
	Symbol   string  `json:"symbol"`   // Required: Commodity symbol (e.g., "CLUSD")
	Interval string  `json:"interval"` // Required: Time interval (1min, 5min, 15min, 30min, 1hour, 4hour)
	From     *string `json:"from"`     // Optional: Start date (YYYY-MM-DD format)
	To       *string `json:"to"`       // Optional: End date (YYYY-MM-DD format)
}

// CommoditiesIntervalChartsResponse represents the response from the Commodities Interval Charts API
type CommoditiesIntervalChartsResponse struct {
	Date   time.Time  `json:"date"`
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
}

// commoditiesIntervalChart is a helper function to avoid code duplication for interval-based charts
func (c *Client) commoditiesIntervalChart(interval string, params CommoditiesIntervalChartsParams) ([]CommoditiesIntervalChartsResponse, error) {
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

	var result []CommoditiesIntervalChartsResponse
	err := c.doRequest(fmt.Sprintf(c.BaseURL+"/historical-chart/%s", interval), urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CommoditiesChart1Min retrieves 1-minute interval data for commodities
func (c *Client) CommoditiesChart1Min(params CommoditiesIntervalChartsParams) ([]CommoditiesIntervalChartsResponse, error) {
	return c.commoditiesIntervalChart("1min", params)
}

// CommoditiesChart5Min retrieves 5-minute interval data for commodities
func (c *Client) CommoditiesChart5Min(params CommoditiesIntervalChartsParams) ([]CommoditiesIntervalChartsResponse, error) {
	return c.commoditiesIntervalChart("5min", params)
}

// CommoditiesChart1Hour retrieves 1-hour interval data for commodities
func (c *Client) CommoditiesChart1Hour(params CommoditiesIntervalChartsParams) ([]CommoditiesIntervalChartsResponse, error) {
	return c.commoditiesIntervalChart("1hour", params)
}
