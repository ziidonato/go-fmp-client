package go_fmp

import (
	"encoding/json"
	"fmt"
)

// BatchForexQuotesParams represents the parameters for the Batch Forex Quotes API
type BatchForexQuotesParams struct {
	Short bool `json:"short"` // Required: Whether to return short format (true) or full format (false)
}

// BatchForexQuotesResponse represents the response from the Batch Forex Quotes API
type BatchForexQuotesResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
	Volume int64   `json:"volume"`
}

// BatchForexQuotes retrieves real-time quotes for multiple forex pairs simultaneously
func (c *Client) BatchForexQuotes(params BatchForexQuotesParams) ([]BatchForexQuotesResponse, error) {
	urlParams := map[string]string{
		"short": fmt.Sprintf("%t", params.Short),
	}

	return doRequest[[]BatchForexQuotesResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
}
