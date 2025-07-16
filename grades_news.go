package go_fmp

import (
	"fmt"
)

// GradesNewsParams represents the parameters for the Stock Grade News API
type GradesNewsParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Page   *int   `json:"page"`   // Required: Page number (e.g., 0)
	Limit  *int   `json:"limit"`  // Required: Number of results (Maximum 100 records per request)
}

// GradesNewsResponse represents the response from the Stock Grade News API
type GradesNewsResponse struct {
	Symbol          string  `json:"symbol"`
	PublishedDate   string  `json:"publishedDate"`
	NewsURL         string  `json:"newsURL"`
	NewsTitle       string  `json:"newsTitle"`
	NewsBaseURL     string  `json:"newsBaseURL"`
	NewsPublisher   string  `json:"newsPublisher"`
	NewGrade        string  `json:"newGrade"`
	PreviousGrade   string  `json:"previousGrade"`
	GradingCompany  string  `json:"gradingCompany"`
	Action          string  `json:"action"`
	PriceWhenPosted float64 `json:"priceWhenPosted"`
}

// GradesNews retrieves real-time updates on stock rating changes
func (c *Client) GradesNews(params GradesNewsParams) ([]GradesNewsResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	if params.Page == nil {
		return nil, fmt.Errorf("page parameter is required")
	}

	if params.Limit == nil {
		return nil, fmt.Errorf("limit parameter is required")
	}

	if *params.Limit > 100 {
		return nil, fmt.Errorf("limit cannot exceed 100")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
		"page":   fmt.Sprintf("%d", *params.Page),
		"limit":  fmt.Sprintf("%d", *params.Limit),
	}

	var result []GradesNewsResponse
	err := c.doRequest(c.BaseURL+"/grades-news", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
