package go_fmp

// COTListResponse represents the response from the COT Report List API
type COTListResponse struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

// COTList retrieves a comprehensive list of available Commitment of Traders (COT) reports
func (c *Client) COTList() ([]COTListResponse, error) {
	var result []COTListResponse
	err := c.doRequest("https://financialmodelingprep.com/stable/commitment-of-traders-list", nil, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
