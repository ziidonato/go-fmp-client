package go_fmp

import (
	"fmt"
	"time"
)

// InsiderTradingSearchResponse represents the response from the Insider Trading Search API
type InsiderTradingSearchResponse struct {
	Symbol          string    `json:"symbol"`
	FilingDate      time.Time `json:"filingDate"`
	TransactionDate time.Time `json:"transactionDate"`
	ReporterName    string    `json:"reporterName"`
	TransactionType TransactionType    `json:"transactionType"`
	SecurityName    string    `json:"securityName"`
	TypeOfOwner     TypeOfOwner    `json:"typeOfOwner"`
	AcquiredOrDisposed AcquiredOrDisposed `json:"acquiredOrDisposed"`
	Price           float64   `json:"price"`
	Quantity        float64   `json:"quantity"`
	OwnedShares     float64   `json:"ownedShares"`
	Link            string    `json:"link"`
}

// GetInsiderTradingSearch searches for insider trading transactions based on symbol
func (c *Client) GetInsiderTradingSearch(symbol string) ([]InsiderTradingSearchResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	url := c.BaseURL + "/insider-trading-search"
	params := map[string]string{"symbol": symbol}
	var result []InsiderTradingSearchResponse
	if err := c.doRequest(url, params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
