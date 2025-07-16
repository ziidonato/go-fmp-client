package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HistoricalSP500Response represents the response from the historical S&P 500 API
type HistoricalSP500Response struct {
	DateAdded       string `json:"dateAdded"`
	AddedSecurity   string `json:"addedSecurity"`
	RemovedTicker   string `json:"removedTicker"`
	RemovedSecurity string `json:"removedSecurity"`
	Date            string `json:"date"`
	Symbol          string `json:"symbol"`
	Reason          string `json:"reason"`
}

// GetHistoricalSP500 retrieves historical data for the S&P 500 index
func (c *Client) GetHistoricalSP500() ([]HistoricalSP500Response, error) {
	url := "https://financialmodelingprep.com/stable/historical-sp500-constituent"

	resp, err := c.get(url, map[string]string{})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var result []HistoricalSP500Response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
