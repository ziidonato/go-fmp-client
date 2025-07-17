package go_fmp

import (
	"fmt"
	"time"
)

// CommoditiesLightChartParams represents the parameters for the Commodities Light Chart API
type CommoditiesLightChartParams struct {
	Symbol string  `json:"symbol"` // Required: Commodity symbol (e.g., "CLUSD")
	From   *string `json:"from"`   // Optional: Start date (YYYY-MM-DD format)
	To     *string `json:"to"`     // Optional: End date (YYYY-MM-DD format)
}

// CommoditiesLightChartResponse represents the response from the Commodities Light Chart API
type CommoditiesLightChartResponse struct {
	Date   time.Time  `json:"date"`
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
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
