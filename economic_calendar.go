package go_fmp

import (
	"encoding/json"
)

// EconomicCalendarResponse represents the response from the Economic Data Releases Calendar API
type EconomicCalendarResponse struct {
	Date      string   `json:"date"`
	Time      string   `json:"time"`
	Country   string   `json:"country"`
	Event     string   `json:"event"`
	Reference string   `json:"reference"`
	Source    string   `json:"source"`
	Actual    *float64 `json:"actual"`
	Previous  *float64 `json:"previous"`
	Forecast  *float64 `json:"forecast"`
	Impact    string   `json:"impact"`
	Currency  string   `json:"currency"`
	Unit      string   `json:"unit"`

// EconomicCalendar retrieves a comprehensive calendar of upcoming economic data releases
func (c *Client) EconomicCalendar() ([]EconomicCalendarResponse, error) {
	return doRequest[[]EconomicCalendarResponse](c, "https://financialmodelingprep.com/stable/economic-calendar", nil)
