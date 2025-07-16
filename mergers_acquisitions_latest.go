package go_fmp

import (
	"fmt"
)

// MergersAcquisitionsLatestParams represents the parameters for the Latest Mergers & Acquisitions API
type MergersAcquisitionsLatestParams struct {
	Page  *int `json:"page"`  // Required: Page number (e.g., 0)
	Limit *int `json:"limit"` // Required: Number of results (Maximum 1000 records per request)
}

// MergersAcquisitionsLatestResponse represents the response from the Latest Mergers & Acquisitions API
type MergersAcquisitionsLatestResponse struct {
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

// MergersAcquisitionsLatest retrieves real-time data on the latest mergers and acquisitions
func (c *Client) MergersAcquisitionsLatest(params MergersAcquisitionsLatestParams) ([]MergersAcquisitionsLatestResponse, error) {
	if params.Page == nil {
		return nil, fmt.Errorf("page parameter is required")
	}

	if params.Limit == nil {
		return nil, fmt.Errorf("limit parameter is required")
	}

	if *params.Limit > 1000 {
		return nil, fmt.Errorf("limit cannot exceed 1000")
	}

	urlParams := map[string]string{
		"page":  fmt.Sprintf("%d", *params.Page),
		"limit": fmt.Sprintf("%d", *params.Limit),
	}

	var result []MergersAcquisitionsLatestResponse
	err := c.doRequest(c.BaseURL+"/mergers-acquisitions-latest", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
