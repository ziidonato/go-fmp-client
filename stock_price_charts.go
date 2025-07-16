package go_fmp

import (
	"fmt"
)

// DailyPriceChartParams represents the parameters for daily stock price chart APIs
type DailyPriceChartParams struct {
	Symbol string  `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	From   *string `json:"from"`   // Optional: Start date (e.g., "2025-01-10")
	To     *string `json:"to"`     // Optional: End date (e.g., "2025-04-10")
}

// DailyPriceChartResponse represents the response from daily stock price chart APIs
type DailyPriceChartResponse struct {
	Symbol   string  `json:"symbol"`
	Date     string  `json:"date"`
	AdjOpen  float64 `json:"adjOpen"`
	AdjHigh  float64 `json:"adjHigh"`
	AdjLow   float64 `json:"adjLow"`
	AdjClose float64 `json:"adjClose"`
	Volume   int64   `json:"volume"`
}

// UnadjustedStockPriceChart retrieves stock price and volume data without adjustments for stock splits
func (c *Client) UnadjustedStockPriceChart(params DailyPriceChartParams) ([]DailyPriceChartResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	if params.From != nil {
		urlParams["from"] = *params.From
	}

	if params.To != nil {
		urlParams["to"] = *params.To
	}

	url := c.BaseURL + "/historical-price-eod/non-split-adjusted"
	var result []DailyPriceChartResponse
	if err := c.doRequest(url, urlParams, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// DividendAdjustedPriceChart retrieves stock price and volume data with dividend adjustments
func (c *Client) DividendAdjustedPriceChart(params DailyPriceChartParams) ([]DailyPriceChartResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	if params.From != nil {
		urlParams["from"] = *params.From
	}

	if params.To != nil {
		urlParams["to"] = *params.To
	}

	var result []DailyPriceChartResponse
	err := c.doRequest(c.BaseURL+"/historical-price-eod/dividend-adjusted", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
