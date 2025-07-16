package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// IndustryPESnapshotResponse represents the response from the industry PE snapshot API
type IndustryPESnapshotResponse struct {
	Date     string  `json:"date"`
	Industry string  `json:"industry"`
	Exchange string  `json:"exchange"`
	PE       float64 `json:"pe"`
}

// GetIndustryPESnapshot retrieves price-to-earnings (P/E) ratios for different industries
func (c *Client) GetIndustryPESnapshot(date, exchange, industry string) ([]IndustryPESnapshotResponse, error) {
	if date == "" {
		return nil, fmt.Errorf("date is required")
	}

	params := map[string]string{
		"date": date,
	}

	if exchange != "" {
		params["exchange"] = exchange
	}
	if industry != "" {
		params["industry"] = industry
	}

	url := "https://financialmodelingprep.com/stable/industry-pe-snapshot"

	resp, err := c.get(url, params)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var result []IndustryPESnapshotResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
