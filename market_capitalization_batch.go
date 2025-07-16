package go_fmp

import (
	"encoding/json"
	"fmt"
)

// MarketCapitalizationBatchParams represents the parameters for the Batch Market Cap API
type MarketCapitalizationBatchParams struct {
	Symbols string `json:"symbols"` // Required: Comma-separated stock symbols (e.g., "AAPL,MSFT,GOOG")
}

// MarketCapitalizationBatchResponse represents the response from the Batch Market Cap API
type MarketCapitalizationBatchResponse struct {
	Symbol    string `json:"symbol"`
	Date      string `json:"date"`
	MarketCap int64  `json:"marketCap"`
}

// MarketCapitalizationBatch retrieves market capitalization data for multiple companies in a single request
func (c *Client) MarketCapitalizationBatch(params MarketCapitalizationBatchParams) ([]MarketCapitalizationBatchResponse, error) {
	if params.Symbols == "" {
		return nil, fmt.Errorf("symbols parameter is required")
	}

	urlParams := map[string]string{
		"symbols": params.Symbols,
	}

	return doRequest[[]MarketCapitalizationBatchResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
}
