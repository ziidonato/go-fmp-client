package go_fmp

import (
	"encoding/json"
	"fmt"
)

// CommoditiesQuoteShortParams represents the parameters for the Commodities Quote Short API
type CommoditiesQuoteShortParams struct {
	Symbol string `json:"symbol"` // Required: Commodity symbol (e.g., "GCUSD")
}

// CommoditiesQuoteShortResponse represents the response from the Commodities Quote Short API
type CommoditiesQuoteShortResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
	Volume int64   `json:"volume"`
}

// CommoditiesQuoteShort retrieves fast and accurate quotes for commodities
func (c *Client) CommoditiesQuoteShort(params CommoditiesQuoteShortParams) ([]CommoditiesQuoteShortResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	resp, err := c.get("https://financialmodelingprep.com/stable/quote-short", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []CommoditiesQuoteShortResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
