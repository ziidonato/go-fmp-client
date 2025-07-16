package go_fmp

import (
	"encoding/json"
	"fmt"
)

// InstitutionalOwnershipSymbolPositionsSummaryParams represents the parameters for the Institutional Ownership Symbol Positions Summary API
type InstitutionalOwnershipSymbolPositionsSummaryParams struct {
	Symbol  string `json:"symbol"`  // Required: Stock symbol (e.g., "AAPL")
	Year    string `json:"year"`    // Required: Year (e.g., "2023")
	Quarter string `json:"quarter"` // Required: Quarter (e.g., "3")
}

// InstitutionalOwnershipSymbolPositionsSummaryResponse represents the response from the Institutional Ownership Symbol Positions Summary API
type InstitutionalOwnershipSymbolPositionsSummaryResponse struct {
	Symbol                   string  `json:"symbol"`
	CIK                      string  `json:"cik"`
	Date                     string  `json:"date"`
	InvestorsHolding         int     `json:"investorsHolding"`
	LastInvestorsHolding     int     `json:"lastInvestorsHolding"`
	InvestorsHoldingChange   int     `json:"investorsHoldingChange"`
	NumberOf13Fshares        int64   `json:"numberOf13Fshares"`
	LastNumberOf13Fshares    int64   `json:"lastNumberOf13Fshares"`
	NumberOf13FsharesChange  int64   `json:"numberOf13FsharesChange"`
	TotalInvested            int64   `json:"totalInvested"`
	LastTotalInvested        int64   `json:"lastTotalInvested"`
	TotalInvestedChange      int64   `json:"totalInvestedChange"`
	OwnershipPercent         float64 `json:"ownershipPercent"`
	LastOwnershipPercent     float64 `json:"lastOwnershipPercent"`
	OwnershipPercentChange   float64 `json:"ownershipPercentChange"`
	NewPositions             int     `json:"newPositions"`
	LastNewPositions         int     `json:"lastNewPositions"`
	NewPositionsChange       int     `json:"newPositionsChange"`
	IncreasedPositions       int     `json:"increasedPositions"`
	LastIncreasedPositions   int     `json:"lastIncreasedPositions"`
	IncreasedPositionsChange int     `json:"increasedPositionsChange"`
	ClosedPositions          int     `json:"closedPositions"`
	LastClosedPositions      int     `json:"lastClosedPositions"`
	ClosedPositionsChange    int     `json:"closedPositionsChange"`
	ReducedPositions         int     `json:"reducedPositions"`
	LastReducedPositions     int     `json:"lastReducedPositions"`
	ReducedPositionsChange   int     `json:"reducedPositionsChange"`
	TotalCalls               int64   `json:"totalCalls"`
	LastTotalCalls           int64   `json:"lastTotalCalls"`
	TotalCallsChange         int64   `json:"totalCallsChange"`
	TotalPuts                int64   `json:"totalPuts"`
	LastTotalPuts            int64   `json:"lastTotalPuts"`
	TotalPutsChange          int64   `json:"totalPutsChange"`
	PutCallRatio             float64 `json:"putCallRatio"`
	LastPutCallRatio         float64 `json:"lastPutCallRatio"`
	PutCallRatioChange       float64 `json:"putCallRatioChange"`
}

// InstitutionalOwnershipSymbolPositionsSummary retrieves positions summary for a specific stock symbol
func (c *Client) InstitutionalOwnershipSymbolPositionsSummary(params InstitutionalOwnershipSymbolPositionsSummaryParams) ([]InstitutionalOwnershipSymbolPositionsSummaryResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	if params.Year == "" {
		return nil, fmt.Errorf("year parameter is required")
	}

	if params.Quarter == "" {
		return nil, fmt.Errorf("quarter parameter is required")
	}

	urlParams := map[string]string{
		"symbol":  params.Symbol,
		"year":    params.Year,
		"quarter": params.Quarter,
	}

	resp, err := c.get("https://financialmodelingprep.com/stable/institutional-ownership/symbol-positions-summary", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []InstitutionalOwnershipSymbolPositionsSummaryResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
