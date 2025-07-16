package go_fmp

import (
	"encoding/json"
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

	resp, err := c.get("https://financialmodelingprep.com/stable/quote-short", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []ForexQuoteShortResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
