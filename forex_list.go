package go_fmp

import (
	"encoding/json"
)

// ForexListResponse represents the response from the Forex Currency Pairs API
type ForexListResponse struct {
	Symbol       string `json:"symbol"`
	FromCurrency string `json:"fromCurrency"`
	ToCurrency   string `json:"toCurrency"`
	FromName     string `json:"fromName"`
	ToName       string `json:"toName"`
}

// ForexList retrieves a comprehensive list of all currency pairs traded on the forex market
func (c *Client) ForexList() ([]ForexListResponse, error) {
	resp, err := c.get("https://financialmodelingprep.com/stable/forex-list", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []ForexListResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
