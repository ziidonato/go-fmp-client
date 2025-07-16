package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
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
}

// GetNasdaqConstituent retrieves comprehensive data for the Nasdaq index
func (c *Client) GetNasdaqConstituent() ([]NasdaqConstituentResponse, error) {
	url := "https://financialmodelingprep.com/stable/nasdaq-constituent"

	resp, err := c.get(url, map[string]string{})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var result []NasdaqConstituentResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
