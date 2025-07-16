package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// ExchangeStockQuotesResponse represents the response from the exchange stock quotes API
type ExchangeStockQuotesResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
	Volume int64   `json:"volume"`
}

// GetExchangeStockQuotes retrieves real-time stock quotes for all listed stocks on a specific exchange
func (c *Client) GetExchangeStockQuotes(exchange string, short bool) ([]ExchangeStockQuotesResponse, error) {
	if exchange == "" {
		return nil, fmt.Errorf("exchange is required")
	}

	url := "https://financialmodelingprep.com/stable/batch-exchange-quote"

	resp, err := c.get(url, map[string]string{
		"exchange": exchange,
		"short":    strconv.FormatBool(short),
	})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var result []ExchangeStockQuotesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
