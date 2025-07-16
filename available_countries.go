package go_fmp

import (
	"encoding/json"
)

// AvailableCountriesResponse represents the response from the Available Countries API
type AvailableCountriesResponse struct {
	Country string `json:"country"`
}

// AvailableCountries retrieves a comprehensive list of countries where stock symbols are available
func (c *Client) AvailableCountries() ([]AvailableCountriesResponse, error) {
	resp, err := c.get("https://financialmodelingprep.com/stable/available-countries", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []AvailableCountriesResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
