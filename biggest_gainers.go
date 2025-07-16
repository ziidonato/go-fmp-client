package go_fmp

import (
	"encoding/json"
	"fmt"
)

// BiggestGainersResponse represents the response from the biggest stock gainers API
type BiggestGainersResponse struct {
	Symbol            string  `json:"symbol"`
	Price             float64 `json:"price"`
	Name              string  `json:"name"`
	Change            float64 `json:"change"`
	ChangesPercentage float64 `json:"changesPercentage"`
	Exchange          string  `json:"exchange"`

// GetBiggestGainers retrieves the stocks with the largest price increases
func (c *Client) BiggestGainers() ([]BiggestGainersResponse, error) {
	url := "https://financialmodelingprep.com/stable/biggest-gainers"

	return doRequest[[]BiggestGainersResponse](c, url, map[string]string{}
