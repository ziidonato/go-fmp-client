package go_fmp

import (
	"fmt"
	"time"
)

// CryptoHistoricalLightParams represents the parameters for the Crypto Historical Light API
type CryptoHistoricalLightParams struct {
	Symbol string  `json:"symbol"` // Required: Crypto symbol (e.g., "BTCUSD")
	From   *string `json:"from"`   // Optional: Start date (YYYY-MM-DD format)
	To     *string `json:"to"`     // Optional: End date (YYYY-MM-DD format)
}

// CryptoHistoricalLightResponse represents the response from the Crypto Historical Light API
type CryptoHistoricalLightResponse struct {
	Date   time.Time  `json:"date"`
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
	Volume float64 `json:"volume"`
}

// CryptoHistoricalLight retrieves light historical price data for cryptocurrencies
func (c *Client) CryptoHistoricalLight(params CryptoHistoricalLightParams) ([]CryptoHistoricalLightResponse, error) {
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

	url := fmt.Sprintf("%s/historical-price-full/%s", c.BaseURL, params.Symbol)
	var result []CryptoHistoricalLightResponse
	err := c.doRequest(url, urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}