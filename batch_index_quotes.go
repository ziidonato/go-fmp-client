package go_fmp

import (
	"strconv"
)

// BatchIndexQuoteResponse represents the response from the batch index quotes API
type BatchIndexQuoteResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
	Volume int64   `json:"volume"`
}

// GetBatchIndexQuotes retrieves real-time quotes for multiple stock indices
func (c *Client) GetBatchIndexQuotes(short bool) ([]BatchIndexQuoteResponse, error) {
	url := "https://financialmodelingprep.com/stable/batch-index-quotes"
	params := map[string]string{"short": strconv.FormatBool(short)}
	var result []BatchIndexQuoteResponse
	if err := c.doRequest(url, params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
