package go_fmp

import (
	"fmt"
)

// ForexQuoteShortParams represents the parameters for the Forex Quote Short API
type ForexQuoteShortParams struct {
	Symbol string `json:"symbol"` // Required: Forex symbol (e.g., "EURUSD")
}

// ForexQuoteShortResponse represents the response from the Forex Quote Short API
type ForexQuoteShortResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
	Volume int64   `json:"volume"`
}

// ForexQuoteShort retrieves concise forex pair quotes with live currency exchange rates
func (c *Client) ForexQuoteShort(params ForexQuoteShortParams) ([]ForexQuoteShortResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	var result []ForexQuoteShortResponse
	err := c.doRequest(c.BaseURL+"/quote-short", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
