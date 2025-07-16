package go_fmp

import (
	"encoding/json"
	"fmt"
)

// NasdaqConstituentResponse represents the response from the Nasdaq index API
type NasdaqConstituentResponse struct {
	Symbol         string  `json:"symbol"`
	Name           string  `json:"name"`
	Sector         string  `json:"sector"`
	SubSector      string  `json:"subSector"`
	HeadQuarter    string  `json:"headQuarter"`
	DateFirstAdded *string `json:"dateFirstAdded"`
	CIK            string  `json:"cik"`
	Founded        string  `json:"founded"`

// GetNasdaqConstituent retrieves comprehensive data for the Nasdaq index
func (c *Client) GetNasdaqConstituent() ([]NasdaqConstituentResponse, error) {
	url := "https://financialmodelingprep.com/stable/nasdaq-constituent"

	return doRequest[[]NasdaqConstituentResponse](c, url, map[string]string{}
