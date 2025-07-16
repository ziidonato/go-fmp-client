package go_fmp

import (
	"encoding/json"
	"fmt"
)

// GradesHistoricalParams represents the parameters for the Historical Stock Grades API
type GradesHistoricalParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Limit  *int   `json:"limit"`  // Optional: Number of results (Maximum 1000 records per request)
}

// GradesHistoricalResponse represents the response from the Historical Stock Grades API
type GradesHistoricalResponse struct {
	Symbol                   string `json:"symbol"`
	Date                     string `json:"date"`
	AnalystRatingsBuy        int    `json:"analystRatingsBuy"`
	AnalystRatingsHold       int    `json:"analystRatingsHold"`
	AnalystRatingsSell       int    `json:"analystRatingsSell"`
	AnalystRatingsStrongSell int    `json:"analystRatingsStrongSell"`
}

// GradesHistorical retrieves a comprehensive record of analyst grades for specific stock symbols
func (c *Client) GradesHistorical(params GradesHistoricalParams) ([]GradesHistoricalResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	if params.Limit != nil {
		if *params.Limit > 1000 {
			return nil, fmt.Errorf("limit cannot exceed 1000")
		}
		urlParams["limit"] = fmt.Sprintf("%d", *params.Limit)
	}

	resp, err := c.get("https://financialmodelingprep.com/stable/grades-historical", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []GradesHistoricalResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
