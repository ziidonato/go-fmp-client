package go_fmp

import (
	"time"
)

// SP500ConstituentResponse represents the response from the S&P 500 Constituent API
type SP500ConstituentResponse struct {
	Symbol         string    `json:"symbol"`
	DateFirstAdded time.Time `json:"dateFirstAdded"`
	CIK            string    `json:"cik"`
	Founded        string    `json:"founded"`
	Description    string    `json:"description"`
	Name           string    `json:"name"`
}

// GetSP500Constituent retrieves detailed data on the S&P 500 index
func (c *Client) GetSP500Constituent() ([]SP500ConstituentResponse, error) {
	var result []SP500ConstituentResponse
	err := c.doRequest(c.BaseURL+"/sp500_constituent", nil, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
