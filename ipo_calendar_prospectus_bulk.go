package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	// Using direct HTTP request since it returns a file
	url := fmt.Sprintf("%s/ipo-calendar-prospectus-bulk/%s", c.BaseURL, params.Year)
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	if c.APIKey != "" {
		req.URL.Query().Add("apikey", c.APIKey)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code: %d", resp.StatusCode)
	}

	var result []IPOCalendarProspectusBulkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}