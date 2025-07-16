package go_fmp

import (
	"encoding/json"
	"fmt"
)

// GradesLatestNewsParams represents the parameters for the Stock Grade Latest News API
type GradesLatestNewsParams struct {
	Page  *int `json:"page"`  // Required: Page number (Page maxed at 100)
	Limit *int `json:"limit"` // Required: Number of results (Maximum 1000 records per request)
}

// GradesLatestNewsResponse represents the response from the Stock Grade Latest News API
type GradesLatestNewsResponse struct {
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

// GradesLatestNews retrieves the most recent updates on analyst ratings for all stock symbols
func (c *Client) GradesLatestNews(params GradesLatestNewsParams) ([]GradesLatestNewsResponse, error) {
	if params.Page == nil {
		return nil, fmt.Errorf("page parameter is required")
	}

	if *params.Page > 100 {
		return nil, fmt.Errorf("page cannot exceed 100")
	}

	if params.Limit == nil {
		return nil, fmt.Errorf("limit parameter is required")
	}

	if *params.Limit > 1000 {
		return nil, fmt.Errorf("limit cannot exceed 1000")
	}

	urlParams := map[string]string{
		"page":  fmt.Sprintf("%d", *params.Page),
		"limit": fmt.Sprintf("%d", *params.Limit),
	}

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/grades-latest-news", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []GradesLatestNewsResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
