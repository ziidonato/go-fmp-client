package go_fmp

import (
	"time"
)

// NasdaqConstituentResponse represents the response from the Nasdaq Constituent API
type NasdaqConstituentResponse struct {
	Symbol         string    `json:"symbol"`
	Name           string    `json:"name"`
	Sector         string    `json:"sector"`
	SubSector      string    `json:"subSector"`
	HeadQuarter    string    `json:"headQuarter"`
	DateFirstAdded time.Time `json:"dateFirstAdded"`
	CIK            string    `json:"cik"`
	Founded        string    `json:"founded"`
}

// GetNasdaqConstituent retrieves detailed data on the Nasdaq index
func (c *Client) GetNasdaqConstituent() ([]NasdaqConstituentResponse, error) {
	var result []NasdaqConstituentResponse
	err := c.doRequest(c.BaseURL+"/nasdaq_constituent", nil, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
