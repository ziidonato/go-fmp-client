package go_fmp

import (
	"encoding/json"
)

// AvailableIndustriesResponse represents the response from the Available Industries API
type AvailableIndustriesResponse struct {
	Industry string `json:"industry"`
}

// AvailableIndustries retrieves a comprehensive list of industries where stock symbols are available
func (c *Client) AvailableIndustries() ([]AvailableIndustriesResponse, error) {
	resp, err := c.doRequest("https://financialmodelingprep.com/stable/available-industries", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []AvailableIndustriesResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
