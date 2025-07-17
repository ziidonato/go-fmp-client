package go_fmp

import (
	"time"
	"fmt"
)

// StockChartLightParams represents the parameters for the Stock Chart Light API
type StockChartLightParams struct {
	Symbol string  `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	From   *string `json:"from"`   // Optional: Start date (e.g., "2025-01-10")
	To     *string `json:"to"`     // Optional: End date (e.g., "2025-04-10")
}

// StockChartLightResponse represents the response from the Stock Chart Light API
type StockChartLightResponse struct {
	Symbol string  `json:"symbol"`
	Date time.Time `json:"date"`
	Price  float64 `json:"price"`
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
