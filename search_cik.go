package go_fmp

import (
	"fmt"
)

// SearchCIKParams represents the parameters for the CIK Search API
type SearchCIKParams struct {
	CIK   string `json:"cik"`   // Required: Central Index Key (e.g., "320193")
	Limit *int   `json:"limit"` // Optional: Number of results to return (default: 50)
}

// SearchCIKResponse represents the response from the CIK Search API
type SearchCIKResponse struct {
	Symbol           string `json:"symbol"`
	CompanyName      string `json:"companyName"`
	CIK              string `json:"cik"`
	ExchangeFullName string `json:"exchangeFullName"`
	Exchange Exchange `json:"exchange"`
	Currency Currency `json:"currency"`
}

// SearchCIK retrieves the Central Index Key (CIK) for publicly traded companies
func (c *Client) SearchCIK(params SearchCIKParams) ([]SearchCIKResponse, error) {
	if params.CIK == "" {
		return nil, fmt.Errorf("cik parameter is required")
	}

	urlParams := map[string]string{
		"cik": params.CIK,
	}

	if params.Limit != nil {
		urlParams["limit"] = fmt.Sprintf("%d", *params.Limit)
	}

	var result []SearchCIKResponse
	err := c.doRequest(c.BaseURL+"/search-cik", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
