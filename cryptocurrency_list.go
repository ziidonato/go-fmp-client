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
	resp, err := c.get("https://financialmodelingprep.com/stable/cryptocurrency-list", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []CryptocurrencyListResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
