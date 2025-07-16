package go_fmp

import (
	"fmt"
)

// ExecutiveCompensationParams represents the parameters for the Executive Compensation API
type ExecutiveCompensationParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// ExecutiveCompensationResponse represents the response from the Executive Compensation API
type ExecutiveCompensationResponse struct {
	CIK                       string `json:"cik"`
	Symbol                    string `json:"symbol"`
	CompanyName               string `json:"companyName"`
	FilingDate                string `json:"filingDate"`
	AcceptedDate              string `json:"acceptedDate"`
	NameAndPosition           string `json:"nameAndPosition"`
	Year                      int    `json:"year"`
	Salary                    int    `json:"salary"`
	Bonus                     int    `json:"bonus"`
	StockAward                int    `json:"stockAward"`
	OptionAward               int    `json:"optionAward"`
	IncentivePlanCompensation int    `json:"incentivePlanCompensation"`
	AllOtherCompensation      int    `json:"allOtherCompensation"`
	Total                     int    `json:"total"`
	Link                      string `json:"link"`
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
	err := c.doRequest("https://financialmodelingprep.com/stable/governance-executive-compensation", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
