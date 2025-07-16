package go_fmp

import (
	"encoding/json"
)

// AvailableSectorsResponse represents the response from the Available Sectors API
type AvailableSectorsResponse struct {
	Sector string `json:"sector"`
}

// AvailableSectors retrieves a complete list of industry sectors
func (c *Client) AvailableSectors() ([]AvailableSectorsResponse, error) {
	resp, err := c.doRequest("https://financialmodelingprep.com/stable/available-sectors", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []AvailableSectorsResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
