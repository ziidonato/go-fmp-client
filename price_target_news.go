package go_fmp

import (
	"fmt"
)

// PriceTargetNewsParams represents the parameters for the Price Target News API
type PriceTargetNewsParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Limit  *int   `json:"limit"`  // Required: Number of results (e.g., 10)
	Page   *int   `json:"page"`   // Required: Page number (e.g., 0)
}

// PriceTargetNewsResponse represents the response from the Price Target News API
type PriceTargetNewsResponse struct {
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

// PriceTargetNews retrieves real-time updates on analysts' price targets for stocks
func (c *Client) PriceTargetNews(params PriceTargetNewsParams) ([]PriceTargetNewsResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	if params.Limit == nil {
		return nil, fmt.Errorf("limit parameter is required")
	}

	if params.Page == nil {
		return nil, fmt.Errorf("page parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
		"limit":  fmt.Sprintf("%d", *params.Limit),
		"page":   fmt.Sprintf("%d", *params.Page),
	}

	var result []PriceTargetNewsResponse
	err := c.doRequest(c.BaseURL+"/price-target-news", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, err
}
