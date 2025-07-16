package go_fmp

import (
	"encoding/json"
	"fmt"
)

// ETFHolderBulkParams represents the parameters for the ETF Holder Bulk API
type ETFHolderBulkParams struct {
	Part string `json:"part"` // Required: part number (e.g., "1")
}

// ETFHolderBulkResponse represents the response from the ETF Holder Bulk API
type ETFHolderBulkResponse struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	SharesNumber     string `json:"sharesNumber"`
	Asset            string `json:"asset"`
	WeightPercentage string `json:"weightPercentage"`
	CUSIP            string `json:"cusip"`
	ISIN             string `json:"isin"`
	MarketValue      string `json:"marketValue"`
	LastUpdated      string `json:"lastUpdated"`
}

// ETFHolderBulk retrieves detailed information about assets and shares held by ETFs
func (c *Client) ETFHolderBulk(params ETFHolderBulkParams) ([]ETFHolderBulkResponse, error) {
	// Validate required parameters
	if params.Part == "" {
		return nil, fmt.Errorf("part parameter is required")
	}

	urlParams := map[string]string{
		"part": params.Part,
	}

	return doRequest[[]ETFHolderBulkResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
}
