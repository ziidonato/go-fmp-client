package go_fmp

import (
	"encoding/json"
)

// AvailableSectorsResponse represents the response from the Available Sectors API
type AvailableSectorsResponse struct {
	Sector string `json:"sector"`

// AvailableSectors retrieves a complete list of industry sectors
func (c *Client) AvailableSectors() ([]AvailableSectorsResponse, error) {
	return doRequest[[]AvailableSectorsResponse](c, "https://financialmodelingprep.com/stable/available-sectors", nil)
