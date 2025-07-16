package go_fmp

// AvailableSectorsResponse represents the response from the Available Sectors API
type AvailableSectorsResponse struct {
	Sector string `json:"sector"`
}

// AvailableSectors retrieves a complete list of industry sectors
func (c *Client) AvailableSectors() ([]AvailableSectorsResponse, error) {
	url := c.BaseURL + "/available-sectors"
	var result []AvailableSectorsResponse
	if err := c.doRequest(url, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
