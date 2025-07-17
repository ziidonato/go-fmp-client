package go_fmp

import (
	"fmt"
	"time"
)

// RatingRecommendationParams represents the parameters for the Rating Recommendation API
type RatingRecommendationParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// RatingRecommendationResponse represents the response from the Rating Recommendation API
type RatingRecommendationResponse struct {
	Symbol              string    `json:"symbol"`
	Date                time.Time `json:"date"`
	Rating              string    `json:"rating"`
	RatingScore         int       `json:"ratingScore"`
	RatingRecommendation string   `json:"ratingRecommendation"`
	RatingDetailsDCFScore          int       `json:"ratingDetailsDCFScore"`
	RatingDetailsDCFRecommendation string    `json:"ratingDetailsDCFRecommendation"`
	RatingDetailsROEScore          int       `json:"ratingDetailsROEScore"`
	RatingDetailsROERecommendation string    `json:"ratingDetailsROERecommendation"`
	RatingDetailsROAScore          int       `json:"ratingDetailsROAScore"`
	RatingDetailsROARecommendation string    `json:"ratingDetailsROARecommendation"`
	RatingDetailsDEScore           int       `json:"ratingDetailsDEScore"`
	RatingDetailsDERecommendation  string    `json:"ratingDetailsDERecommendation"`
	RatingDetailsPEScore           int       `json:"ratingDetailsPEScore"`
	RatingDetailsPERecommendation  string    `json:"ratingDetailsPERecommendation"`
	RatingDetailsPBScore           int       `json:"ratingDetailsPBScore"`
	RatingDetailsPBRecommendation  string    `json:"ratingDetailsPBRecommendation"`
}

// RatingRecommendation retrieves rating recommendations for a specific stock
func (c *Client) RatingRecommendation(params RatingRecommendationParams) ([]RatingRecommendationResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	var result []RatingRecommendationResponse
	err := c.doRequest(c.BaseURL+"/rating", urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}