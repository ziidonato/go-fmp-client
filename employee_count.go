package go_fmp

import (
	"fmt"
)

// EmployeeCountParams represents the parameters for the Company Employee Count API
type EmployeeCountParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Limit  *int   `json:"limit"`  // Optional: Number of results (Maximum 10000 records per request)
}

// EmployeeCountResponse represents the response from the Company Employee Count API
type EmployeeCountResponse struct {
	Symbol         string `json:"symbol"`
	CIK            string `json:"cik"`
	AcceptanceTime string `json:"acceptanceTime"`
	PeriodOfReport string `json:"periodOfReport"`
	CompanyName    string `json:"companyName"`
	FormType       string `json:"formType"`
	FilingDate     string `json:"filingDate"`
	EmployeeCount  int    `json:"employeeCount"`
	Source         string `json:"source"`
}

// EmployeeCount retrieves detailed workforce information for companies
func (c *Client) EmployeeCount(params EmployeeCountParams) ([]EmployeeCountResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	if params.Limit != nil {
		if *params.Limit > 10000 {
			return nil, fmt.Errorf("limit cannot exceed 10000")
		}
		urlParams["limit"] = fmt.Sprintf("%d", *params.Limit)
	}

	var result []EmployeeCountResponse
	err := c.doRequest("https://financialmodelingprep.com/stable/employee-count", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
