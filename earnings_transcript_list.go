package go_fmp

// EarningsTranscriptListResponse represents the response from the Earnings Transcript List API
type EarningsTranscriptListResponse struct {
	Symbol string `json:"symbol"`
	Date   string `json:"date"`
}

// EarningsTranscriptList retrieves earnings transcripts for all transcripts from a time period
func (c *Client) EarningsTranscriptList() ([]EarningsTranscriptListResponse, error) {
	url := "https://financialmodelingprep.com/stable/earnings-transcript-list"
	var result []EarningsTranscriptListResponse
	if err := c.doRequest(url, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
