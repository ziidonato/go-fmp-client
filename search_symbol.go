package go_fmp

import (
	"fmt"
)

// SearchSymbolParams represents the parameters for the Stock Symbol Search API
type SearchSymbolParams struct {
	Query    string  `json:"query"`    // Required: Search query (e.g., "AAPL")
	Limit    *int    `json:"limit"`    // Optional: Number of results to return (default: 50)
	Exchange *Exchange `json:"exchange"` // Optional: Exchange filter (e.g., "NASDAQ")
}

// SearchSymbolResponse represents the response from the Stock Symbol Search API
type SearchSymbolResponse struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency Currency `json:"currency"`
	ExchangeFullName string `json:"exchangeFullName"`
	Exchange Exchange `json:"exchange"`
}

// SearchSymbol searches for stock symbols by query across multiple global markets
func (c *Client) SearchSymbol(params SearchSymbolParams) ([]SearchSymbolResponse, error) {
	if params.Query == "" {
		return nil, fmt.Errorf("query parameter is required")
	}

	urlParams := map[string]string{
		"query": params.Query,
	}

	if params.Limit != nil {
		urlParams["limit"] = fmt.Sprintf("%d", *params.Limit)
	}

	if params.Exchange != nil {
		urlParams["exchange"] = *params.Exchange
	}

	var result []SearchSymbolResponse
	err := c.doRequest(c.BaseURL+"/search-symbol", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, err
}
