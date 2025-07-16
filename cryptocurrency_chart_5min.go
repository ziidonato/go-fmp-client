package go_fmp

import (
	"encoding/json"
	"fmt"
)

// CryptocurrencyChart5MinParams represents the parameters for the 5-Minute Cryptocurrency Chart API
type CryptocurrencyChart5MinParams struct {
	Symbol string  `json:"symbol"` // Required: Cryptocurrency symbol (e.g., "BTCUSD")
	From   *string `json:"from"`   // Optional: Start date (e.g., "2024-01-01")
	To     *string `json:"to"`     // Optional: End date (e.g., "2024-03-01")
}

// CryptocurrencyChart5MinResponse represents the response from the 5-Minute Cryptocurrency Chart API
type CryptocurrencyChart5MinResponse struct {
	Date   string  `json:"date"`
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
}

// CryptocurrencyChart5Min retrieves real-time, 5-minute interval price data for cryptocurrencies
func (c *Client) CryptocurrencyChart5Min(params CryptocurrencyChart5MinParams) ([]CryptocurrencyChart5MinResponse, error) {
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

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/historical-chart/5min", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []CryptocurrencyChart5MinResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
