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
	return doRequest[[]AvailableCountriesResponse](c, "https://financialmodelingprep.com/stable/available-countries", nil)
}
