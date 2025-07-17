package go_fmp

import (
	"time"
)

// EconomicCalendarResponse represents the response from the Economic Calendar API
type EconomicCalendarResponse struct {
	Event    string `json:"event"`
	Date time.Time `json:"date"`
	Country Country `json:"country"`
	Actual   string `json:"actual"`
	Previous string `json:"previous"`
	Estimate string `json:"estimate"`
	Impact EconomicImpact `json:"impact"`
	Currency Currency `json:"currency"`
}

// EconomicCalendar retrieves upcoming economic events and their impact on the markets
func (c *Client) EconomicCalendar() ([]EconomicCalendarResponse, error) {
	url := c.BaseURL + "/economic-calendar"
	var result []EconomicCalendarResponse
	if err := c.doRequest(url, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
