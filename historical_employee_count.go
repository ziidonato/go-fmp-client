package go_fmp

import (
	"fmt"
	"time"
)

// HistoricalEmployeeCountParams represents the parameters for the Historical Employee Count API
type HistoricalEmployeeCountParams struct {
	Symbol string  `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Limit  *int    `json:"limit"`  // Optional: Number of records to return
}

// HistoricalEmployeeCountResponse represents the response from the Historical Employee Count API
type HistoricalEmployeeCountResponse struct {
	Symbol               string    `json:"symbol"`
	CIK                  string    `json:"cik"`
	AcceptanceTime       string    `json:"acceptanceTime"`
	PeriodOfReport       string    `json:"periodOfReport"`
	CompanyName          string    `json:"companyName"`
	FormType             FormType  `json:"formType"`
	FilingDate           time.Time `json:"filingDate"`
	EmployeeCount        int       `json:"employeeCount"`
	Source               string    `json:"source"`
}

// HistoricalEmployeeCount retrieves historical employee count data for a company based on specific reporting periods
func (c *Client) HistoricalEmployeeCount(params HistoricalEmployeeCountParams) ([]HistoricalEmployeeCountResponse, error) {
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

	var result []HistoricalEmployeeCountResponse
	err := c.doRequest(c.BaseURL+"/historical-employee-count", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
