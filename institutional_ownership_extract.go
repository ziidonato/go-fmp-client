package go_fmp

import (
	"fmt"
	"time"
)

// InstitutionalOwnershipExtractParams represents the parameters for the Institutional Ownership Extract API
type InstitutionalOwnershipExtractParams struct {
	CIK     string `json:"cik"`     // Required: CIK number (e.g., "0001388838")
	Year    string `json:"year"`    // Required: Year (e.g., "2023")
	Quarter string `json:"quarter"` // Required: Quarter (e.g., "3")
}

// InstitutionalOwnershipExtractResponse represents the response from the Institutional Ownership Extract API
type InstitutionalOwnershipExtractResponse struct {
	CIK           string    `json:"cik"`
	Date          time.Time `json:"date"`
	FilingDate    time.Time `json:"filingDate"`
	AcceptedDate  time.Time `json:"acceptedDate"`
	InvestorName  string    `json:"investorName"`
	Symbol        string    `json:"symbol"`
	SecurityName  string    `json:"securityName"`
	TypeOfSecurity string   `json:"typeOfSecurity"`
	SecurityCusip string    `json:"securityCusip"`
	SharesType    SharesType `json:"sharesType"`
	PutCallShare  string    `json:"putCallShare"`
	InvestmentDiscretion string `json:"investmentDiscretion"`
	OtherManager  string    `json:"otherManager"`
	Sole          int64     `json:"sole"`
	Shared        int64     `json:"shared"`
	None          int64     `json:"none"`
}

// InstitutionalOwnershipExtract retrieves detailed data from SEC filings for institutional ownership
func (c *Client) InstitutionalOwnershipExtract(params InstitutionalOwnershipExtractParams) ([]InstitutionalOwnershipExtractResponse, error) {
	if params.CIK == "" {
		return nil, fmt.Errorf("cik parameter is required")
	}

	if params.Year == "" {
		return nil, fmt.Errorf("year parameter is required")
	}

	if params.Quarter == "" {
		return nil, fmt.Errorf("quarter parameter is required")
	}

	urlParams := map[string]string{
		"cik":     params.CIK,
		"year":    params.Year,
		"quarter": params.Quarter,
	}

	var result []InstitutionalOwnershipExtractResponse
	err := c.doRequest(c.BaseURL+"/institutional-ownership/extract", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
