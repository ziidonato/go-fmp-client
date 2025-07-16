package go_fmp

import (
	"encoding/json"
	"fmt"
)

// StockPriceChangeResponse represents the response from the stock price change API
type StockPriceChangeResponse struct {
	Symbol string  `json:"symbol"`
	OneD   float64 `json:"1D"`
	FiveD  float64 `json:"5D"`
	OneM   float64 `json:"1M"`
	ThreeM float64 `json:"3M"`
	SixM   float64 `json:"6M"`
	YTD    float64 `json:"ytd"`
	OneY   float64 `json:"1Y"`
	ThreeY float64 `json:"3Y"`
	FiveY  float64 `json:"5Y"`
	TenY   float64 `json:"10Y"`
	Max    float64 `json:"max"`
}

// GetStockPriceChange retrieves stock price fluctuations in real-time
func (c *Client) GetStockPriceChange(symbol string) ([]StockPriceChangeResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	url := "https://financialmodelingprep.com/stable/stock-price-change"

	resp, err := c.doRequest(url, map[string]string{
		"symbol": symbol,
	})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()


	var result []StockPriceChangeResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
