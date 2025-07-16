package go_fmp

import (
	"fmt"
	"strconv"
)

// ExchangeStockQuotesResponse represents the response from the exchange stock quotes API
type ExchangeStockQuotesResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
	Volume int64   `json:"volume"`
}

// ExchangeStockQuotes retrieves real-time stock quotes for all listed stocks on a specific exchange
func (c *Client) ExchangeStockQuotes(exchange string, short bool) ([]ExchangeStockQuotesResponse, error) {
	if exchange == "" {
		return nil, fmt.Errorf("exchange is required")
	}

	url := c.BaseURL + "/exchange-stock-quotes"
	params := map[string]string{
		"exchange": exchange,
		"short":    strconv.FormatBool(short),
	}

	var result []ExchangeStockQuotesResponse
	if err := c.doRequest(url, params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
