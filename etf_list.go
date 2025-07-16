package go_fmp

import (
	"encoding/json"
)

// ETFListResponse represents the response from the ETF Symbol Search API
type ETFListResponse struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

// ETFList retrieves ticker symbols and company names for Exchange Traded Funds (ETFs)
func (c *Client) ETFList() ([]ETFListResponse, error) {
	return doRequest[[]ETFListResponse](c, "https://financialmodelingprep.com/stable/etf-list", nil)
}
