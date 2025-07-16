package go_fmp

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// ETFQuotesResponse represents the response from the ETF price quotes API
type ETFQuotesResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
	Volume int64   `json:"volume"`

// GetETFQuotes retrieves real-time price quotes for exchange-traded funds (ETFs)
func (c *Client) GetETFQuotes(short bool) ([]ETFQuotesResponse, error) {
	url := "https://financialmodelingprep.com/stable/batch-etf-quotes"

	return doRequest[[]ETFQuotesResponse](c, url, map[string]string{
		"short": strconv.FormatBool(short)
	}
