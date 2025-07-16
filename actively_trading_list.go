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
	return doRequest[[]ActivelyTradingListResponse](c, "https://financialmodelingprep.com/stable/actively-trading-list", nil)
}
