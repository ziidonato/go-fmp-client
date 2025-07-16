package go_fmp

// BiggestGainersResponse represents the response from the biggest gainers API
type BiggestGainersResponse struct {
	Symbol           string  `json:"symbol"`
	Name             string  `json:"name"`
	Change           float64 `json:"change"`
	Price            float64 `json:"price"`
	ChangePercentage float64 `json:"changePercentage"`
}

// BiggestGainers retrieves the stocks with the biggest price increases
func (c *Client) BiggestGainers() ([]BiggestGainersResponse, error) {
	url := "https://financialmodelingprep.com/stable/stock-market/gainers"
	var result []BiggestGainersResponse
	if err := c.doRequest(url, map[string]string{}, &result); err != nil {
		return nil, err
	}
	return result, nil
}
