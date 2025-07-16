package go_fmp

import (
	"encoding/json"
	"fmt"
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

	return doRequest[[]CryptoQuotesResponse](c, url, map[string]string{
		"short": strconv.FormatBool(short)
}
