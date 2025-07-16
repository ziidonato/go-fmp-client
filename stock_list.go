package go_fmp

import (
	"encoding/json"
)

// StockListResponse represents the response from the Company Symbols List API
type StockListResponse struct {
	Symbol      string `json:"symbol"`
	CompanyName string `json:"companyName"`
}

// StockList retrieves a comprehensive list of financial symbols from various global exchanges
func (c *Client) StockList() ([]StockListResponse, error) {
	resp, err := c.get("https://financialmodelingprep.com/stable/stock-list", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []StockListResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
