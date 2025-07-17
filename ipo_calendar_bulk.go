package go_fmp

import (
	"fmt"
	"time"
)

// IPOCalendarBulkParams represents the parameters for the IPO Calendar Bulk API
type IPOCalendarBulkParams struct {
	Year string `json:"year"` // Required: Year (e.g., "2024")
}

// IPOCalendarBulkResponse represents the response from the IPO Calendar Bulk API
type IPOCalendarBulkResponse struct {
	Symbol         string    `json:"symbol"`
	Company        string    `json:"company"`
	Date           time.Time `json:"date"`
	Exchange       string    `json:"exchange"`
	PriceRange     string    `json:"priceRange"`
	Price          float64   `json:"price"`
	Currency       string    `json:"currency"`
	Shares         float64   `json:"shares"`
}

// IPOCalendarBulk retrieves all IPO calendar data for a specific year
func (c *Client) IPOCalendarBulk(params IPOCalendarBulkParams) ([]IPOCalendarBulkResponse, error) {
	if params.Year == "" {
		return nil, fmt.Errorf("year parameter is required")
	}

	url := fmt.Sprintf("ipo-calendar-bulk/%s", params.Year)
	
	var result []IPOCalendarBulkResponse
	err := c.doRequest(c.BaseURL+"/"+url, nil, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}