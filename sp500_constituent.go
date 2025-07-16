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
}

// GetSP500Constituent retrieves detailed data on the S&P 500 index
func (c *Client) GetSP500Constituent() ([]SP500ConstituentResponse, error) {
	url := "https://financialmodelingprep.com/stable/sp500-constituent"

	resp, err := c.doRequest(url, map[string]string{})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()


	var result []SP500ConstituentResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
