package go_fmp

import (
	"encoding/json"
	"fmt"
)

// IndexListResponse represents the response from the index list API
type IndexListResponse struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Exchange string `json:"exchange"`
	Currency string `json:"currency"`
}

// GetIndexList retrieves a comprehensive list of stock market indexes across global exchanges
func (c *Client) GetIndexList() ([]IndexListResponse, error) {
	url := "https://financialmodelingprep.com/stable/index-list"

	resp, err := c.doRequest(url, map[string]string{})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()


	var result []IndexListResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
