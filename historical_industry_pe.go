package go_fmp

import (
	"encoding/json"
	"fmt"
)

// HistoricalIndustryPEResponse represents the response from the historical industry PE API
type HistoricalIndustryPEResponse struct {
	Date     string  `json:"date"`
	Industry string  `json:"industry"`
	Exchange string  `json:"exchange"`
	PE       float64 `json:"pe"`

// GetHistoricalIndustryPE retrieves historical price-to-earnings (P/E) ratios by industry
func (c *Client) GetHistoricalIndustryPE(industry, from, to, exchange string) ([]HistoricalIndustryPEResponse, error) {
	if industry == "" {
		return nil, fmt.Errorf("industry is required")
	}

	params := map[string]string{
		"industry": industry,
	}

	if from != "" {
		params["from"] = from
	}
	if to != "" {
		params["to"] = to
	}
	if exchange != "" {
		params["exchange"] = exchange
	}

	url := "https://financialmodelingprep.com/stable/historical-industry-pe"

	return doRequest[[]HistoricalIndustryPEResponse](c, url, params)
