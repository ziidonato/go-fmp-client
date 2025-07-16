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

// GetPriceTargetSummaryBulk retrieves comprehensive overview of price targets for all listed symbols
func (c *Client) GetPriceTargetSummaryBulk() ([]PriceTargetSummaryBulkResponse, error) {
	return doRequest[[]PriceTargetSummaryBulkResponse](c, "https://financialmodelingprep.com/stable/price-target-summary-bulk", nil)
