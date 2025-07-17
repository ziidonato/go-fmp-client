package go_fmp

import (
	"fmt"
	"time"
)

// MarketCapitalizationParams represents the parameters for the Market Capitalization API
type MarketCapitalizationParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// MarketCapitalizationResponse represents the response from the Market Capitalization API
type MarketCapitalizationResponse struct {
	Symbol    string    `json:"symbol"`
	Date      time.Time `json:"date"`
	MarketCap float64   `json:"marketCap"`
}

// MarketCapitalization retrieves the market capitalization for a specific company on any given date
func (c *Client) MarketCapitalization(params MarketCapitalizationParams) ([]MarketCapitalizationResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	var result []MarketCapitalizationResponse
	err := c.doRequest(c.BaseURL+"/market-capitalization", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
