package go_fmp

import (
	"fmt"
	"time"
)

// SymbolChangeParams represents the parameters for the Symbol Change API
type SymbolChangeParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// SymbolChangeResponse represents the response from the Symbol Change API
type SymbolChangeResponse struct {
	Date        time.Time `json:"date"`
	Name        string    `json:"name"`
	OldSymbol   string    `json:"oldSymbol"`
	NewSymbol   string    `json:"newSymbol"`
}

// SymbolChange retrieves information about stock symbol changes due to mergers, acquisitions, stock splits, and name changes
func (c *Client) SymbolChange(params SymbolChangeParams) ([]SymbolChangeResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	var result []SymbolChangeResponse
	err := c.doRequest(c.BaseURL+"/symbol-change", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, err
}
