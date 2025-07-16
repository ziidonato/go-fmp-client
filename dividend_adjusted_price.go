package go_fmp

import (
	"encoding/json"
	"fmt"
)

// DividendAdjustedPriceParams represents the parameters for the Dividend Adjusted Price Chart API
type DividendAdjustedPriceParams struct {
	Symbol string  `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	From   *string `json:"from"`   // Optional: Start date (e.g., "2025-01-10")
	To     *string `json:"to"`     // Optional: End date (e.g., "2025-04-10")
}

// DividendAdjustedPriceResponse represents the response from the Dividend Adjusted Price Chart API
type DividendAdjustedPriceResponse struct {
	Symbol   string  `json:"symbol"`
	Date     string  `json:"date"`
	AdjOpen  float64 `json:"adjOpen"`
	AdjHigh  float64 `json:"adjHigh"`
	AdjLow   float64 `json:"adjLow"`
	AdjClose float64 `json:"adjClose"`
	Volume   int64   `json:"volume"`
}

// DividendAdjustedPrice retrieves stock price and volume data with dividend adjustments
func (c *Client) DividendAdjustedPrice(params DividendAdjustedPriceParams) ([]DividendAdjustedPriceResponse, error) {
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

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/historical-price-eod/dividend-adjusted", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []DividendAdjustedPriceResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
