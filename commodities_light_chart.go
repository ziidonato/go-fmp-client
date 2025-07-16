package go_fmp

import (
	"fmt"
)

// CommoditiesLightChartParams represents the parameters for the Light Chart API
type CommoditiesLightChartParams struct {
	Symbol string  `json:"symbol"` // Required: Commodity symbol (e.g., "GCUSD")
	From   *string `json:"from"`   // Optional: Start date (e.g., "2025-01-10")
	To     *string `json:"to"`     // Optional: End date (e.g., "2025-04-10")
}

// CommoditiesLightChartResponse represents the response from the Light Chart API
type CommoditiesLightChartResponse struct {
	Symbol string  `json:"symbol"`
	Date   string  `json:"date"`
	Price  float64 `json:"price"`
	Volume int64   `json:"volume"`
}

// CommoditiesLightChart retrieves historical end-of-day prices for commodities
func (c *Client) CommoditiesLightChart(params CommoditiesLightChartParams) ([]CommoditiesLightChartResponse, error) {
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

	var result []CommoditiesLightChartResponse
	err := c.doRequest(c.BaseURL+"/historical-price-eod/light", urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
