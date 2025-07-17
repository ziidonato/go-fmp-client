package go_fmp

import (
	"fmt"
	"time"
)

// ExecutiveCompensationParams represents the parameters for the Executive Compensation API
type ExecutiveCompensationParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// ExecutiveCompensationResponse represents the response from the Executive Compensation API
type ExecutiveCompensationResponse struct {
	Symbol                string    `json:"symbol"`
	CIK                   string    `json:"cik"`
	CompanyName           string    `json:"companyName"`
	FilingDate            time.Time `json:"filingDate"`
	AcceptedDate          time.Time `json:"acceptedDate"`
	Name                  string    `json:"name"`
	Title                 string    `json:"title"`
	Year                  int       `json:"year"`
	Salary                float64   `json:"salary"`
	Bonus                 float64   `json:"bonus"`
	StockAward            float64   `json:"stockAward"`
	IncentivePlanCompensation float64 `json:"incentivePlanCompensation"`
	AllOtherCompensation  float64   `json:"allOtherCompensation"`
	Total                 float64   `json:"total"`
	NameAndPosition       string    `json:"nameAndPosition"`
	URL                   string    `json:"url"`
}

// ExecutiveCompensation retrieves comprehensive compensation data for company executives
func (c *Client) ExecutiveCompensation(params ExecutiveCompensationParams) ([]ExecutiveCompensationResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	var result []ExecutiveCompensationResponse
	err := c.doRequest(c.BaseURL+"/governance-executive-compensation", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
