package go_fmp

import (
	"time"
	"fmt"
)

// IndustryPESnapshotResponse represents the response from the industry PE snapshot API
type IndustryPESnapshotResponse struct {
	Date time.Time `json:"date"`
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

	url := c.BaseURL + "/industry-pe-snapshot"

	var result []IndustryPESnapshotResponse
	err := c.doRequest(url, params, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
