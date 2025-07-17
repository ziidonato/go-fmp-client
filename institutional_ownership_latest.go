package go_fmp

import (
	"fmt"
	"time"
)

// InstitutionalOwnershipLatestParams represents the parameters for the Institutional Ownership Latest API
type InstitutionalOwnershipLatestParams struct {
	Symbol string  `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Page   *int    `json:"page"`   // Optional: Page number (default: 0)
}

// InstitutionalOwnershipLatestResponse represents the response from the Institutional Ownership Latest API
type InstitutionalOwnershipLatestResponse struct {
	Symbol       string    `json:"symbol"`
	Date         time.Time `json:"date"`
	FilingDate   time.Time `json:"filingDate"`
	AcceptedDate time.Time `json:"acceptedDate"`
	FormType     FormType  `json:"formType"`
	CIK          string    `json:"cik"`
	CUSIP        string    `json:"cusip"`
	InvestorName string    `json:"investorName"`
	Shares       int64     `json:"shares"`
	Value        float64   `json:"value"`
	Weight       float64   `json:"weight"`
}

// InstitutionalOwnershipLatest retrieves the latest institutional ownership filings
func (c *Client) InstitutionalOwnershipLatest(params InstitutionalOwnershipLatestParams) ([]InstitutionalOwnershipLatestResponse, error) {
	urlParams := map[string]string{}

	if params.Page != nil {
		urlParams["page"] = fmt.Sprintf("%d", *params.Page)
	}

	var result []InstitutionalOwnershipLatestResponse
	err := c.doRequest(c.BaseURL+"/institutional-ownership/latest", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
