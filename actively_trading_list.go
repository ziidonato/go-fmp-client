package go_fmp

// ActivelyTradingListResponse represents the response from the Actively Trading List API
type ActivelyTradingListResponse struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

// ActivelyTradingList retrieves all actively trading companies and financial instruments
func (c *Client) ActivelyTradingList() ([]ActivelyTradingListResponse, error) {
	url := "https://financialmodelingprep.com/stable/actively-trading-list"
	var result []ActivelyTradingListResponse
	if err := c.doRequest(url, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
