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

	resp, err := c.doRequest(url, map[string]string{
		"short": strconv.FormatBool(short),
	})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()


	var result []CommodityQuotesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
