package go_fmp

import (
	"encoding/json"
)

// AvailableIndustriesResponse represents the response from the Available Industries API
type AvailableIndustriesResponse struct {
	Industry string `json:"industry"`

// AvailableIndustries retrieves a comprehensive list of industries where stock symbols are available
func (c *Client) AvailableIndustries() ([]AvailableIndustriesResponse, error) {
	return doRequest[[]AvailableIndustriesResponse](c, "https://financialmodelingprep.com/stable/available-industries", nil)
