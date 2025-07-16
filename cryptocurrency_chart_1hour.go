package go_fmp

import (
	"encoding/json"
	"fmt"
)

// CryptocurrencyChart1HourParams represents the parameters for the 1-Hour Cryptocurrency Chart API
type CryptocurrencyChart1HourParams struct {
	Symbol string  `json:"symbol"` // Required: Cryptocurrency symbol (e.g., "BTCUSD")
	From   *string `json:"from"`   // Optional: Start date (e.g., "2024-01-01")
	To     *string `json:"to"`     // Optional: End date (e.g., "2024-03-01")

// CryptocurrencyChart1HourResponse represents the response from the 1-Hour Cryptocurrency Chart API
type CryptocurrencyChart1HourResponse struct {
	Date   string  `json:"date"`
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`

// CryptocurrencyChart1Hour retrieves real-time, 1-hour interval price data for cryptocurrencies
func (c *Client) CryptocurrencyChart1Hour(params CryptocurrencyChart1HourParams) ([]CryptocurrencyChart1HourResponse, error) {
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

	return doRequest[[]CryptocurrencyChart1HourResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
