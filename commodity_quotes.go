package go_fmp

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// CommodityQuotesResponse represents the response from the full commodities quotes API
type CommodityQuotesResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
	Volume int64   `json:"volume"`
}

// GetCommodityQuotes retrieves up-to-the-minute quotes for commodities
func (c *Client) GetCommodityQuotes(short bool) ([]CommodityQuotesResponse, error) {
	url := "https://financialmodelingprep.com/stable/batch-commodity-quotes"

	return doRequest[[]CommodityQuotesResponse](c, url, map[string]string{
		"short": strconv.FormatBool(short)
}
