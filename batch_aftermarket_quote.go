package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// BatchAftermarketQuoteResponse represents the response from the batch aftermarket quote API
type BatchAftermarketQuoteResponse struct {
	Symbol    string  `json:"symbol"`
	BidSize   int64   `json:"bidSize"`
	BidPrice  float64 `json:"bidPrice"`
	AskSize   int64   `json:"askSize"`
	AskPrice  float64 `json:"askPrice"`
	Volume    int64   `json:"volume"`
	Timestamp int64   `json:"timestamp"`
}

// GetBatchAftermarketQuote retrieves real-time aftermarket quotes for multiple stocks
func (c *Client) GetBatchAftermarketQuote(symbols string) ([]BatchAftermarketQuoteResponse, error) {
	if symbols == "" {
		return nil, fmt.Errorf("symbols is required")
	}

	url := "https://financialmodelingprep.com/stable/batch-aftermarket-quote"

	resp, err := c.get(url, map[string]string{
		"symbols": symbols,
	})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var result []BatchAftermarketQuoteResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
