package go_fmp

import (
	"fmt"
)

// ETFListResponse represents the response from the ETF Symbol Search API
type ETFListResponse struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

// ETFList retrieves ticker symbols and company names for Exchange Traded Funds (ETFs)
func (c *Client) ETFList() ([]ETFListResponse, error) {
	url := c.BaseURL + "/etf-list"

	var result []ETFListResponse
	err := c.doRequest(url, map[string]string{}, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
