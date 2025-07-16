package go_fmp

import (
	"encoding/json"
	"fmt"
)

// ScoresBulkResponse represents the response from the Scores Bulk API
type ScoresBulkResponse struct {
	Symbol           string `json:"symbol"`
	ReportedCurrency string `json:"reportedCurrency"`
	AltmanZScore     string `json:"altmanZScore"`
	PiotroskiScore   string `json:"piotroskiScore"`
	WorkingCapital   string `json:"workingCapital"`
	TotalAssets      string `json:"totalAssets"`
	RetainedEarnings string `json:"retainedEarnings"`
	EBIT             string `json:"ebit"`
	MarketCap        string `json:"marketCap"`
	TotalLiabilities string `json:"totalLiabilities"`
	Revenue          string `json:"revenue"`

// ScoresBulk retrieves key financial scores and metrics for multiple symbols
func (c *Client) ScoresBulk() ([]ScoresBulkResponse, error) {
	return doRequest[[]ScoresBulkResponse](c, "https://financialmodelingprep.com/stable/scores-bulk", nil)
