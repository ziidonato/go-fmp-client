package go_fmp

import (
	"fmt"
	"time"
)

// SharesFloatAllParams represents the parameters for the Shares Float All API
type SharesFloatAllParams struct {
	Limit *int `json:"limit"` // Optional: Number of records to return
	Page  *int `json:"page"`  // Optional: Page number (default: 0)
}

// SharesFloatAllResponse represents the response from the Shares Float All API
type SharesFloatAllResponse struct {
	Symbol            string    `json:"symbol"`
	Date              time.Time `json:"date"`
	FreeFloat         float64   `json:"freeFloat"`
	FloatShares       float64   `json:"floatShares"`
	OutstandingShares float64   `json:"outstandingShares"`
	Source            string    `json:"source"`
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
