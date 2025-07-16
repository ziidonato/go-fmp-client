package go_fmp

import (
	"encoding/json"
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
	Date   string  `json:"date"`
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

	resp, err := c.get("https://financialmodelingprep.com/stable/historical-price-eod/light", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []StockChartLightResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
