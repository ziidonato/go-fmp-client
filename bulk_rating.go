package go_fmp

import (
	"fmt"
	"time"
)

// BulkRatingParams represents the parameters for the Bulk Rating API
type BulkRatingParams struct {
	Date string `json:"date"` // Required: date (e.g., "2024-10-22")
}

// BulkRatingResponse represents the response from the Bulk Rating API
type BulkRatingResponse struct {
	Symbol                   string `json:"symbol"`
	Date                    time.Time `json:"date"`
	Rating                  string `json:"rating"`
	RatingScore             float64 `json:"ratingScore"`
	RatingRecommendation    string `json:"ratingRecommendation"`
	DCFScore                float64 `json:"dcfScore"`
	DCFRecommendation       string `json:"dcfRecommendation"`
	ROEScore                float64 `json:"roeScore"`
	ROERecommendation       string `json:"roeRecommendation"`
	ROAScore                float64 `json:"roaScore"`
	ROARecommendation       string `json:"roaRecommendation"`
	DEScore                 float64 `json:"deScore"`
	DERecommendation        string `json:"deRecommendation"`
	PEScore                 float64 `json:"peScore"`
	PERecommendation        string `json:"peRecommendation"`
	PBScore                 float64 `json:"pbScore"`
	PBRecommendation        string `json:"pbRecommendation"`
}

// BulkRating retrieves bulk stock ratings for a specific date
func (c *Client) BulkRating(params BulkRatingParams) ([]BulkRatingResponse, error) {
	urlParams := map[string]string{
		"date": params.Date,
	}

	var result []BulkRatingResponse
	err := c.doRequest(c.BaseURL+"/bulk-rating", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to get bulk rating: %w", err)
	}

	return result, nil
}
