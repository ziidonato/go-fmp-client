package go_fmp

// FinancialStatementSymbolListResponse represents the response from the Financial Statement Symbol List API
type FinancialStatementSymbolListResponse struct {
	Symbol string `json:"symbol"`
}

// FinancialStatementSymbolList retrieves a comprehensive list of financial symbols that have statements available
func (c *Client) FinancialStatementSymbolList() ([]FinancialStatementSymbolListResponse, error) {
	url := c.BaseURL + "/financial-statement-symbol-list"
	var result []FinancialStatementSymbolListResponse
	if err := c.doRequest(url, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
