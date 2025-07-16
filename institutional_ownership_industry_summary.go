package go_fmp

import (
	"encoding/json"
	"fmt"
)

// InstitutionalOwnershipIndustrySummaryParams represents the parameters for the Institutional Ownership Industry Summary API
type InstitutionalOwnershipIndustrySummaryParams struct {
	Year    string `json:"year"`    // Required: Year (e.g., "2023")
	Quarter string `json:"quarter"` // Required: Quarter (e.g., "3")

// InstitutionalOwnershipIndustrySummaryResponse represents the response from the Institutional Ownership Industry Summary API
type InstitutionalOwnershipIndustrySummaryResponse struct {
	IndustryTitle string `json:"industryTitle"`
	IndustryValue int64  `json:"industryValue"`
	Date          string `json:"date"`

// InstitutionalOwnershipIndustrySummary retrieves industry performance summary
func (c *Client) InstitutionalOwnershipIndustrySummary(params InstitutionalOwnershipIndustrySummaryParams) ([]InstitutionalOwnershipIndustrySummaryResponse, error) {
	if params.Year == "" {
		return nil, fmt.Errorf("year parameter is required")
	}

	if params.Quarter == "" {
		return nil, fmt.Errorf("quarter parameter is required")
	}

	urlParams := map[string]string{
		"year":    params.Year,
		"quarter": params.Quarter,
	}

	return doRequest[[]InstitutionalOwnershipIndustrySummaryResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
