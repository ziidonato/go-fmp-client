package go_fmp

import (
	"fmt"
	"time"
)

// AnalystRecommendationsParams represents the parameters for the Analyst Recommendations API
type AnalystRecommendationsParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// AnalystRecommendationsResponse represents the response from the Analyst Recommendations API
type AnalystRecommendationsResponse struct {
	Symbol                       string    `json:"symbol"`
	Date                         time.Time `json:"date"`
	AnalystCount                 int       `json:"analystCount"`
	AnalystRecommendationBuy     int       `json:"analystRecommendationBuy"`
	AnalystRecommendationHold    int       `json:"analystRecommendationHold"`
	AnalystRecommendationSell    int       `json:"analystRecommendationSell"`
	AnalystRecommendationStrongBuy  int    `json:"analystRecommendationStrongBuy"`
	AnalystRecommendationStrongSell int    `json:"analystRecommendationStrongSell"`
}

// AnalystRecommendations retrieves analyst recommendations for a specific stock
func (c *Client) AnalystRecommendations(params AnalystRecommendationsParams) ([]AnalystRecommendationsResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	var result []AnalystRecommendationsResponse
	err := c.doRequest(c.BaseURL+"/analyst-stock-recommendations", urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}