package go_fmp

import (
	"encoding/json"
)

// MarketRiskPremiumResponse represents the response from the Market Risk Premium API
type MarketRiskPremiumResponse struct {
	Date              string  `json:"date"`
	MarketRiskPremium float64 `json:"marketRiskPremium"`
	RiskFreeRate      float64 `json:"riskFreeRate"`
	MarketReturn      float64 `json:"marketReturn"`
	Description       string  `json:"description"`
}

// MarketRiskPremium retrieves the market risk premium for specific dates
func (c *Client) MarketRiskPremium() ([]MarketRiskPremiumResponse, error) {
	resp, err := c.get("https://financialmodelingprep.com/stable/market-risk-premium", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []MarketRiskPremiumResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
