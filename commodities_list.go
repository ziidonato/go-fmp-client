package go_fmp

// CommoditiesListResponse represents the response from the Commodities List API
type CommoditiesListResponse struct {
	Symbol     string  `json:"symbol"`
	Name       string  `json:"name"`
	Exchange *Exchange `json:"exchange"`
	TradeMonth string  `json:"tradeMonth"`
	Currency Currency `json:"currency"`
}

// CommoditiesList retrieves an extensive list of tracked commodities across various sectors
func (c *Client) CommoditiesList() ([]CommoditiesListResponse, error) {
	url := c.BaseURL + "/commodities-list"
	var result []CommoditiesListResponse
	if err := c.doRequest(url, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
