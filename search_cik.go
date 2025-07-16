package go_fmp

import (
	"encoding/json"
	"fmt"
)

// SearchCIKParams represents the parameters for the CIK Search API
type SearchCIKParams struct {
	CIK   string `json:"cik"`   // Required: Central Index Key (e.g., "320193")
	Limit *int   `json:"limit"` // Optional: Number of results to return (default: 50)

// SearchCIKResponse represents the response from the CIK Search API
type SearchCIKResponse struct {
	Symbol           string `json:"symbol"`
	CompanyName      string `json:"companyName"`
	CIK              string `json:"cik"`
	ExchangeFullName string `json:"exchangeFullName"`
	Exchange         string `json:"exchange"`
	Currency         string `json:"currency"`

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

	return doRequest[[]SearchCIKResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
