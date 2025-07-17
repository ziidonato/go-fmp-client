package go_fmp

import "fmt"

// StockListResponse represents the response from the Company Symbols List API
type StockListResponse struct {
	Symbol      string `json:"symbol"`
	CompanyName string `json:"companyName"`
}

// StockList retrieves a comprehensive list of financial symbols from various global exchanges
func (c *Client) StockList() ([]StockListResponse, error) {
	var result []StockListResponse
	err := c.doRequest(c.BaseURL+"/stock-list", map[string]string{}, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
