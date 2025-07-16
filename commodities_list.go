package go_fmp

import (
	"encoding/json"
)

// CommoditiesListResponse represents the response from the Commodities List API
type CommoditiesListResponse struct {
	Symbol     string  `json:"symbol"`
	Name       string  `json:"name"`
	Exchange   *string `json:"exchange"`
	TradeMonth string  `json:"tradeMonth"`
	Currency   string  `json:"currency"`
}

// CommoditiesList retrieves an extensive list of tracked commodities across various sectors
func (c *Client) CommoditiesList() ([]CommoditiesListResponse, error) {
	resp, err := c.doRequest("https://financialmodelingprep.com/stable/commodities-list", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []CommoditiesListResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
