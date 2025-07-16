package go_fmp

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
func (c *Client) CommoditiesQuote(symbol string) ([]CommoditiesQuoteResponse, error) {
	url := "https://financialmodelingprep.com/stable/quote"
	params := map[string]string{"symbol": symbol}
	var result []CommoditiesQuoteResponse
	if err := c.doRequest(url, params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
