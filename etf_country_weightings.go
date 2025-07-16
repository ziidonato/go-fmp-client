package go_fmp

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
func (c *Client) ETFCountryWeightings(symbol string) ([]ETFCountryWeightingsResponse, error) {
	url := c.BaseURL + "/etf/country-weightings"
	params := map[string]string{"symbol": symbol}
	var result []ETFCountryWeightingsResponse
	if err := c.doRequest(url, params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
