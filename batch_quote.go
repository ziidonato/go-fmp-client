package go_fmp

import "fmt"

// BatchQuoteResponse represents the response from the stock batch quote API
type BatchQuoteResponse struct {
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
	Exchange Exchange `json:"exchange"`
	Open             float64 `json:"open"`
	PreviousClose    float64 `json:"previousClose"`
	Timestamp        int64   `json:"timestamp"`
}

// GetBatchQuote retrieves multiple real-time stock quotes in a single request
func (c *Client) GetBatchQuote(symbols string) ([]BatchQuoteResponse, error) {
	if symbols == "" {
		return nil, fmt.Errorf("symbols is required")
	}

	url := c.BaseURL + "/batch-quote"

	var result []BatchQuoteResponse
	err := c.doRequest(url, map[string]string{
		"symbols": symbols,
	}, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
