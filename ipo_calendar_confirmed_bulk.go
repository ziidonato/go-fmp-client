package go_fmp

import (
	"fmt"
	"time"
)

// IPOCalendarConfirmedBulkParams represents the parameters for the IPO Calendar Confirmed Bulk API
type IPOCalendarConfirmedBulkParams struct {
	Year string `json:"year"` // Required: Year (e.g., "2024")
}

// IPOCalendarConfirmedBulkResponse represents the response from the IPO Calendar Confirmed Bulk API
type IPOCalendarConfirmedBulkResponse struct {
	Symbol         string    `json:"symbol"`
	Company        string    `json:"company"`
	Date           time.Time `json:"date"`
	Exchange       string    `json:"exchange"`
	PriceRange     string    `json:"priceRange"`
	Price          float64   `json:"price"`
	Currency       string    `json:"currency"`
	Shares         float64   `json:"shares"`
}

// IPOCalendarConfirmedBulk retrieves all confirmed IPO calendar data for a specific year
func (c *Client) IPOCalendarConfirmedBulk(params IPOCalendarConfirmedBulkParams) ([]IPOCalendarConfirmedBulkResponse, error) {
	if params.Year == "" {
		return nil, fmt.Errorf("year parameter is required")
	}

	url := fmt.Sprintf("ipo-calendar-confirmed-bulk/%s", params.Year)
	
	var result []IPOCalendarConfirmedBulkResponse
	err := c.doRequest(c.BaseURL+"/"+url, nil, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}