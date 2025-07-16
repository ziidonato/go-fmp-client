package go_fmp

import (
	"encoding/json"
	"fmt"
)

// CommoditiesQuoteParams represents the parameters for the Commodities Quote API
type CommoditiesQuoteParams struct {
	Symbol string `json:"symbol"` // Required: Commodity symbol (e.g., "GCUSD")
}

// CommoditiesQuoteResponse represents the response from the Commodities Quote API
type CommoditiesQuoteResponse struct {
	Symbol           string  `json:"symbol"`
	Name             string  `json:"name"`
	Price            float64 `json:"price"`
	ChangePercentage float64 `json:"changePercentage"`
	Change           float64 `json:"change"`
	Volume           int64   `json:"volume"`
	DayLow           float64 `json:"dayLow"`
	DayHigh          float64 `json:"dayHigh"`
	YearHigh         float64 `json:"yearHigh"`
	YearLow          float64 `json:"yearLow"`
	MarketCap        *int64  `json:"marketCap"`
	PriceAvg50       float64 `json:"priceAvg50"`
	PriceAvg200      float64 `json:"priceAvg200"`
	Exchange         string  `json:"exchange"`
	Open             float64 `json:"open"`
	PreviousClose    float64 `json:"previousClose"`
	Timestamp        int64   `json:"timestamp"`
}

// CommoditiesQuote retrieves real-time price quotes for commodities traded worldwide
func (c *Client) CommoditiesQuote(params CommoditiesQuoteParams) ([]CommoditiesQuoteResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	resp, err := c.get("https://financialmodelingprep.com/stable/quote", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []CommoditiesQuoteResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
