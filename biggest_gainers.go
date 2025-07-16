package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// BiggestGainersResponse represents the response from the biggest stock gainers API
type BiggestGainersResponse struct {
	Symbol            string  `json:"symbol"`
	Price             float64 `json:"price"`
	Name              string  `json:"name"`
	Change            float64 `json:"change"`
	ChangesPercentage float64 `json:"changesPercentage"`
	Exchange          string  `json:"exchange"`
}

// GetBiggestGainers retrieves the stocks with the largest price increases
func (c *Client) BiggestGainers() ([]BiggestGainersResponse, error) {
	url := "https://financialmodelingprep.com/stable/biggest-gainers"

	resp, err := c.get(url, map[string]string{})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var result []BiggestGainersResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
