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

	resp, err := c.get("https://financialmodelingprep.com/stable/etf/country-weightings", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []ETFCountryWeightingsResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
