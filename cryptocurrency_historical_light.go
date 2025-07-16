package go_fmp

import (
	"encoding/json"
	"fmt"
)

// CryptocurrencyHistoricalLightParams represents the parameters for the Historical Cryptocurrency Light Chart API
type CryptocurrencyHistoricalLightParams struct {
	Symbol string  `json:"symbol"` // Required: Cryptocurrency symbol (e.g., "BTCUSD")
	From   *string `json:"from"`   // Optional: Start date (e.g., "2024-01-01")
	To     *string `json:"to"`     // Optional: End date (e.g., "2024-03-01")
}

// CryptocurrencyHistoricalLightResponse represents the response from the Historical Cryptocurrency Light Chart API
type CryptocurrencyHistoricalLightResponse struct {
	Symbol string  `json:"symbol"`
	Date   string  `json:"date"`
	Price  float64 `json:"price"`
	Volume int64   `json:"volume"`
}

// CryptocurrencyHistoricalLight retrieves historical end-of-day prices for cryptocurrencies
func (c *Client) CryptocurrencyHistoricalLight(params CryptocurrencyHistoricalLightParams) ([]CryptocurrencyHistoricalLightResponse, error) {
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

	resp, err := c.get("https://financialmodelingprep.com/stable/historical-price-eod/light", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []CryptocurrencyHistoricalLightResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
