package go_fmp

import "fmt"

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
	var result []AvailableExchangesResponse
	err := c.doRequest(c.BaseURL+"/available-exchanges", map[string]string{}, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
