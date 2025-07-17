package go_fmp

import (
	"fmt"
)

// PriceTargetLatestNewsParams represents the parameters for the Price Target Latest News API
type PriceTargetLatestNewsParams struct {
	Limit *int `json:"limit"` // Required: Number of results (Maximum 1000 records per request)
	Page  *int `json:"page"`  // Required: Page number (Page maxed at 100)
}

// PriceTargetLatestNewsResponse represents the response from the Price Target Latest News API
type PriceTargetLatestNewsResponse struct {
	Symbol          string  `json:"symbol"`
	PublishedDate   string  `json:"publishedDate"`
	NewsURL         string  `json:"newsURL"`
	NewsTitle       string  `json:"newsTitle"`
	AnalystName     string  `json:"analystName"`
	PriceTarget     float64 `json:"priceTarget"`
	AdjPriceTarget  float64 `json:"adjPriceTarget"`
	PriceWhenPosted float64 `json:"priceWhenPosted"`
	NewsPublisher   string  `json:"newsPublisher"`
	NewsBaseURL     string  `json:"newsBaseURL"`
	AnalystCompany  string  `json:"analystCompany"`
}

// PriceTargetLatestNews retrieves the most recent analyst price target updates for all stock symbols
func (c *Client) PriceTargetLatestNews(params PriceTargetLatestNewsParams) ([]PriceTargetLatestNewsResponse, error) {
	if params.Limit == nil {
		return nil, fmt.Errorf("limit parameter is required")
	}

	if *params.Limit > 1000 {
		return nil, fmt.Errorf("limit cannot exceed 1000")
	}

	if params.Page == nil {
		return nil, fmt.Errorf("page parameter is required")
	}

	if *params.Page > 100 {
		return nil, fmt.Errorf("page cannot exceed 100")
	}

	urlParams := map[string]string{
		"limit": fmt.Sprintf("%d", *params.Limit),
		"page":  fmt.Sprintf("%d", *params.Page),
	}

	var result []PriceTargetLatestNewsResponse
	err := c.doRequest(c.BaseURL+"/price-target-latest-news", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
