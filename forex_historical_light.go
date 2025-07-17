package go_fmp

import (
	"fmt"
	"time"
)

// ForexHistoricalLightParams represents the parameters for the Forex Historical Light API
type ForexHistoricalLightParams struct {
	Symbol string  `json:"symbol"` // Required: Forex pair symbol (e.g., "EURUSD")
	From   *string `json:"from"`   // Optional: Start date (YYYY-MM-DD format)
	To     *string `json:"to"`     // Optional: End date (YYYY-MM-DD format)
}

// ForexHistoricalLightResponse represents the response from the Forex Historical Light API
type ForexHistoricalLightResponse struct {
	Date   time.Time  `json:"date"`
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
}

// ForexHistoricalLight retrieves historical end-of-day forex prices for currency pairs
func (c *Client) ForexHistoricalLight(params ForexHistoricalLightParams) ([]ForexHistoricalLightResponse, error) {
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

	var result []ForexHistoricalLightResponse
	err := c.doRequest(c.BaseURL+"/historical-price-eod/light", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
