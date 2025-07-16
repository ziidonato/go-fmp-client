package go_fmp

import (
	"encoding/json"
)

// AvailableExchangesResponse represents the response from the Available Exchanges API
type AvailableExchangesResponse struct {
	Exchange     string `json:"exchange"`
	Name         string `json:"name"`
	CountryName  string `json:"countryName"`
	CountryCode  string `json:"countryCode"`
	SymbolSuffix string `json:"symbolSuffix"`
	Delay        string `json:"delay"`
}

// AvailableExchanges retrieves a complete list of supported stock exchanges
func (c *Client) AvailableExchanges() ([]AvailableExchangesResponse, error) {
	resp, err := c.get("https://financialmodelingprep.com/stable/available-exchanges", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []AvailableExchangesResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
