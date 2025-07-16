package go_fmp

import (
	"encoding/json"
	"fmt"
)

// SymbolChangeParams represents the parameters for the Symbol Changes List API
type SymbolChangeParams struct {
	Invalid *string `json:"invalid"` // Required: Filter for invalid symbols (e.g., "false")
	Limit   *int    `json:"limit"`   // Required: Number of results (e.g., 100)
}

// SymbolChangeResponse represents the response from the Symbol Changes List API
type SymbolChangeResponse struct {
	Date        string `json:"date"`
	CompanyName string `json:"companyName"`
	OldSymbol   string `json:"oldSymbol"`
	NewSymbol   string `json:"newSymbol"`
}

// SymbolChange retrieves information about stock symbol changes due to mergers, acquisitions, stock splits, and name changes
func (c *Client) SymbolChange(params SymbolChangeParams) ([]SymbolChangeResponse, error) {
	if params.Invalid == nil {
		return nil, fmt.Errorf("invalid parameter is required")
	}

	if params.Limit == nil {
		return nil, fmt.Errorf("limit parameter is required")
	}

	urlParams := map[string]string{
		"invalid": *params.Invalid,
		"limit":   fmt.Sprintf("%d", *params.Limit),
	}

	resp, err := c.get("https://financialmodelingprep.com/stable/symbol-change", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []SymbolChangeResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
