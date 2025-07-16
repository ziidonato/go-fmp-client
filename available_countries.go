package go_fmp

// AvailableCountriesResponse represents the response from the Available Countries API
type AvailableCountriesResponse struct {
	Country string `json:"country"`
}

// AvailableCountries retrieves a comprehensive list of countries where stock symbols are available
func (c *Client) AvailableCountries() ([]AvailableCountriesResponse, error) {
	url := c.BaseURL + "/available-countries"
	var result []AvailableCountriesResponse
	if err := c.doRequest(url, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
