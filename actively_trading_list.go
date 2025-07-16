package go_fmp

import (
	"encoding/json"
)

// ActivelyTradingListResponse represents the response from the Actively Trading List API
type ActivelyTradingListResponse struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

// ActivelyTradingList retrieves all actively trading companies and financial instruments
func (c *Client) ActivelyTradingList() ([]ActivelyTradingListResponse, error) {
	resp, err := c.get("https://financialmodelingprep.com/stable/actively-trading-list", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []ActivelyTradingListResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
