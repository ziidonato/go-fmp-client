package go_fmp

import (
	"encoding/json"
	"fmt"
)

// BatchCommodityQuotesParams represents the parameters for the All Commodities Quotes API
type BatchCommodityQuotesParams struct {
	Short bool `json:"short"` // Required: Whether to return short format (true) or full format (false)
}

// BatchCommodityQuotesResponse represents the response from the All Commodities Quotes API
type BatchCommodityQuotesResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
	Volume int64   `json:"volume"`
}

// BatchCommodityQuotes retrieves real-time quotes for multiple commodities at once
func (c *Client) BatchCommodityQuotes(params BatchCommodityQuotesParams) ([]BatchCommodityQuotesResponse, error) {
	urlParams := map[string]string{
		"short": fmt.Sprintf("%t", params.Short),
	}

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/batch-commodity-quotes", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []BatchCommodityQuotesResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
