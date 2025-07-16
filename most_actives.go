package go_fmp

// MostActivesResponse represents the response from the most actives API
type MostActivesResponse struct {
	Symbol           string  `json:"symbol"`
	Name             string  `json:"name"`
	Change           float64 `json:"change"`
	Price            float64 `json:"price"`
	ChangePercentage float64 `json:"changePercentage"`
}

// MostActives retrieves the stocks with the highest trading volume
func (c *Client) MostActives() ([]MostActivesResponse, error) {
	url := c.BaseURL + "/stock-market/actives"
	var result []MostActivesResponse
	if err := c.doRequest(url, map[string]string{}, &result); err != nil {
		return nil, err
	}
	return result, nil
}
