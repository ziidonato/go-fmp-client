package go_fmp

import (
	"fmt"
	"time"
)

// HistoricalMarketCapitalizationParams represents the parameters for the Historical Market Capitalization API
type HistoricalMarketCapitalizationParams struct {
	Symbol string  `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Limit  *int    `json:"limit"`  // Optional: Number of records to return
	From   *string `json:"from"`   // Optional: Start date (YYYY-MM-DD format)
	To     *string `json:"to"`     // Optional: End date (YYYY-MM-DD format)
}

// HistoricalMarketCapitalizationResponse represents the response from the Historical Market Capitalization API
type HistoricalMarketCapitalizationResponse struct {
	Symbol               string    `json:"symbol"`
	Date                 time.Time `json:"date"`
	MarketCap            int64     `json:"marketCap"`
}

// HistoricalMarketCapitalization retrieves historical market capitalization data for a company
func (c *Client) HistoricalMarketCapitalization(params HistoricalMarketCapitalizationParams) ([]HistoricalMarketCapitalizationResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	if params.Limit != nil {
		if *params.Limit > 5000 {
			return nil, fmt.Errorf("limit cannot exceed 5000")
		}
		urlParams["limit"] = fmt.Sprintf("%d", *params.Limit)
	}

	if params.From != nil {
		urlParams["from"] = *params.From
	}

	if params.To != nil {
		urlParams["to"] = *params.To
	}

	var result []HistoricalMarketCapitalizationResponse
	err := c.doRequest(c.BaseURL+"/historical-market-capitalization", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
