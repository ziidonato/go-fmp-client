package go_fmp

import (
	"encoding/json"
	"fmt"
)

// StockPeersParams represents the parameters for the Stock Peer Comparison API
type StockPeersParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")

// StockPeersResponse represents the response from the Stock Peer Comparison API
type StockPeersResponse struct {
	Symbol      string  `json:"symbol"`
	CompanyName string  `json:"companyName"`
	Price       float64 `json:"price"`
	MktCap      int64   `json:"mktCap"`

// StockPeers identifies and compares companies within the same sector and market capitalization range
func (c *Client) StockPeers(params StockPeersParams) ([]StockPeersResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	return doRequest[[]StockPeersResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
