package go_fmp

import (
	"encoding/json"
	"fmt"
)

// CommoditiesFullChartParams represents the parameters for the Full Chart API
type CommoditiesFullChartParams struct {
	Symbol string  `json:"symbol"` // Required: Commodity symbol (e.g., "GCUSD")
	From   *string `json:"from"`   // Optional: Start date (e.g., "2025-01-10")
	To     *string `json:"to"`     // Optional: End date (e.g., "2025-04-10")
}

// CommoditiesFullChartResponse represents the response from the Full Chart API
type CommoditiesFullChartResponse struct {
	Symbol        string  `json:"symbol"`
	Date          string  `json:"date"`
	Open          float64 `json:"open"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	Close         float64 `json:"close"`
	Volume        int64   `json:"volume"`
	Change        float64 `json:"change"`
	ChangePercent float64 `json:"changePercent"`
	VWAP          float64 `json:"vwap"`
}

// CommoditiesFullChart retrieves full historical end-of-day price data for commodities
func (c *Client) CommoditiesFullChart(params CommoditiesFullChartParams) ([]CommoditiesFullChartResponse, error) {
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

	return doRequest[[]CommoditiesFullChartResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
}
