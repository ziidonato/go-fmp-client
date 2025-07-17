package go_fmp

// IndexListResponse represents the response from the index list API
type IndexListResponse struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Exchange Exchange `json:"exchange"`
	Currency Currency `json:"currency"`
}

// GetIndexList retrieves a comprehensive list of stock market indexes across global exchanges
func (c *Client) GetIndexList() ([]IndexListResponse, error) {
	url := c.BaseURL + "/index-list"

	var result []IndexListResponse
	err := c.doRequest(url, map[string]string{}, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
