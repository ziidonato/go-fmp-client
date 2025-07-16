package go_fmp

import (
	"encoding/json"
	"fmt"
)

// SearchCUSIPParams represents the parameters for the CUSIP Search API
type SearchCUSIPParams struct {
	CUSIP string `json:"cusip"` // Required: CUSIP number (e.g., "037833100")

// SearchCUSIPResponse represents the response from the CUSIP Search API
type SearchCUSIPResponse struct {
	Symbol      string `json:"symbol"`
	CompanyName string `json:"companyName"`
	CUSIP       string `json:"cusip"`
	MarketCap   int64  `json:"marketCap"`

// SearchCUSIP searches and retrieves financial securities information by CUSIP number
func (c *Client) SearchCUSIP(params SearchCUSIPParams) ([]SearchCUSIPResponse, error) {
	if params.CUSIP == "" {
		return nil, fmt.Errorf("cusip parameter is required")
	}

	urlParams := map[string]string{
		"cusip": params.CUSIP,
	}

	return doRequest[[]SearchCUSIPResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
