package go_fmp

import (
	"encoding/json"
	"fmt"
)

// GradesParams represents the parameters for the Stock Grades API
type GradesParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// GradesResponse represents the response from the Stock Grades API
type GradesResponse struct {
	Symbol         string `json:"symbol"`
	Date           string `json:"date"`
	GradingCompany string `json:"gradingCompany"`
	PreviousGrade  string `json:"previousGrade"`
	NewGrade       string `json:"newGrade"`
	Action         string `json:"action"`
}

// Grades retrieves the latest stock grades from top analysts and financial institutions
func (c *Client) Grades(params GradesParams) ([]GradesResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/grades", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []GradesResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
