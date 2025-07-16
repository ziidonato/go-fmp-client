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
	return doRequest[[]DCFBulkResponse](c, "https://financialmodelingprep.com/stable/dcf-bulk", nil)
}
