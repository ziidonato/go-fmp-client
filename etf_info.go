package go_fmp

import (
	"encoding/json"
	"fmt"
)

// ETFInfoParams represents the parameters for the ETF & Mutual Fund Information API
type ETFInfoParams struct {
	Symbol string `json:"symbol"` // Required: ETF/Fund symbol (e.g., "SPY")
}

// SectorExposure represents sector exposure data
type SectorExposure struct {
	Industry string  `json:"industry"`
	Exposure float64 `json:"exposure"`
}

// ETFInfoResponse represents the response from the ETF & Mutual Fund Information API
type ETFInfoResponse struct {
	Symbol                string           `json:"symbol"`
	Name                  string           `json:"name"`
	Description           string           `json:"description"`
	ISIN                  string           `json:"isin"`
	AssetClass            string           `json:"assetClass"`
	SecurityCusip         string           `json:"securityCusip"`
	Domicile              string           `json:"domicile"`
	Website               string           `json:"website"`
	ETFCompany            string           `json:"etfCompany"`
	ExpenseRatio          float64          `json:"expenseRatio"`
	AssetsUnderManagement int64            `json:"assetsUnderManagement"`
	AvgVolume             int64            `json:"avgVolume"`
	InceptionDate         string           `json:"inceptionDate"`
	NAV                   float64          `json:"nav"`
	NAVCurrency           string           `json:"navCurrency"`
	HoldingsCount         int              `json:"holdingsCount"`
	UpdatedAt             string           `json:"updatedAt"`
	SectorsList           []SectorExposure `json:"sectorsList"`
}

// ETFInfo retrieves comprehensive data on ETFs and mutual funds
func (c *Client) ETFInfo(params ETFInfoParams) ([]ETFInfoResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	return doRequest[[]ETFInfoResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
}
