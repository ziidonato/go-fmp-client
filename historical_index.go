package go_fmp

import (
	"time"
)

// HistoricalIndexResponse represents the response from the historical index constituent change APIs
// (S&P 500, Nasdaq, Dow Jones)
type HistoricalIndexResponse struct {
	DateAdded time.Time `json:"dateAdded"`
	AddedSecurity   string `json:"addedSecurity"`
	RemovedTicker   string `json:"removedTicker"`
	RemovedSecurity string `json:"removedSecurity"`
	Date time.Time `json:"date"`
	Symbol          string `json:"symbol"`
	Reason          string `json:"reason"`
}

// HistoricalSP500 retrieves historical data for the S&P 500 index
func (c *Client) HistoricalSP500() ([]HistoricalIndexResponse, error) {
	url := c.BaseURL + "/historical-sp500-constituent"
	var result []HistoricalIndexResponse
	if err := c.doRequest(url, map[string]string{}, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// HistoricalNasdaq retrieves historical data for the Nasdaq index
func (c *Client) HistoricalNasdaq() ([]HistoricalIndexResponse, error) {
	url := c.BaseURL + "/historical-nasdaq-constituent"
	var result []HistoricalIndexResponse
	if err := c.doRequest(url, map[string]string{}, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// HistoricalDowJones retrieves historical data for the Dow Jones Industrial Average
func (c *Client) HistoricalDowJones() ([]HistoricalIndexResponse, error) {
	url := c.BaseURL + "/historical-dowjones-constituent"
	var result []HistoricalIndexResponse
	if err := c.doRequest(url, map[string]string{}, &result); err != nil {
		return nil, err
	}
	return result, nil
}
