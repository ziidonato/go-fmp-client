package go_fmp

import (
	"fmt"
	"time"
)

// InstitutionalOwnershipIndustrySummaryParams represents the parameters for the Institutional Ownership Industry Summary API
type InstitutionalOwnershipIndustrySummaryParams struct {
	Industry string  `json:"industry"` // Required: Industry name (e.g., "Software")
	Page     *int    `json:"page"`     // Optional: Page number (default: 0)
}

// InstitutionalOwnershipIndustrySummaryResponse represents the response from the Institutional Ownership Industry Summary API
type InstitutionalOwnershipIndustrySummaryResponse struct {
	Industry string    `json:"industry"`
	Date     time.Time `json:"date"`
	Holdings int       `json:"holdings"`
	Holders  int       `json:"holders"`
	Shares   float64   `json:"shares"`
	Value    float64   `json:"value"`
}

// InstitutionalOwnershipIndustrySummary retrieves industry performance summary
func (c *Client) InstitutionalOwnershipIndustrySummary(params InstitutionalOwnershipIndustrySummaryParams) ([]InstitutionalOwnershipIndustrySummaryResponse, error) {
	if params.Industry == "" {
		return nil, fmt.Errorf("industry parameter is required")
	}

	urlParams := map[string]string{
		"industry": params.Industry,
	}

	if params.Page != nil {
		urlParams["page"] = fmt.Sprintf("%d", *params.Page)
	}

	var result []InstitutionalOwnershipIndustrySummaryResponse
	err := c.doRequest(c.BaseURL+"/institutional-ownership/industry-summary", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
