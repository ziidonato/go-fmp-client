package go_fmp

// CryptocurrencyListResponse represents the response from the Cryptocurrency List API
type CryptocurrencyListResponse struct {
	Symbol       string  `json:"symbol"`
	Name         string  `json:"name"`
	Currency Currency `json:"currency"`
	Exchange     *string `json:"exchange"`
	ExchangeName *string `json:"exchangeName"`
}

// CryptocurrencyList retrieves a comprehensive list of all cryptocurrencies traded on exchanges worldwide
func (c *Client) CryptocurrencyList() ([]CryptocurrencyListResponse, error) {
	url := c.BaseURL + "/cryptocurrency-list"
	var result []CryptocurrencyListResponse
	if err := c.doRequest(url, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
