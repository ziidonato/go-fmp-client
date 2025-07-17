package go_fmp

import (
	"time"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// RatingBulkResponse represents the response from the Rating Bulk API
type RatingBulkResponse struct {
	Symbol                  string `json:"symbol"`
	Date time.Time `json:"date"`
	Rating                  string `json:"rating"`
	DiscountedCashFlowScore string `json:"discountedCashFlowScore"`
	ReturnOnEquityScore     string `json:"returnOnEquityScore"`
	ReturnOnAssetsScore     string `json:"returnOnAssetsScore"`
	DebtToEquityScore       string `json:"debtToEquityScore"`
	PriceToEarningsScore    string `json:"priceToEarningsScore"`
	PriceToBookScore        string `json:"priceToBookScore"`
}

// GetRatingBulk retrieves comprehensive rating data for multiple stocks
func (c *Client) GetRatingBulk() ([]RatingBulkResponse, error) {
	// Build the URL
	baseURL := c.BaseURL + "/rating-bulk"
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %w", err)
	}

	// Add API key if available
	if c.APIKey != "" {
		q := u.Query()
		q.Set("apikey", c.APIKey)
		u.RawQuery = q.Encode()
	}

	// Create the request
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set headers
	req.Header.Set("User-Agent", "fmp-go-client")
	req.Header.Set("Accept", "application/json")

	// Make the request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	// Parse the response
	var result []RatingBulkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
