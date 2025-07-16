package go_fmp

import (
	"fmt"
)

// GradesConsensusParams represents the parameters for the Stock Grades Summary API
type GradesConsensusParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// GradesConsensusResponse represents the response from the Stock Grades Summary API
type GradesConsensusResponse struct {
	Symbol     string `json:"symbol"`
	StrongBuy  int    `json:"strongBuy"`
	Buy        int    `json:"buy"`
	Hold       int    `json:"hold"`
	Sell       int    `json:"sell"`
	StrongSell int    `json:"strongSell"`
	Consensus  string `json:"consensus"`
}

// GradesConsensus retrieves an overall view of analyst ratings for individual stock symbols
func (c *Client) GradesConsensus(params GradesConsensusParams) ([]GradesConsensusResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	var result []GradesConsensusResponse
	err := c.doRequest("https://financialmodelingprep.com/stable/grades-consensus", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
