package go_fmp

import (
	"fmt"
)

// SharesFloatAllParams represents the parameters for the All Shares Float API
type SharesFloatAllParams struct {
	Limit *int `json:"limit"` // Required: Number of results (Maximum 5000 records per request)
	Page  *int `json:"page"`  // Required: Page number (e.g., 0)
}

// SharesFloatAllResponse represents the response from the All Shares Float API
type SharesFloatAllResponse struct {
	Symbol            string  `json:"symbol"`
	Date              string  `json:"date"`
	FreeFloat         float64 `json:"freeFloat"`
	FloatShares       int64   `json:"floatShares"`
	OutstandingShares int64   `json:"outstandingShares"`
}

// SharesFloatAll retrieves comprehensive shares float data for all available companies
func (c *Client) SharesFloatAll(params SharesFloatAllParams) ([]SharesFloatAllResponse, error) {
	if params.Limit == nil {
		return nil, fmt.Errorf("limit parameter is required")
	}

	if params.Page == nil {
		return nil, fmt.Errorf("page parameter is required")
	}

	if *params.Limit > 5000 {
		return nil, fmt.Errorf("limit cannot exceed 5000")
	}

	urlParams := map[string]string{
		"limit": fmt.Sprintf("%d", *params.Limit),
		"page":  fmt.Sprintf("%d", *params.Page),
	}

	var result []SharesFloatAllResponse
	err := c.doRequest(c.BaseURL+"/shares-float-all", urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, err
}
