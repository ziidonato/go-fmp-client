package go_fmp

import (
	"encoding/json"
	"fmt"
)

// MergersAcquisitionsSearchParams represents the parameters for the Search Mergers & Acquisitions API
type MergersAcquisitionsSearchParams struct {
	Name string `json:"name"` // Required: Company name to search for (e.g., "Apple")
}

// MergersAcquisitionsSearchResponse represents the response from the Search Mergers & Acquisitions API
type MergersAcquisitionsSearchResponse struct {
	Symbol              string `json:"symbol"`
	CompanyName         string `json:"companyName"`
	CIK                 string `json:"cik"`
	TargetedCompanyName string `json:"targetedCompanyName"`
	TargetedCIK         string `json:"targetedCik"`
	TargetedSymbol      string `json:"targetedSymbol"`
	TransactionDate     string `json:"transactionDate"`
	AcceptedDate        string `json:"acceptedDate"`
	Link                string `json:"link"`
}

// MergersAcquisitionsSearch searches for specific mergers and acquisitions data
func (c *Client) MergersAcquisitionsSearch(params MergersAcquisitionsSearchParams) ([]MergersAcquisitionsSearchResponse, error) {
	if params.Name == "" {
		return nil, fmt.Errorf("name parameter is required")
	}

	urlParams := map[string]string{
		"name": params.Name,
	}

	return doRequest[[]MergersAcquisitionsSearchResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
}
