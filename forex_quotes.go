package go_fmp

import (
	"strconv"
)

// ForexQuotesResponse represents the response from the full forex quote API
type ForexQuotesResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
	Volume int64   `json:"volume"`
}

// ForexQuotes retrieves real-time quotes for multiple forex currency pairs
func (c *Client) ForexQuotes(short bool) ([]ForexQuotesResponse, error) {
	url := "https://financialmodelingprep.com/stable/batch-forex-quotes"

	var result []ForexQuotesResponse
	if err := c.doRequest(url, map[string]string{"short": strconv.FormatBool(short)}, &result); err != nil {
		return nil, err
	}
	return result, nil
}
