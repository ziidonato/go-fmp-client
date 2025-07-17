package go_fmp

import (
	"fmt"
	"time"
)

// InstitutionalOwnershipSymbolPositionsSummaryParams represents the parameters for the Institutional Ownership Symbol Positions Summary API
type InstitutionalOwnershipSymbolPositionsSummaryParams struct {
	Symbol  string  `json:"symbol"`  // Required: Stock symbol (e.g., "AAPL")
	Year    string  `json:"year"`    // Required: Year (e.g., "2024")
	Quarter string  `json:"quarter"` // Required: Quarter (e.g., "Q1")
	Page    *int    `json:"page"`    // Optional: Page number (default: 0)
}

// InstitutionalOwnershipSymbolPositionsSummaryResponse represents the response from the Institutional Ownership Symbol Positions Summary API
type InstitutionalOwnershipSymbolPositionsSummaryResponse struct {
	Symbol                   string  `json:"symbol"`
	CIK                      string  `json:"cik"`
	Date                     time.Time  `json:"date"`
	InstitutionalInvestors   int     `json:"institutionalInvestors"`
	IncreasedPositions       int     `json:"increasedPositions"`
	DecreasedPositions       int     `json:"decreasedPositions"`
	NewPositions             int     `json:"newPositions"`
	ClosedPositions          int     `json:"closedPositions"`
	TotalSharesHeld          float64 `json:"totalSharesHeld"`
	TotalCallSharesHeld      float64 `json:"totalCallSharesHeld"`
	TotalPutSharesHeld       float64 `json:"totalPutSharesHeld"`
	TotalInvestedShares      float64 `json:"totalInvestedShares"`
	TotalCalls               float64 `json:"totalCalls"`
	TotalPuts                float64 `json:"totalPuts"`
	AvgSharesHeld            float64 `json:"avgSharesHeld"`
	AvgCallSharesHeld        float64 `json:"avgCallSharesHeld"`
	AvgPutSharesHeld         float64 `json:"avgPutSharesHeld"`
	AvgInvestedShares        float64 `json:"avgInvestedShares"`
	AvgCalls                 float64 `json:"avgCalls"`
	AvgPuts                  float64 `json:"avgPuts"`
	TotalNewInvestedShares   float64 `json:"totalNewInvestedShares"`
	TotalSoldOutShares       float64 `json:"totalSoldOutShares"`
	TotalIncreasedShares     float64 `json:"totalIncreasedShares"`
	TotalReducedShares       float64 `json:"totalReducedShares"`
	AvgNewInvestedShares     float64 `json:"avgNewInvestedShares"`
	AvgSoldOutShares         float64 `json:"avgSoldOutShares"`
	AvgIncreasedShares       float64 `json:"avgIncreasedShares"`
	AvgReducedShares         float64 `json:"avgReducedShares"`
	PctOfPortfolio           float64 `json:"pctOfPortfolio"`
	PrevPctOfPortfolio       float64 `json:"prevPctOfPortfolio"`
	PortfolioPercentageChange float64 `json:"portfolioPercentageChange"`
	HoldingsValue            float64 `json:"holdingsValue"`
	PrevHoldingsValue        float64 `json:"prevHoldingsValue"`
	HoldingsValueChange      float64 `json:"holdingsValueChange"`
	HoldingsValuePctChange   float64 `json:"holdingsValuePctChange"`
	SharesChange             float64 `json:"sharesChange"`
	SharesChangePct          float64 `json:"sharesChangePct"`
	TotalOwnership           float64 `json:"totalOwnership"`
	PrevTotalOwnership       float64 `json:"prevTotalOwnership"`
	OwnershipPctChange       float64 `json:"ownershipPctChange"`
	TotalFloat               float64 `json:"totalFloat"`
	PrevTotalFloat           float64 `json:"prevTotalFloat"`
	FloatPctChange           float64 `json:"floatPctChange"`
	TotalOutstandingShares   float64 `json:"totalOutstandingShares"`
	PrevOutstandingShares    float64 `json:"prevOutstandingShares"`
	OutstandingPctChange     float64 `json:"outstandingPctChange"`
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

	var result []InstitutionalOwnershipSymbolPositionsSummaryResponse
	err := c.doRequest(c.BaseURL+"/institutional-ownership/symbol-positions-summary", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
