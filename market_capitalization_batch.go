package go_fmp

import (
	"fmt"
	"time"
)

// MarketCapitalizationBatchParams represents the parameters for the Market Capitalization Batch API
type MarketCapitalizationBatchParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL,MSFT")
}

// MarketCapitalizationBatchResponse represents the response from the Market Capitalization Batch API
type MarketCapitalizationBatchResponse struct {
	Symbol    string    `json:"symbol"`
	Date      time.Time `json:"date"`
	MarketCap float64   `json:"marketCap"`
}

// MarketCapitalizationBatch retrieves market capitalization data for multiple companies in a single request
func (c *Client) MarketCapitalizationBatch(params MarketCapitalizationBatchParams) ([]MarketCapitalizationBatchResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	var result []MarketCapitalizationBatchResponse
	err := c.doRequest(c.BaseURL+"/market-capitalization-batch", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
