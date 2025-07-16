package go_fmp

import (
	"encoding/json"
)

// COTListResponse represents the response from the COT Report List API
type COTListResponse struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

// COTList retrieves a comprehensive list of available Commitment of Traders (COT) reports
func (c *Client) COTList() ([]COTListResponse, error) {
	resp, err := c.get("https://financialmodelingprep.com/stable/commitment-of-traders-list", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []COTListResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
