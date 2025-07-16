package go_fmp

import (
	"encoding/json"
	"fmt"
)

// DelistedCompaniesParams represents the parameters for the Delisted Companies API
type DelistedCompaniesParams struct {
	Page  *int `json:"page"`  // Required: Page number (e.g., 0)
	Limit *int `json:"limit"` // Required: Number of results (Maximum 100 records per request)
}

// DelistedCompaniesResponse represents the response from the Delisted Companies API
type DelistedCompaniesResponse struct {
	Symbol       string `json:"symbol"`
	CompanyName  string `json:"companyName"`
	Exchange     string `json:"exchange"`
	IPODate      string `json:"ipoDate"`
	DelistedDate string `json:"delistedDate"`
}

// DelistedCompanies retrieves a comprehensive list of companies that have been delisted from US exchanges
func (c *Client) DelistedCompanies(params DelistedCompaniesParams) ([]DelistedCompaniesResponse, error) {
	if params.Page == nil {
		return nil, fmt.Errorf("page parameter is required")
	}

	if params.Limit == nil {
		return nil, fmt.Errorf("limit parameter is required")
	}

	if *params.Limit > 100 {
		return nil, fmt.Errorf("limit cannot exceed 100")
	}

	urlParams := map[string]string{
		"page":  fmt.Sprintf("%d", *params.Page),
		"limit": fmt.Sprintf("%d", *params.Limit),
	}

	return doRequest[[]DelistedCompaniesResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
}
