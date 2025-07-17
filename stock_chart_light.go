package go_fmp

import (
	"fmt"
	"time"
)

// StockChartLightParams represents the parameters for the Stock Chart Light API
type StockChartLightParams struct {
	Symbol string  `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	From   *string `json:"from"`   // Optional: Start date (YYYY-MM-DD format)
	To     *string `json:"to"`     // Optional: End date (YYYY-MM-DD format)
}

// StockChartLightResponse represents the response from the Stock Chart Light API
type StockChartLightResponse struct {
	Date   time.Time  `json:"date"`
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
}

// StockChartLight retrieves simplified stock chart data with essential charting information
func (c *Client) StockChartLight(params StockChartLightParams) ([]StockChartLightResponse, error) {
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

	url := c.BaseURL + "/historical-price-eod/light"
	var result []StockChartLightResponse
	if err := c.doRequest(url, urlParams, &result); err != nil {
		return nil, err
	}
	return result, nil
}
