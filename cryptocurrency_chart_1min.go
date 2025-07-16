package go_fmp

import (
	"encoding/json"
	"fmt"
)

// CryptocurrencyChart1MinParams represents the parameters for the 1-Minute Cryptocurrency Chart API
type CryptocurrencyChart1MinParams struct {
	Symbol string  `json:"symbol"` // Required: Cryptocurrency symbol (e.g., "BTCUSD")
	From   *string `json:"from"`   // Optional: Start date (e.g., "2024-01-01")
	To     *string `json:"to"`     // Optional: End date (e.g., "2024-03-01")
}

// CryptocurrencyChart1MinResponse represents the response from the 1-Minute Cryptocurrency Chart API
type CryptocurrencyChart1MinResponse struct {
	Date   string  `json:"date"`
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
}

// CryptocurrencyChart1Min retrieves real-time, 1-minute interval price data for cryptocurrencies
func (c *Client) CryptocurrencyChart1Min(params CryptocurrencyChart1MinParams) ([]CryptocurrencyChart1MinResponse, error) {
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

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/historical-chart/1min", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []CryptocurrencyChart1MinResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
