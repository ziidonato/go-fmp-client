package go_fmp

import (
	"fmt"
	"time"
)

// ForexHistoricalFullParams represents the parameters for the Forex Historical Full API
type ForexHistoricalFullParams struct {
	Symbol string  `json:"symbol"` // Required: Forex pair symbol (e.g., "EURUSD")
	From   *string `json:"from"`   // Optional: Start date (YYYY-MM-DD format)
	To     *string `json:"to"`     // Optional: End date (YYYY-MM-DD format)
}

// ForexHistoricalFullResponse represents the response from the Forex Historical Full API
type ForexHistoricalFullResponse struct {
	Date          time.Time  `json:"date"`
	Open          float64 `json:"open"`
	Low           float64 `json:"low"`
	High          float64 `json:"high"`
	Close         float64 `json:"close"`
	AdjClose      float64 `json:"adjClose"`
	Volume        int64   `json:"volume"`
	UnadjustedVolume int64 `json:"unadjustedVolume"`
	Change        float64 `json:"change"`
	ChangePercent float64 `json:"changePercent"`
	Vwap          float64 `json:"vwap"`
	Label         string  `json:"label"`
	ChangeOverTime float64 `json:"changeOverTime"`
}

// ForexHistoricalFull retrieves comprehensive historical end-of-day forex price data for currency pairs
func (c *Client) ForexHistoricalFull(params ForexHistoricalFullParams) ([]ForexHistoricalFullResponse, error) {
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

	var result []ForexHistoricalFullResponse
	err := c.doRequest(c.BaseURL+"/historical-price-eod/full", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
