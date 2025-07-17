package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	// Using direct HTTP request since it returns a file
	url := fmt.Sprintf("%s/ipo-calendar-confirmed-bulk/%s", c.BaseURL, params.Year)
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	if c.APIKey != "" {
		req.URL.Query().Add("apikey", c.APIKey)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code: %d", resp.StatusCode)
	}

	var result []IPOCalendarConfirmedBulkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}