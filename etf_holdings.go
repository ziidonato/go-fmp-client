package go_fmp

import (
	"encoding/json"
	"fmt"
)

// ETFHoldingsParams represents the parameters for the ETF & Fund Holdings API
type ETFHoldingsParams struct {
	Symbol string `json:"symbol"` // Required: ETF/Fund symbol (e.g., "SPY")
}

// ETFHoldingsResponse represents the response from the ETF & Fund Holdings API
type ETFHoldingsResponse struct {
	Symbol           string  `json:"symbol"`
	Asset            string  `json:"asset"`
	Name             string  `json:"name"`
	ISIN             string  `json:"isin"`
	SecurityCusip    string  `json:"securityCusip"`
	SharesNumber     int64   `json:"sharesNumber"`
	WeightPercentage float64 `json:"weightPercentage"`
	MarketValue      float64 `json:"marketValue"`
	UpdatedAt        string  `json:"updatedAt"`
	Updated          string  `json:"updated"`
}

// ETFHoldings retrieves a detailed breakdown of the assets held within ETFs and mutual funds
func (c *Client) ETFHoldings(params ETFHoldingsParams) ([]ETFHoldingsResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	resp, err := c.get("https://financialmodelingprep.com/stable/etf/holdings", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []ETFHoldingsResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
