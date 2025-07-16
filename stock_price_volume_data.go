package go_fmp

import (
	"encoding/json"
	"fmt"
)

// StockPriceVolumeDataParams represents the parameters for the Stock Price and Volume Data API
type StockPriceVolumeDataParams struct {
	Symbol string  `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	From   *string `json:"from"`   // Optional: Start date (e.g., "2025-01-10")
	To     *string `json:"to"`     // Optional: End date (e.g., "2025-04-10")
}

// StockPriceVolumeDataResponse represents the response from the Stock Price and Volume Data API
type StockPriceVolumeDataResponse struct {
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
}

// StockPriceVolumeData retrieves full price and volume data for any stock symbol
func (c *Client) StockPriceVolumeData(params StockPriceVolumeDataParams) ([]StockPriceVolumeDataResponse, error) {
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

	var result []StockPriceVolumeDataResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
