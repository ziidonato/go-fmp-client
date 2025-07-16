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
	return doRequest[[]ForexListResponse](c, "https://financialmodelingprep.com/stable/forex-list", nil)
}
