package go_fmp

import (
	"encoding/json"
	"fmt"
)

// CryptocurrencyQuoteShortParams represents the parameters for the Cryptocurrency Quote Short API
type CryptocurrencyQuoteShortParams struct {
	Symbol string `json:"symbol"` // Required: Cryptocurrency symbol (e.g., "BTCUSD")
}

// CryptocurrencyQuoteShortResponse represents the response from the Cryptocurrency Quote Short API
type CryptocurrencyQuoteShortResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
	Volume int64   `json:"volume"`
}

// CryptocurrencyQuoteShort retrieves concise cryptocurrency quotes with current prices, changes, and trading volume
func (c *Client) CryptocurrencyQuoteShort(params CryptocurrencyQuoteShortParams) ([]CryptocurrencyQuoteShortResponse, error) {
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

	var result []CryptocurrencyQuoteShortResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
