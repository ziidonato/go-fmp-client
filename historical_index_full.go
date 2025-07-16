package go_fmp

import (
	"encoding/json"
	"fmt"
)

// HistoricalIndexFullResponse represents the response from the historical index full chart API
type HistoricalIndexFullResponse struct {
	Symbol        string  `json:"symbol"`
	Date          string  `json:"date"`
	Open          float64 `json:"open"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	Close         float64 `json:"close"`
	Volume        int64   `json:"volume"`
	Change        float64 `json:"change"`
	ChangePercent float64 `json:"changePercent"`
	VWAP          float64 `json:"vwap"`

// HistoricalIndexFull retrieves full historical end-of-day prices for stock indexes
func (c *Client) HistoricalIndexFull(symbol, from, to string) ([]HistoricalIndexFullResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	params := map[string]string{
		"symbol": symbol,
	}

	if from != "" {
		params["from"] = from
	}
	if to != "" {
		params["to"] = to
	}

	url := "https://financialmodelingprep.com/stable/historical-price-eod/full"

	return doRequest[[]HistoricalIndexFullResponse](c, url, params)
