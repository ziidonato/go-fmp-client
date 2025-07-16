package go_fmp

import (
	"encoding/json"
	"fmt"
)

// UnadjustedStockPriceParams represents the parameters for the Unadjusted Stock Price API
type UnadjustedStockPriceParams struct {
	Symbol string  `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	From   *string `json:"from"`   // Optional: Start date (e.g., "2025-01-10")
	To     *string `json:"to"`     // Optional: End date (e.g., "2025-04-10")
}

// UnadjustedStockPriceResponse represents the response from the Unadjusted Stock Price API
type UnadjustedStockPriceResponse struct {
	Symbol   string  `json:"symbol"`
	Date     string  `json:"date"`
	AdjOpen  float64 `json:"adjOpen"`
	AdjHigh  float64 `json:"adjHigh"`
	AdjLow   float64 `json:"adjLow"`
	AdjClose float64 `json:"adjClose"`
	Volume   int64   `json:"volume"`
}

// UnadjustedStockPrice retrieves stock price and volume data without adjustments for stock splits
func (c *Client) UnadjustedStockPrice(params UnadjustedStockPriceParams) ([]UnadjustedStockPriceResponse, error) {
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

	return doRequest[[]UnadjustedStockPriceResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
}
