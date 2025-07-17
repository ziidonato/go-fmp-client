package go_fmp

import (
	"fmt"
)

// IndustryPESnapshotResponse represents the response from the Industry PE Snapshot API
type IndustryPESnapshotResponse struct {
	Date     string  `json:"date"`
	Industry string  `json:"industry"`
	Exchange string  `json:"exchange"`
	PE       float64 `json:"pe"`
}

// GetIndustryPESnapshot retrieves the price-to-earnings (P/E) ratios for various industries
func (c *Client) GetIndustryPESnapshot(date, exchange, industry string) ([]IndustryPESnapshotResponse, error) {
	if date == "" {
		return nil, fmt.Errorf("date parameter is required")
	}

	url := fmt.Sprintf("%s/industries-pe/%s", c.BaseURL, date)
	params := map[string]string{}

	if exchange != "" {
		params["exchange"] = exchange
	}

	if industry != "" {
		params["industry"] = industry
	}

	var result []IndustryPESnapshotResponse
	err := c.doRequest(url, params, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
