package go_fmp

import (
	"fmt"
	"time"
)

// InsiderTradingLatestResponse represents the response from the Insider Trading Latest API
type InsiderTradingLatestResponse struct {
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

// GetInsiderTradingLatest retrieves the most recent insider trading transactions for a specific symbol
func (c *Client) GetInsiderTradingLatest(symbol string) ([]InsiderTradingLatestResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	url := c.BaseURL + "/insider-trading-latest"
	params := map[string]string{"symbol": symbol}
	var result []InsiderTradingLatestResponse
	if err := c.doRequest(url, params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
