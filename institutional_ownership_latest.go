package go_fmp

import (
	"fmt"
	"time"
)

// InstitutionalOwnershipLatestParams represents the parameters for the Institutional Ownership Latest API
type InstitutionalOwnershipLatestParams struct {
	Page  *int `json:"page"`  // Optional: Page number (default: 0)
	Limit *int `json:"limit"` // Optional: Number of results (default: 100)
}

// InstitutionalOwnershipLatestResponse represents the response from the Institutional Ownership Latest API
type InstitutionalOwnershipLatestResponse struct {
	CIK          string    `json:"cik"`
	Name         string    `json:"name"`
	Date         time.Time `json:"date"`
	FilingDate   time.Time `json:"filingDate"`
	AcceptedDate time.Time `json:"acceptedDate"`
	FormType     FormType  `json:"formType"`
	Link         string `json:"link"`
	FinalLink    string `json:"finalLink"`
}

// InstitutionalOwnershipLatest retrieves the latest institutional ownership filings
func (c *Client) InstitutionalOwnershipLatest(params InstitutionalOwnershipLatestParams) ([]InstitutionalOwnershipLatestResponse, error) {
	urlParams := map[string]string{}

	if params.Page != nil {
		urlParams["page"] = fmt.Sprintf("%d", *params.Page)
	}

	if params.Limit != nil {
		urlParams["limit"] = fmt.Sprintf("%d", *params.Limit)
	}

	var result []InstitutionalOwnershipLatestResponse
	err := c.doRequest(c.BaseURL+"/institutional-ownership/latest", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
