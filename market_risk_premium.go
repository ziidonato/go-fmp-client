package go_fmp

// MarketRiskPremiumResponse represents the response from the Market Risk Premium API
type MarketRiskPremiumResponse struct {
	Country Country `json:"country"`
	CountryRiskPremium float64 `json:"countryRiskPremium"`
}

// MarketRiskPremium retrieves the global market risk premium data
func (c *Client) MarketRiskPremium() ([]MarketRiskPremiumResponse, error) {
	url := c.BaseURL + "/market-risk-premium"
	var result []MarketRiskPremiumResponse
	if err := c.doRequest(url, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
