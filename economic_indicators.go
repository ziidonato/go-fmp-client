package go_fmp

import (
	"encoding/json"
	"fmt"
)

// EconomicIndicatorsParams represents the parameters for the Economic Indicators API
type EconomicIndicatorsParams struct {
	Name string `json:"name"` // Required: Economic indicator name (e.g., "GDP")
}

// EconomicIndicatorsResponse represents the response from the Economic Indicators API
type EconomicIndicatorsResponse struct {
	Date        string  `json:"date"`
	Name        string  `json:"name"`
	Value       float64 `json:"value"`
	Unit        string  `json:"unit"`
	Period      string  `json:"period"`
	Source      string  `json:"source"`
	Description string  `json:"description"`
}

// EconomicIndicators retrieves real-time and historical economic data for key indicators
func (c *Client) EconomicIndicators(params EconomicIndicatorsParams) ([]EconomicIndicatorsResponse, error) {
	if params.Name == "" {
		return nil, fmt.Errorf("name parameter is required")
	}

	urlParams := map[string]string{
		"name": params.Name,
	}

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/economic-indicators", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []EconomicIndicatorsResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
