package go_fmp

// HistoricalIndexResponse represents the response from the historical index constituent change APIs
// (S&P 500, Nasdaq, Dow Jones)
type HistoricalIndexResponse struct {
	DateAdded       string `json:"dateAdded"`
	AddedSecurity   string `json:"addedSecurity"`
	RemovedTicker   string `json:"removedTicker"`
	RemovedSecurity string `json:"removedSecurity"`
	Date            string `json:"date"`
	Symbol          string `json:"symbol"`
	Reason          string `json:"reason"`
}

// HistoricalSP500 retrieves historical data for the S&P 500 index
func (c *Client) HistoricalSP500() ([]HistoricalIndexResponse, error) {
	url := "https://financialmodelingprep.com/stable/historical-sp500-constituent"
	var result []HistoricalIndexResponse
	if err := c.doRequest(url, map[string]string{}, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// HistoricalNasdaq retrieves historical data for the Nasdaq index
func (c *Client) HistoricalNasdaq() ([]HistoricalIndexResponse, error) {
	url := "https://financialmodelingprep.com/stable/historical-nasdaq-constituent"
	var result []HistoricalIndexResponse
	if err := c.doRequest(url, map[string]string{}, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// HistoricalDowJones retrieves historical data for the Dow Jones Industrial Average
func (c *Client) HistoricalDowJones() ([]HistoricalIndexResponse, error) {
	url := "https://financialmodelingprep.com/stable/historical-dowjones-constituent"
	var result []HistoricalIndexResponse
	if err := c.doRequest(url, map[string]string{}, &result); err != nil {
		return nil, err
	}
	return result, nil
}
