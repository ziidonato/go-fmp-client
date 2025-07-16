package go_fmp

import (
	"fmt"
)

// HistoricalIndexLightResponse represents the response from the historical index light chart API
type HistoricalIndexLightResponse struct {
	Symbol string  `json:"symbol"`
	Date   string  `json:"date"`
	Price  float64 `json:"price"`
	Volume int64   `json:"volume"`
}

// GetHistoricalIndexLight retrieves end-of-day historical prices for stock indexes
func (c *Client) GetHistoricalIndexLight(symbol, from, to string) ([]HistoricalIndexLightResponse, error) {
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

	url := "https://financialmodelingprep.com/stable/historical-price-eod/light"

	var result []HistoricalIndexLightResponse
	err := c.doRequest(url, params, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
