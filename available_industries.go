package go_fmp

// AvailableIndustriesResponse represents the response from the Available Industries API
type AvailableIndustriesResponse struct {
	Industry string `json:"industry"`
}

// AvailableIndustries retrieves a comprehensive list of industries where stock symbols are available
func (c *Client) AvailableIndustries() ([]AvailableIndustriesResponse, error) {
	url := c.BaseURL + "/available-industries"
	var result []AvailableIndustriesResponse
	if err := c.doRequest(url, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
