package go_fmp

import (
	"encoding/json"
	"fmt"
)

// ETFCountryWeightingsParams represents the parameters for the ETF & Fund Country Allocation API
type ETFCountryWeightingsParams struct {
	Symbol string `json:"symbol"` // Required: ETF/Fund symbol (e.g., "SPY")
}

// ETFCountryWeightingsResponse represents the response from the ETF & Fund Country Allocation API
type ETFCountryWeightingsResponse struct {
	Country          string `json:"country"`
	WeightPercentage string `json:"weightPercentage"`
}

// ETFCountryWeightings retrieves country allocation data for ETFs and mutual funds
func (c *Client) ETFCountryWeightings(params ETFCountryWeightingsParams) ([]ETFCountryWeightingsResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	return doRequest[[]ETFCountryWeightingsResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
}
