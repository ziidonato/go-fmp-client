package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HistoricalDowJonesResponse represents the response from the historical Dow Jones API
type HistoricalDowJonesResponse struct {
	DateAdded       string `json:"dateAdded"`
	AddedSecurity   string `json:"addedSecurity"`
	RemovedTicker   string `json:"removedTicker"`
	RemovedSecurity string `json:"removedSecurity"`
	Date            string `json:"date"`
	Symbol          string `json:"symbol"`
	Reason          string `json:"reason"`
}

// GetHistoricalDowJones retrieves historical data for the Dow Jones Industrial Average
func (c *Client) GetHistoricalDowJones() ([]HistoricalDowJonesResponse, error) {
	url := "https://financialmodelingprep.com/stable/historical-dowjones-constituent"

	resp, err := c.get(url, map[string]string{})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var result []HistoricalDowJonesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
