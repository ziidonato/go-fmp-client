package go_fmp

import (
	"fmt"
	"time"
)

// GradesParams represents the parameters for the Stock Grades API
type GradesParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// GradesResponse represents the response from the Stock Grades API
type GradesResponse struct {
	Symbol         string      `json:"symbol"`
	Date           time.Time   `json:"date"`
	GradingCompany string      `json:"gradingCompany"`
	PreviousGrade  string      `json:"previousGrade"`
	NewGrade       string      `json:"newGrade"`
	Action         GradeAction `json:"action"`
}

// Grades retrieves the latest stock grades from top analysts and financial institutions
func (c *Client) Grades(params GradesParams) ([]GradesResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	var result []GradesResponse
	err := c.doRequest(c.BaseURL+"/grades", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
