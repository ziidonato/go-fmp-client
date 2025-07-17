package go_fmp

import (
	"fmt"
	"time"
)

// SecInsiderTradingStatisticsParams represents the parameters for the SEC Insider Trading Statistics API
type SecInsiderTradingStatisticsParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// SecInsiderTradingStatisticsResponse represents the response from the SEC Insider Trading Statistics API
type SecInsiderTradingStatisticsResponse struct {
	Symbol            string    `json:"symbol"`
	CIK               string    `json:"cik"`
	Date              time.Time `json:"date"`
	PurchaseCount     int       `json:"purchaseCount"`
	PurchaseAmount    float64   `json:"purchaseAmount"`
	SaleCount         int       `json:"saleCount"`
	SaleAmount        float64   `json:"saleAmount"`
	TotalCount        int       `json:"totalCount"`
	TotalAmount       float64   `json:"totalAmount"`
}

// SecInsiderTradingStatistics retrieves SEC insider trading statistics for a specific stock
func (c *Client) SecInsiderTradingStatistics(params SecInsiderTradingStatisticsParams) ([]SecInsiderTradingStatisticsResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	var result []SecInsiderTradingStatisticsResponse
	err := c.doRequest(c.BaseURL+"/insider-trading-statistics", urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}