package go_fmp

import (
	"encoding/json"
	"fmt"
)

// InstitutionalOwnershipLatestParams represents the parameters for the Institutional Ownership Latest API
type InstitutionalOwnershipLatestParams struct {
	Page  *int `json:"page"`  // Optional: Page number (default: 0)
	Limit *int `json:"limit"` // Optional: Number of results (default: 100)
}

// InstitutionalOwnershipLatestResponse represents the response from the Institutional Ownership Latest API
type InstitutionalOwnershipLatestResponse struct {
	CIK          string `json:"cik"`
	Name         string `json:"name"`
	Date         string `json:"date"`
	FilingDate   string `json:"filingDate"`
	AcceptedDate string `json:"acceptedDate"`
	FormType     string `json:"formType"`
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

	resp, err := c.get("https://financialmodelingprep.com/stable/institutional-ownership/latest", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []InstitutionalOwnershipLatestResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
