package go_fmp

import (
	"time"
	"fmt"
)

// DowJonesConstituentResponse represents the response from the Dow Jones API
type DowJonesConstituentResponse struct {
	Symbol         string `json:"symbol"`
	Name           string `json:"name"`
	Sector         string `json:"sector"`
	SubSector      string `json:"subSector"`
	HeadQuarter    string `json:"headQuarter"`
	DateFirstAdded time.Time `json:"dateFirstAdded"`
	CIK            string `json:"cik"`
	Founded        string `json:"founded"`
}

// GetDowJonesConstituent retrieves data on the Dow Jones Industrial Average
func (c *Client) GetDowJonesConstituent() ([]DowJonesConstituentResponse, error) {
	url := c.BaseURL + "/dowjones-constituent"

	var result []DowJonesConstituentResponse
	err := c.doRequest(url, map[string]string{}, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
