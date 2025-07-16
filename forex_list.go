package go_fmp

// ForexListResponse represents the response from the Forex List API
type ForexListResponse struct {
	Symbol     string  `json:"symbol"`
	Name       string  `json:"name"`
	Currency   string  `json:"currency"`
	StockExchange string `json:"stockExchange"`
	ExchangeShortName string `json:"exchangeShortName"`
}

// ForexList retrieves a comprehensive list of all forex trading pairs
func (c *Client) ForexList() ([]ForexListResponse, error) {
	url := "https://financialmodelingprep.com/stable/forex-list"
	var result []ForexListResponse
	if err := c.doRequest(url, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
