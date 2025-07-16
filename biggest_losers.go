package go_fmp

// BiggestLosersResponse represents the response from the biggest losers API
type BiggestLosersResponse struct {
	Symbol           string  `json:"symbol"`
	Name             string  `json:"name"`
	Change           float64 `json:"change"`
	Price            float64 `json:"price"`
	ChangePercentage float64 `json:"changePercentage"`
}

// BiggestLosers retrieves the stocks with the biggest price declines
func (c *Client) BiggestLosers() ([]BiggestLosersResponse, error) {
	url := c.BaseURL + "/stock-market/losers"
	var result []BiggestLosersResponse
	if err := c.doRequest(url, map[string]string{}, &result); err != nil {
		return nil, err
	}
	return result, nil
}
