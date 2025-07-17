package go_fmp

import (
	"fmt"
	"time"
)

// EarningsTranscriptListResponse represents the response from the Earnings Transcript List API
type EarningsTranscriptListResponse struct {
	Symbol string    `json:"symbol"`
	Date   time.Time `json:"date"`
	Quarter int      `json:"quarter"`
	Year    int      `json:"year"`
}

// EarningsTranscriptList retrieves a list of available earnings transcripts
func (c *Client) EarningsTranscriptList(year, quarter string) ([]EarningsTranscriptListResponse, error) {
	if year == "" {
		return nil, fmt.Errorf("year parameter is required")
	}

	if quarter == "" {
		return nil, fmt.Errorf("quarter parameter is required")
	}

	urlParams := map[string]string{
		"year":    year,
		"quarter": quarter,
	}

	var result []EarningsTranscriptListResponse
	err := c.doRequest(c.BaseURL+"/earning_call_transcript", urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
