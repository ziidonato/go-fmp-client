package go_fmp

import (
	"fmt"
)

// PriceTargetSummaryParams represents the parameters for the Price Target Summary API
type PriceTargetSummaryParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// PriceTargetSummaryResponse represents the response from the Price Target Summary API
type PriceTargetSummaryResponse struct {
	Symbol                    string  `json:"symbol"`
	LastMonthCount            int     `json:"lastMonthCount"`
	LastMonthAvgPriceTarget   float64 `json:"lastMonthAvgPriceTarget"`
	LastQuarterCount          int     `json:"lastQuarterCount"`
	LastQuarterAvgPriceTarget float64 `json:"lastQuarterAvgPriceTarget"`
	LastYearCount             int     `json:"lastYearCount"`
	LastYearAvgPriceTarget    float64 `json:"lastYearAvgPriceTarget"`
	AllTimeCount              int     `json:"allTimeCount"`
	AllTimeAvgPriceTarget     float64 `json:"allTimeAvgPriceTarget"`
	Publishers                string  `json:"publishers"`
}

// PriceTargetSummary retrieves average price targets from analysts across various timeframes
func (c *Client) PriceTargetSummary(params PriceTargetSummaryParams) ([]PriceTargetSummaryResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	var result []PriceTargetSummaryResponse
	err := c.doRequest("https://financialmodelingprep.com/stable/price-target-summary", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, err
}
