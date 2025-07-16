package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// BiggestLosersResponse represents the response from the biggest stock losers API
type BiggestLosersResponse struct {
	Symbol            string  `json:"symbol"`
	Price             float64 `json:"price"`
	Name              string  `json:"name"`
	Change            float64 `json:"change"`
	ChangesPercentage float64 `json:"changesPercentage"`
	Exchange          string  `json:"exchange"`
}

// GetBiggestLosers retrieves the stocks with the largest price drops
func (c *Client) BiggestLosers() ([]BiggestLosersResponse, error) {
	url := "https://financialmodelingprep.com/stable/biggest-losers"

	resp, err := c.get(url, map[string]string{})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var result []BiggestLosersResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
