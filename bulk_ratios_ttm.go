package go_fmp

import (
	"encoding/json"
	"fmt"
)

// RatiosTTMBulkResponse represents the response from the Ratios TTM Bulk API
type RatiosTTMBulkResponse struct {
	Symbol                                     string `json:"symbol"`
	GrossProfitMarginTTM                       string `json:"grossProfitMarginTTM"`
	EBITMarginTTM                              string `json:"ebitMarginTTM"`
	EBITDAMarginTTM                            string `json:"ebitdaMarginTTM"`
	OperatingProfitMarginTTM                   string `json:"operatingProfitMarginTTM"`
	PretaxProfitMarginTTM                      string `json:"pretaxProfitMarginTTM"`
	ContinuousOperationsProfitMarginTTM        string `json:"continuousOperationsProfitMarginTTM"`
	NetProfitMarginTTM                         string `json:"netProfitMarginTTM"`
	BottomLineProfitMarginTTM                  string `json:"bottomLineProfitMarginTTM"`
	ReceivablesTurnoverTTM                     string `json:"receivablesTurnoverTTM"`
	PayablesTurnoverTTM                        string `json:"payablesTurnoverTTM"`
	InventoryTurnoverTTM                       string `json:"inventoryTurnoverTTM"`
	FixedAssetTurnoverTTM                      string `json:"fixedAssetTurnoverTTM"`
	AssetTurnoverTTM                           string `json:"assetTurnoverTTM"`
	CurrentRatioTTM                            string `json:"currentRatioTTM"`
	QuickRatioTTM                              string `json:"quickRatioTTM"`
	SolvencyRatioTTM                           string `json:"solvencyRatioTTM"`
	CashRatioTTM                               string `json:"cashRatioTTM"`
	PriceToEarningsRatioTTM                    string `json:"priceToEarningsRatioTTM"`
	PriceToEarningsGrowthRatioTTM              string `json:"priceToEarningsGrowthRatioTTM"`
	ForwardPriceToEarningsGrowthRatioTTM       string `json:"forwardPriceToEarningsGrowthRatioTTM"`
	PriceToBookRatioTTM                        string `json:"priceToBookRatioTTM"`
	PriceToSalesRatioTTM                       string `json:"priceToSalesRatioTTM"`
	PriceToFreeCashFlowRatioTTM                string `json:"priceToFreeCashFlowRatioTTM"`
	PriceToOperatingCashFlowRatioTTM           string `json:"priceToOperatingCashFlowRatioTTM"`
	DebtToAssetsRatioTTM                       string `json:"debtToAssetsRatioTTM"`
	DebtToEquityRatioTTM                       string `json:"debtToEquityRatioTTM"`
	DebtToCapitalRatioTTM                      string `json:"debtToCapitalRatioTTM"`
	LongTermDebtToCapitalRatioTTM              string `json:"longTermDebtToCapitalRatioTTM"`
	FinancialLeverageRatioTTM                  string `json:"financialLeverageRatioTTM"`
	WorkingCapitalTurnoverRatioTTM             string `json:"workingCapitalTurnoverRatioTTM"`
	OperatingCashFlowRatioTTM                  string `json:"operatingCashFlowRatioTTM"`
	OperatingCashFlowSalesRatioTTM             string `json:"operatingCashFlowSalesRatioTTM"`
	FreeCashFlowOperatingCashFlowRatioTTM      string `json:"freeCashFlowOperatingCashFlowRatioTTM"`
	DebtServiceCoverageRatioTTM                string `json:"debtServiceCoverageRatioTTM"`
	InterestCoverageRatioTTM                   string `json:"interestCoverageRatioTTM"`
	ShortTermOperatingCashFlowCoverageRatioTTM string `json:"shortTermOperatingCashFlowCoverageRatioTTM"`
	OperatingCashFlowCoverageRatioTTM          string `json:"operatingCashFlowCoverageRatioTTM"`
	CapitalExpenditureCoverageRatioTTM         string `json:"capitalExpenditureCoverageRatioTTM"`
	DividendPaidAndCapexCoverageRatioTTM       string `json:"dividendPaidAndCapexCoverageRatioTTM"`
	DividendPayoutRatioTTM                     string `json:"dividendPayoutRatioTTM"`
	DividendYieldTTM                           string `json:"dividendYieldTTM"`
	EnterpriseValueTTM                         string `json:"enterpriseValueTTM"`
	RevenuePerShareTTM                         string `json:"revenuePerShareTTM"`
	NetIncomePerShareTTM                       string `json:"netIncomePerShareTTM"`
	InterestDebtPerShareTTM                    string `json:"interestDebtPerShareTTM"`
	CashPerShareTTM                            string `json:"cashPerShareTTM"`
	BookValuePerShareTTM                       string `json:"bookValuePerShareTTM"`
	TangibleBookValuePerShareTTM               string `json:"tangibleBookValuePerShareTTM"`
	ShareholdersEquityPerShareTTM              string `json:"shareholdersEquityPerShareTTM"`
	OperatingCashFlowPerShareTTM               string `json:"operatingCashFlowPerShareTTM"`
	CapexPerShareTTM                           string `json:"capexPerShareTTM"`
	FreeCashFlowPerShareTTM                    string `json:"freeCashFlowPerShareTTM"`
	NetIncomePerEBTTTM                         string `json:"netIncomePerEBTTTM"`
	EBTPerEbitTTM                              string `json:"ebtPerEbitTTM"`
	PriceToFairValueTTM                        string `json:"priceToFairValueTTM"`
	DebtToMarketCapTTM                         string `json:"debtToMarketCapTTM"`
	EffectiveTaxRateTTM                        string `json:"effectiveTaxRateTTM"`
	EnterpriseValueMultipleTTM                 string `json:"enterpriseValueMultipleTTM"`
	DividendPerShareTTM                        string `json:"dividendPerShareTTM"`
}

// GetRatiosTTMBulk retrieves trailing twelve months financial ratios for stocks
func (c *Client) GetRatiosTTMBulk() ([]RatiosTTMBulkResponse, error) {
	resp, err := c.doRequest("https://financialmodelingprep.com/stable/ratios-ttm-bulk", nil)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()



	// Parse the response
	var result []RatiosTTMBulkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
