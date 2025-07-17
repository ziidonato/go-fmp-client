package go_fmp

import (
	"fmt"
	"time"
)

// COTAnalysisParams represents the parameters for the COT Analysis By Dates API
type COTAnalysisParams struct {
	Symbol *string `json:"symbol"` // Optional: Symbol (e.g., "AAPL")
	From   *string `json:"from"`   // Required: Start date (e.g., "2024-01-01")
	To     *string `json:"to"`     // Required: End date (e.g., "2024-03-01")
}

// COTAnalysisResponse represents the response from the COT Analysis By Dates API
type COTAnalysisResponse struct {
	Symbol                       string  `json:"symbol"`
	Date time.Time `json:"date"`
	Name                         string  `json:"name"`
	Sector                       string  `json:"sector"`
	Exchange                     string  `json:"exchange"`
	CurrentLongMarketSituation   float64 `json:"currentLongMarketSituation"`
	CurrentShortMarketSituation  float64 `json:"currentShortMarketSituation"`
	MarketSituation MarketSituation `json:"marketSituation"`
	PreviousLongMarketSituation  float64 `json:"previousLongMarketSituation"`
	PreviousShortMarketSituation float64 `json:"previousShortMarketSituation"`
	PreviousMarketSituation MarketSituation `json:"previousMarketSituation"`
	NetPostion                   int     `json:"netPostion"`
	PreviousNetPosition          int     `json:"previousNetPosition"`
	ChangeInNetPosition          float64 `json:"changeInNetPosition"`
	MarketSentiment MarketSentiment `json:"marketSentiment"`
	ReversalTrend                bool    `json:"reversalTrend"`
}

// COTAnalysis retrieves in-depth insights into market sentiment with COT report analysis
func (c *Client) COTAnalysis(params COTAnalysisParams) ([]COTAnalysisResponse, error) {
	if params.From == nil {
		return nil, fmt.Errorf("from parameter is required")
	}

	if params.To == nil {
		return nil, fmt.Errorf("to parameter is required")
	}

	// Validate date range (max 90 days)
	fromDate, err := time.Parse("2006-01-02", *params.From)
	if err != nil {
		return nil, fmt.Errorf("invalid from date format, expected YYYY-MM-DD")
	}

	toDate, err := time.Parse("2006-01-02", *params.To)
	if err != nil {
		return nil, fmt.Errorf("invalid to date format, expected YYYY-MM-DD")
	}

	daysDiff := toDate.Sub(fromDate).Hours() / 24
	if daysDiff > 90 {
		return nil, fmt.Errorf("date range cannot exceed 90 days")
	}

	urlParams := map[string]string{
		"from": *params.From,
		"to":   *params.To,
	}

	if params.Symbol != nil {
		urlParams["symbol"] = *params.Symbol
	}

	var result []COTAnalysisResponse
	err = c.doRequest(c.BaseURL+"/commitment-of-traders-analysis", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
