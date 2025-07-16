package go_fmp

import (
	"encoding/json"
	"fmt"
)

// ForexChart5MinParams represents the parameters for the 5-Minute Forex Chart API
type ForexChart5MinParams struct {
	Symbol string  `json:"symbol"` // Required: Forex symbol (e.g., "EURUSD")
	From   *string `json:"from"`   // Optional: Start date (e.g., "2024-01-01")
	To     *string `json:"to"`     // Optional: End date (e.g., "2024-03-01")

// ForexChart5MinResponse represents the response from the 5-Minute Forex Chart API
type ForexChart5MinResponse struct {
	Date   string  `json:"date"`
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`

// ForexChart5Min retrieves real-time, 5-minute intraday forex data for currency pairs
func (c *Client) ForexChart5Min(params ForexChart5MinParams) ([]ForexChart5MinResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	if params.From != nil {
		urlParams["from"] = *params.From
	}

	if params.To != nil {
		urlParams["to"] = *params.To
	}

	return doRequest[[]ForexChart5MinResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
