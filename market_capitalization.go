package go_fmp

import (
	"encoding/json"
	"fmt"
)

// MarketCapitalizationParams represents the parameters for the Company Market Cap API
type MarketCapitalizationParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// MarketCapitalizationResponse represents the response from the Company Market Cap API
type MarketCapitalizationResponse struct {
	Symbol    string `json:"symbol"`
	Date      string `json:"date"`
	MarketCap int64  `json:"marketCap"`
}

// MarketCapitalization retrieves the market capitalization for a specific company on any given date
func (c *Client) MarketCapitalization(params MarketCapitalizationParams) ([]MarketCapitalizationResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	return doRequest[[]MarketCapitalizationResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
}
