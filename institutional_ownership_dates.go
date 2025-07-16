package go_fmp

import (
	"encoding/json"
	"fmt"
)

// InstitutionalOwnershipDatesParams represents the parameters for the Institutional Ownership Dates API
type InstitutionalOwnershipDatesParams struct {
	CIK string `json:"cik"` // Required: CIK number (e.g., "0001067983")
}

// InstitutionalOwnershipDatesResponse represents the response from the Institutional Ownership Dates API
type InstitutionalOwnershipDatesResponse struct {
	Date    string `json:"date"`
	Year    int    `json:"year"`
	Quarter int    `json:"quarter"`
}

// InstitutionalOwnershipDates retrieves Form 13F filing dates for institutional investors
func (c *Client) InstitutionalOwnershipDates(params InstitutionalOwnershipDatesParams) ([]InstitutionalOwnershipDatesResponse, error) {
	if params.CIK == "" {
		return nil, fmt.Errorf("cik parameter is required")
	}

	urlParams := map[string]string{
		"cik": params.CIK,
	}

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/institutional-ownership/dates", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []InstitutionalOwnershipDatesResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
