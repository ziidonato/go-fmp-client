package go_fmp

import (
	"encoding/json"
	"fmt"
)

// DowJonesConstituentResponse represents the response from the Dow Jones API
type DowJonesConstituentResponse struct {
	Symbol         string `json:"symbol"`
	Name           string `json:"name"`
	Sector         string `json:"sector"`
	SubSector      string `json:"subSector"`
	HeadQuarter    string `json:"headQuarter"`
	DateFirstAdded string `json:"dateFirstAdded"`
	CIK            string `json:"cik"`
	Founded        string `json:"founded"`
}

// GetDowJonesConstituent retrieves data on the Dow Jones Industrial Average
func (c *Client) GetDowJonesConstituent() ([]DowJonesConstituentResponse, error) {
	url := "https://financialmodelingprep.com/stable/dowjones-constituent"

	return doRequest[[]DowJonesConstituentResponse](c, url, map[string]string{})
}
