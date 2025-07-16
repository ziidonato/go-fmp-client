package go_fmp

import (
	"encoding/json"
)

// EarningsTranscriptListResponse represents the response from the Earnings Transcript List API
type EarningsTranscriptListResponse struct {
	Symbol          string `json:"symbol"`
	CompanyName     string `json:"companyName"`
	NoOfTranscripts string `json:"noOfTranscripts"`
}

// EarningsTranscriptList retrieves a list of companies with earnings transcripts and the number of transcripts available
func (c *Client) EarningsTranscriptList() ([]EarningsTranscriptListResponse, error) {
	resp, err := c.get("https://financialmodelingprep.com/stable/earnings-transcript-list", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []EarningsTranscriptListResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
