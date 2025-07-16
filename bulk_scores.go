package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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
}

// GetScoresBulk retrieves key financial scores and metrics for multiple symbols
func (c *Client) GetScoresBulk() ([]ScoresBulkResponse, error) {
	// Build the URL
	baseURL := "https://financialmodelingprep.com/stable/scores-bulk"
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
	var result []ScoresBulkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
