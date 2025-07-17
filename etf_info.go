package go_fmp

import (
	"fmt"
	"time"
)

// ETFInfoParams represents the parameters for the ETF Info API
type ETFInfoParams struct {
	Symbol string `json:"symbol"` // Required: ETF symbol (e.g., "SPY")
}

// SectorExposure represents sector exposure data
type SectorExposure struct {
	Industry string  `json:"industry"`
	Exposure float64 `json:"exposure"`
}

// ETFPerformance represents ETF performance metrics
type ETFPerformance struct {
	YearToDateReturn       float64 `json:"yearToDateReturn"`
	OneYearReturn          float64 `json:"oneYearReturn"`
	ThreeYearReturn        float64 `json:"threeYearReturn"`
	FiveYearReturn         float64 `json:"fiveYearReturn"`
	MaxDrawDown            float64 `json:"maxDrawDown"`
	VolatilityOneYear      float64 `json:"volatilityOneYear"`
	VolatilityThreeYear    float64 `json:"volatilityThreeYear"`
	VolatilityFiveYear     float64 `json:"volatilityFiveYear"`
	SharpeRatioOneYear     float64 `json:"sharpeRatioOneYear"`
	SharpeRatioThreeYear   float64 `json:"sharpeRatioThreeYear"`
	SharpeRatioFiveYear    float64 `json:"sharpeRatioFiveYear"`
}

// ETFInfoResponse represents the response from the ETF Info API
type ETFInfoResponse struct {
	Symbol                string           `json:"symbol"`
	CompanyName           string           `json:"companyName"`
	Industry              string           `json:"industry"`
	Description           string           `json:"description"`
	IssuerName            string           `json:"issuerName"`
	InceptionDate         time.Time        `json:"inceptionDate"`
	CUSIP                 string           `json:"cusip"`
	ISIN                  string           `json:"isin"`
	DomicileCountry       string           `json:"domicileCountry"`
	AssetClass            AssetClass       `json:"assetClass"`
	UpdatedAt             time.Time        `json:"updatedAt"`
	AverageVolume         int64            `json:"averageVolume"`
	AverageDailyVolume    int64            `json:"averageDailyVolume"`
	NetAssetValue         float64          `json:"netAssetValue"`
	ExpenseRatio          float64          `json:"expenseRatio"`
	PortfolioDiversified  bool             `json:"portfolioDiversified"`
	PortfolioLeveraged    bool             `json:"portfolioLeveraged"`
	PortfolioInverse      bool             `json:"portfolioInverse"`
	LeverageFactor        float64          `json:"leverageFactor"`
	IndexTracked          string           `json:"indexTracked"`
	Performance           ETFPerformance   `json:"performance"`
}

// ETFInfo retrieves comprehensive data on ETFs and mutual funds
func (c *Client) ETFInfo(params ETFInfoParams) ([]ETFInfoResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	var result []ETFInfoResponse
	err := c.doRequest(c.BaseURL+"/etf/info", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
