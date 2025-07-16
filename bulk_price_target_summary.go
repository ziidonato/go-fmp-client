package go_fmp

import (
	"encoding/json"
	"fmt"
)

// PriceTargetSummaryBulkResponse represents the response from the Price Target Summary Bulk API
type PriceTargetSummaryBulkResponse struct {
	Symbol                    string `json:"symbol"`
	LastMonthCount            string `json:"lastMonthCount"`
	LastMonthAvgPriceTarget   string `json:"lastMonthAvgPriceTarget"`
	LastQuarterCount          string `json:"lastQuarterCount"`
	LastQuarterAvgPriceTarget string `json:"lastQuarterAvgPriceTarget"`
	LastYearCount             string `json:"lastYearCount"`
	LastYearAvgPriceTarget    string `json:"lastYearAvgPriceTarget"`
	AllTimeCount              string `json:"allTimeCount"`
	AllTimeAvgPriceTarget     string `json:"allTimeAvgPriceTarget"`
	Publishers                string `json:"publishers"`
}

// GetPriceTargetSummaryBulk retrieves comprehensive overview of price targets for all listed symbols
func (c *Client) GetPriceTargetSummaryBulk() ([]PriceTargetSummaryBulkResponse, error) {
	resp, err := c.doRequest("https://financialmodelingprep.com/stable/price-target-summary-bulk", nil)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()



	// Parse the response
	var result []PriceTargetSummaryBulkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
