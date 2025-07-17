package go_fmp

import (
	"fmt"
)

// SearchNameParams represents the parameters for the Company Name Search API
type SearchNameParams struct {
	Query    string  `json:"query"`    // Required: Search query (e.g., "AA")
	Limit    *int    `json:"limit"`    // Optional: Number of results to return (default: 50)
	Exchange *Exchange `json:"exchange"` // Optional: Exchange filter (e.g., "NASDAQ")
}

// SearchNameResponse represents the response from the Company Name Search API
type SearchNameResponse struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency Currency `json:"currency"`
	ExchangeFullName string `json:"exchangeFullName"`
	Exchange Exchange `json:"exchange"`
}

// SearchName searches for ticker symbols, company names, and exchange details for equity securities and ETFs
func (c *Client) SearchName(params SearchNameParams) ([]SearchNameResponse, error) {
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

	var result []SearchNameResponse
	err := c.doRequest(c.BaseURL+"/search-name", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, err
}
