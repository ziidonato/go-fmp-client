package go_fmp

import (
	"fmt"
	"time"
)

// TreasuryRatesResponse represents the response from the Treasury Rates API
type TreasuryRatesResponse struct {
	Date      time.Time  `json:"date"`
	Month1    float64 `json:"month1"`
	Month2    float64 `json:"month2"`
	Month3    float64 `json:"month3"`
	Month6    float64 `json:"month6"`
	Year1     float64 `json:"year1"`
	Year2     float64 `json:"year2"`
	Year3     float64 `json:"year3"`
	Year5     float64 `json:"year5"`
	Year7     float64 `json:"year7"`
	Year10    float64 `json:"year10"`
	Year20    float64 `json:"year20"`
	Year30    float64 `json:"year30"`
}

// TreasuryRates retrieves real-time and historical Treasury rates for all maturities
func (c *Client) TreasuryRates() ([]TreasuryRatesResponse, error) {
	var result []TreasuryRatesResponse
	err := c.doRequest(c.BaseURL+"/treasury-rates", map[string]string{}, &result)
	if err != nil {
		return nil, err
	}

	return result, err
}
