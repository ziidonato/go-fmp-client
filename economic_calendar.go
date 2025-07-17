package go_fmp

import (
	"time"
)

// EconomicCalendarResponse represents the response from the Economic Calendar API
type EconomicCalendarResponse struct {
	Event      string    `json:"event"`
	Date       time.Time `json:"date"`
	Country    string    `json:"country"`
	Actual     float64   `json:"actual"`
	Previous   float64   `json:"previous"`
	Change     float64   `json:"change"`
	ChangePct  float64   `json:"changePercentage"`
	Estimate   float64   `json:"estimate"`
	Impact     string    `json:"impact"`
}

// EconomicCalendar retrieves upcoming economic events and their impact on the markets
func (c *Client) EconomicCalendar(from, to string) ([]EconomicCalendarResponse, error) {
	urlParams := map[string]string{}

	if from != "" {
		urlParams["from"] = from
	}

	if to != "" {
		urlParams["to"] = to
	}

	var result []EconomicCalendarResponse
	err := c.doRequest(c.BaseURL+"/economic_calendar", urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
