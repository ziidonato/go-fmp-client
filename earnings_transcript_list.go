package go_fmp

import (
	"time"
)

// EarningsTranscriptListResponse represents the response from the Earnings Transcript List API
type EarningsTranscriptListResponse struct {
	Symbol string `json:"symbol"`
	Date time.Time `json:"date"`
}

// EarningsTranscriptList retrieves earnings transcripts for all transcripts from a time period
func (c *Client) EarningsTranscriptList() ([]EarningsTranscriptListResponse, error) {
	url := c.BaseURL + "/earnings-transcript-list"
	var result []EarningsTranscriptListResponse
	if err := c.doRequest(url, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
