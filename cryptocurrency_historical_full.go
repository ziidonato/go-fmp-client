package go_fmp

import (
	"encoding/json"
	"fmt"
)

// CryptocurrencyHistoricalFullParams represents the parameters for the Historical Cryptocurrency Full Chart API
type CryptocurrencyHistoricalFullParams struct {
	Symbol string  `json:"symbol"` // Required: Cryptocurrency symbol (e.g., "BTCUSD")
	From   *string `json:"from"`   // Optional: Start date (e.g., "2024-01-01")
	To     *string `json:"to"`     // Optional: End date (e.g., "2024-03-01")
}

// CryptocurrencyHistoricalFullResponse represents the response from the Historical Cryptocurrency Full Chart API
type CryptocurrencyHistoricalFullResponse struct {
	Symbol        string  `json:"symbol"`
	Date          string  `json:"date"`
	Open          float64 `json:"open"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	Close         float64 `json:"close"`
	Volume        int64   `json:"volume"`
	Change        float64 `json:"change"`
	ChangePercent float64 `json:"changePercent"`
	Vwap          float64 `json:"vwap"`
}

// CryptocurrencyHistoricalFull retrieves comprehensive historical end-of-day data for cryptocurrencies
func (c *Client) CryptocurrencyHistoricalFull(params CryptocurrencyHistoricalFullParams) ([]CryptocurrencyHistoricalFullResponse, error) {
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

	resp, err := c.get("https://financialmodelingprep.com/stable/historical-price-eod/full", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []CryptocurrencyHistoricalFullResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
