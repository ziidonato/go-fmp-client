package go_fmp

import (
	"time"
	"fmt"
)

// ForexHistoricalLightParams represents the parameters for the Historical Forex Light Chart API
type ForexHistoricalLightParams struct {
	Symbol string  `json:"symbol"` // Required: Forex symbol (e.g., "EURUSD")
	From   *string `json:"from"`   // Optional: Start date (e.g., "2024-01-01")
	To     *string `json:"to"`     // Optional: End date (e.g., "2024-03-01")
}

// ForexHistoricalLightResponse represents the response from the Historical Forex Light Chart API
type ForexHistoricalLightResponse struct {
	Symbol string  `json:"symbol"`
	Date time.Time `json:"date"`
	Price  float64 `json:"price"`
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
