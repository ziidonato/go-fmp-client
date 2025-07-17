package go_fmp

import (
	"fmt"
	"time"
)

// CommoditiesFullChartParams represents the parameters for the Commodities Full Chart API
type CommoditiesFullChartParams struct {
	Symbol string  `json:"symbol"` // Required: Commodity symbol (e.g., "CLUSD")
	From   *string `json:"from"`   // Optional: Start date (YYYY-MM-DD format)
	To     *string `json:"to"`     // Optional: End date (YYYY-MM-DD format)
}

// CommoditiesFullChartResponse represents the response from the Commodities Full Chart API
type CommoditiesFullChartResponse struct {
	Date          time.Time  `json:"date"`
	Open          float64 `json:"open"`
	Low           float64 `json:"low"`
	High          float64 `json:"high"`
	Close         float64 `json:"close"`
	AdjClose      float64 `json:"adjClose"`
	Volume        int64   `json:"volume"`
	UnadjustedVolume int64 `json:"unadjustedVolume"`
	Change        float64 `json:"change"`
	ChangePercent float64 `json:"changePercent"`
	Vwap          float64 `json:"vwap"`
	Label         string  `json:"label"`
	ChangeOverTime float64 `json:"changeOverTime"`
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

	var result []CommoditiesFullChartResponse
	err := c.doRequest(c.BaseURL+"/historical-price-eod/full", urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
