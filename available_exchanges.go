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

// AvailableExchanges retrieves a complete list of supported stock exchanges
func (c *Client) AvailableExchanges() ([]AvailableExchangesResponse, error) {
	return doRequest[[]AvailableExchangesResponse](c, "https://financialmodelingprep.com/stable/available-exchanges", nil)
