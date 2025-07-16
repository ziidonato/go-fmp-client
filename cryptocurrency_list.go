package go_fmp

import (
	"encoding/json"
)

// CryptocurrencyListResponse represents the response from the Cryptocurrency List API
type CryptocurrencyListResponse struct {
	Symbol       string  `json:"symbol"`
	Name         string  `json:"name"`
	Currency     string  `json:"currency"`
	Exchange     *string `json:"exchange"`
	ExchangeName *string `json:"exchangeName"`
}

// CryptocurrencyList retrieves a comprehensive list of all cryptocurrencies traded on exchanges worldwide
func (c *Client) CryptocurrencyList() ([]CryptocurrencyListResponse, error) {
	return doRequest[[]CryptocurrencyListResponse](c, "https://financialmodelingprep.com/stable/cryptocurrency-list", nil)
}
