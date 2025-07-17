package go_fmp

import (
	"fmt"
	"time"
)

// IPOCalendarProspectusBulkParams represents the parameters for the IPO Calendar Prospectus Bulk API
type IPOCalendarProspectusBulkParams struct {
	Year string `json:"year"` // Required: Year (e.g., "2024")
}

// IPOCalendarProspectusBulkResponse represents the response from the IPO Calendar Prospectus Bulk API
type IPOCalendarProspectusBulkResponse struct {
	Symbol              string    `json:"symbol"`
	CIK                 string    `json:"cik"`
	CompanyName         string    `json:"companyName"`
	ProspectusDate      time.Time `json:"prospectusDate"`
	FilingDate          time.Time `json:"filingDate"`
	AcceptedDate        time.Time `json:"acceptedDate"`
	FormType            FormType  `json:"formType"`
	SECFilingLink       string    `json:"secFilingLink"`
	FinalProspectusLink string    `json:"finalProspectusLink"`
}

// IPOCalendarProspectusBulk retrieves all IPO prospectus data for a specific year
func (c *Client) IPOCalendarProspectusBulk(params IPOCalendarProspectusBulkParams) ([]IPOCalendarProspectusBulkResponse, error) {
	if params.Year == "" {
		return nil, fmt.Errorf("year parameter is required")
	}

	url := fmt.Sprintf("ipo-calendar-prospectus-bulk/%s", params.Year)
	
	var result []IPOCalendarProspectusBulkResponse
	err := c.doRequest(c.BaseURL+"/"+url, nil, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}