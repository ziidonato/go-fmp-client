package go_fmp

import (
	"encoding/json"
	"fmt"
)

// SharesFloatParams represents the parameters for the Company Share Float & Liquidity API
type SharesFloatParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")

// SharesFloatResponse represents the response from the Company Share Float & Liquidity API
type SharesFloatResponse struct {
	Symbol            string  `json:"symbol"`
	Date              string  `json:"date"`
	FreeFloat         float64 `json:"freeFloat"`
	FloatShares       int64   `json:"floatShares"`
	OutstandingShares int64   `json:"outstandingShares"`

// SharesFloat retrieves the total number of publicly traded shares for any company
func (c *Client) SharesFloat(params SharesFloatParams) ([]SharesFloatResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	return doRequest[[]SharesFloatResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
