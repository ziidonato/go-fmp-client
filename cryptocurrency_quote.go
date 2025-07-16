package go_fmp

import (
	"encoding/json"
	"fmt"
)

// CryptocurrencyQuoteParams represents the parameters for the Cryptocurrency Quote API
type CryptocurrencyQuoteParams struct {
	Symbol string `json:"symbol"` // Required: Cryptocurrency symbol (e.g., "BTCUSD")

// CryptocurrencyQuoteResponse represents the response from the Cryptocurrency Quote API
type CryptocurrencyQuoteResponse struct {
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

// CryptocurrencyQuote retrieves real-time quotes for cryptocurrencies
func (c *Client) CryptocurrencyQuote(params CryptocurrencyQuoteParams) ([]CryptocurrencyQuoteResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	return doRequest[[]CryptocurrencyQuoteResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
