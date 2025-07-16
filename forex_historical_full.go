package go_fmp

import (
	"encoding/json"
	"fmt"
)

// ForexHistoricalFullParams represents the parameters for the Historical Forex Full Chart API
type ForexHistoricalFullParams struct {
	Symbol string  `json:"symbol"` // Required: Forex symbol (e.g., "EURUSD")
	From   *string `json:"from"`   // Optional: Start date (e.g., "2024-01-01")
	To     *string `json:"to"`     // Optional: End date (e.g., "2024-03-01")
}

// ForexHistoricalFullResponse represents the response from the Historical Forex Full Chart API
type ForexHistoricalFullResponse struct {
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

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/historical-price-eod/full", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []ForexHistoricalFullResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
