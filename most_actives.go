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

// GetMostActives retrieves the most actively traded stocks
func (c *Client) GetMostActives() ([]MostActivesResponse, error) {
	url := "https://financialmodelingprep.com/stable/most-actives"

	return doRequest[[]MostActivesResponse](c, url, map[string]string{}
