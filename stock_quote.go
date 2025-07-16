package go_fmp

import (
	"encoding/json"
	"fmt"
)

// StockQuoteResponse represents the response from the stock quote API
type StockQuoteResponse struct {
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
	MarketCap        int64   `json:"marketCap"`
	PriceAvg50       float64 `json:"priceAvg50"`
	PriceAvg200      float64 `json:"priceAvg200"`
	Exchange         string  `json:"exchange"`
	Open             float64 `json:"open"`
	PreviousClose    float64 `json:"previousClose"`
	Timestamp        int64   `json:"timestamp"`

// GetStockQuote retrieves real-time stock quotes
func (c *Client) GetStockQuote(symbol string) ([]StockQuoteResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	return doRequest[[]StockQuoteResponse](c, "https://financialmodelingprep.com/stable/quote", map[string]string{
		"symbol": symbol,
	}
