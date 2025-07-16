package go_fmp

import (
	"encoding/json"
	"fmt"
)

// SP500ConstituentResponse represents the response from the S&P 500 index API
type SP500ConstituentResponse struct {
	Symbol         string `json:"symbol"`
	Name           string `json:"name"`
	Sector         string `json:"sector"`
	SubSector      string `json:"subSector"`
	HeadQuarter    string `json:"headQuarter"`
	DateFirstAdded string `json:"dateFirstAdded"`
	CIK            string `json:"cik"`
	Founded        string `json:"founded"`

// GetSP500Constituent retrieves detailed data on the S&P 500 index
func (c *Client) GetSP500Constituent() ([]SP500ConstituentResponse, error) {
	url := "https://financialmodelingprep.com/stable/sp500-constituent"

	return doRequest[[]SP500ConstituentResponse](c, url, map[string]string{}
