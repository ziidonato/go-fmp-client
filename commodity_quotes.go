package go_fmp

import (
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
	url := c.BaseURL + "/batch-commodity-quotes"

	var result []CommodityQuotesResponse
	err := c.doRequest(url, map[string]string{
		"short": strconv.FormatBool(short),
	}, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
