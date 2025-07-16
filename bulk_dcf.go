package go_fmp

import (
	"encoding/json"
	"fmt"
)

// DCFBulkResponse represents the response from the DCF Bulk API
type DCFBulkResponse struct {
	Symbol     string `json:"symbol"`
	Date       string `json:"date"`
	DCF        string `json:"dcf"`
	StockPrice string `json:"Stock Price"`
}

// GetDCFBulk retrieves discounted cash flow valuations for multiple symbols
func (c *Client) GetDCFBulk() ([]DCFBulkResponse, error) {
	resp, err := c.doRequest("https://financialmodelingprep.com/stable/dcf-bulk", nil)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()



	// Parse the response
	var result []DCFBulkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
