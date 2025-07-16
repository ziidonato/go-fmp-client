package go_fmp

import (
	"fmt"
)

// CIKListParams represents the parameters for the CIK List API
type CIKListParams struct {
	Limit *int `json:"limit"` // Required: Number of results (Maximum 10000 records per request)
}

// CIKListResponse represents the response from the CIK List API
type CIKListResponse struct {
	CIK         string `json:"cik"`
	CompanyName string `json:"companyName"`
}

// CIKList retrieves a comprehensive database of CIK numbers for SEC-registered entities
func (c *Client) CIKList(params CIKListParams) ([]CIKListResponse, error) {
	if params.Limit == nil {
		return nil, fmt.Errorf("limit parameter is required")
	}

	if *params.Limit > 10000 {
		return nil, fmt.Errorf("limit cannot exceed 10000")
	}

	urlParams := map[string]string{
		"limit": fmt.Sprintf("%d", *params.Limit),
	}

	var result []CIKListResponse
	err := c.doRequest("https://financialmodelingprep.com/stable/cik-list", urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
