package go_fmp

// InsiderTradingTransactionTypeResponse represents the response from the insider trading transaction type API
type InsiderTradingTransactionTypeResponse struct {
	TransactionType string `json:"transactionType"`
}

// InsiderTradingTransactionType retrieves all available transaction types for insider trading data
func (c *Client) InsiderTradingTransactionType() ([]InsiderTradingTransactionTypeResponse, error) {
	url := "https://financialmodelingprep.com/stable/insider-trading-transaction-type"
	var result []InsiderTradingTransactionTypeResponse
	if err := c.doRequest(url, map[string]string{}, &result); err != nil {
		return nil, err
	}
	return result, nil
}
