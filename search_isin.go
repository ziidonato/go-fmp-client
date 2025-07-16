package go_fmp

import (
	"fmt"
)

// SearchISINParams represents the parameters for the ISIN Search API
type SearchISINParams struct {
	ISIN string `json:"isin"` // Required: International Securities Identification Number (e.g., "US0378331005")
}

// SearchISINResponse represents the response from the ISIN Search API
type SearchISINResponse struct {
	Symbol    string `json:"symbol"`
	Name      string `json:"name"`
	ISIN      string `json:"isin"`
	MarketCap int64  `json:"marketCap"`
}

// SearchISIN searches and retrieves the International Securities Identification Number (ISIN) for financial securities
func (c *Client) SearchISIN(params SearchISINParams) ([]SearchISINResponse, error) {
	if params.ISIN == "" {
		return nil, fmt.Errorf("isin parameter is required")
	}

	urlParams := map[string]string{
		"isin": params.ISIN,
	}

	var result []SearchISINResponse
	err := c.doRequest(c.BaseURL+"/search-isin", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, err
}
