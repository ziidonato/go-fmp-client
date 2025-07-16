package go_fmp

import (
	"encoding/json"
	"fmt"
)

// ForexQuoteParams represents the parameters for the Forex Quote API
type ForexQuoteParams struct {
	Symbol string `json:"symbol"` // Required: Forex symbol (e.g., "EURUSD")
}

// ForexQuoteResponse represents the response from the Forex Quote API
type ForexQuoteResponse struct {
	Symbol           string  `json:"symbol"`
	Name             string  `json:"name"`
	Price            float64 `json:"price"`
	ChangePercentage float64 `json:"changePercentage"`
	Change           float64 `json:"change"`
	Volume           int64   `json:"volume"`
	DayLow           float64 `json:"dayLow"`
	DayHigh          float64 `json:"dayHigh"`
	YearHigh         float64 `json:"yearHigh"`
	YearLow          float64 `json:"yearLow"`
	MarketCap        *int64  `json:"marketCap"`
	PriceAvg50       float64 `json:"priceAvg50"`
	PriceAvg200      float64 `json:"priceAvg200"`
	Exchange         string  `json:"exchange"`
	Open             float64 `json:"open"`
	PreviousClose    float64 `json:"previousClose"`
	Timestamp        int64   `json:"timestamp"`
}

// ForexQuote retrieves real-time forex quotes for currency pairs
func (c *Client) ForexQuote(params ForexQuoteParams) ([]ForexQuoteResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	return doRequest[[]ForexQuoteResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
}
