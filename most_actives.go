package go_fmp

import (
	"encoding/json"
	"fmt"
)

// MostActivesResponse represents the response from the top traded stocks API
type MostActivesResponse struct {
	Symbol            string  `json:"symbol"`
	Price             float64 `json:"price"`
	Name              string  `json:"name"`
	Change            float64 `json:"change"`
	ChangesPercentage float64 `json:"changesPercentage"`
	Exchange          string  `json:"exchange"`
}

// GetMostActives retrieves the most actively traded stocks
func (c *Client) GetMostActives() ([]MostActivesResponse, error) {
	url := "https://financialmodelingprep.com/stable/most-actives"

	resp, err := c.doRequest(url, map[string]string{})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()


	var result []MostActivesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
