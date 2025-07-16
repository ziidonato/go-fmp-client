package go_fmp

import (
	"strconv"
)

// CryptoQuotesResponse represents the response from the full cryptocurrency quotes API
type CryptoQuotesResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
	Volume int64   `json:"volume"`
}

// CryptoQuotes retrieves real-time cryptocurrency quotes
func (c *Client) CryptoQuotes(short bool) ([]CryptoQuotesResponse, error) {
	url := "https://financialmodelingprep.com/stable/batch-crypto-quotes"
	params := map[string]string{"short": strconv.FormatBool(short)}
	var result []CryptoQuotesResponse
	if err := c.doRequest(url, params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
