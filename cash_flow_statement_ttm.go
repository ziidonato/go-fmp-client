package go_fmp

import (
	"fmt"
	"time"
)

// CashFlowStatementTTMParams represents the parameters for the Cash Flow Statement TTM API
type CashFlowStatementTTMParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Limit  *int   `json:"limit"`  // Optional: Number of results (Maximum 1000 records per request)
}

// CashFlowStatementTTMResponse represents the response from the Cash Flow Statement TTM API
type CashFlowStatementTTMResponse struct {
	Date                                   time.Time `json:"date"`
	Symbol                                 string `json:"symbol"`
	ReportedCurrency                       string `json:"reportedCurrency"`
	CIK                                    string `json:"cik"`
	FilingDate                             time.Time `json:"filingDate"`
	AcceptedDate                           time.Time `json:"acceptedDate"`
	CalendarYear                           string `json:"calendarYear"`
	Period                                 string `json:"period"`
	NetIncomeTTM                           float64 `json:"netIncomeTTM"`
	DepreciationAndAmortizationTTM         float64 `json:"depreciationAndAmortizationTTM"`
	DeferredIncomeTaxTTM                   float64 `json:"deferredIncomeTaxTTM"`
	StockBasedCompensationTTM              float64 `json:"stockBasedCompensationTTM"`
	ChangeInWorkingCapitalTTM              float64 `json:"changeInWorkingCapitalTTM"`
	AccountsReceivablesTTM                 float64 `json:"accountsReceivablesTTM"`
	InventoryTTM                           float64 `json:"inventoryTTM"`
	AccountsPayablesTTM                    float64 `json:"accountsPayablesTTM"`
	OtherWorkingCapitalTTM                 float64 `json:"otherWorkingCapitalTTM"`
	OtherNonCashItemsTTM                   float64 `json:"otherNonCashItemsTTM"`
	NetCashProvidedByOperatingActivitiesTTM float64 `json:"netCashProvidedByOperatingActivitiesTTM"`
	InvestmentsInPropertyPlantAndEquipmentTTM float64 `json:"investmentsInPropertyPlantAndEquipmentTTM"`
	AcquisitionsNetTTM                     float64 `json:"acquisitionsNetTTM"`
	PurchasesOfInvestmentsTTM              float64 `json:"purchasesOfInvestmentsTTM"`
	SalesMaturitiesOfInvestmentsTTM        float64 `json:"salesMaturitiesOfInvestmentsTTM"`
	OtherInvestingActivitesTTM             float64 `json:"otherInvestingActivitesTTM"`
	NetCashUsedForInvestingActivitesTTM    float64 `json:"netCashUsedForInvestingActivitesTTM"`
	DebtRepaymentTTM                       float64 `json:"debtRepaymentTTM"`
	CommonStockIssuedTTM                   float64 `json:"commonStockIssuedTTM"`
	CommonStockRepurchasedTTM              float64 `json:"commonStockRepurchasedTTM"`
	DividendsPaidTTM                       float64 `json:"dividendsPaidTTM"`
	OtherFinancingActivitesTTM             float64 `json:"otherFinancingActivitesTTM"`
	NetCashUsedProvidedByFinancingActivitiesTTM float64 `json:"netCashUsedProvidedByFinancingActivitiesTTM"`
	EffectOfForexChangesOnCashTTM          float64 `json:"effectOfForexChangesOnCashTTM"`
	NetChangeInCashTTM                     float64 `json:"netChangeInCashTTM"`
	CashAtEndOfPeriodTTM                   float64 `json:"cashAtEndOfPeriodTTM"`
	CashAtBeginningOfPeriodTTM             float64 `json:"cashAtBeginningOfPeriodTTM"`
	OperatingCashFlowTTM                   float64 `json:"operatingCashFlowTTM"`
	CapitalExpenditureTTM                  float64 `json:"capitalExpenditureTTM"`
	FreeCashFlowTTM                        float64 `json:"freeCashFlowTTM"`
	Link                                   string `json:"link"`
	FinalLink                              string `json:"finalLink"`
}

// CashFlowStatementTTM retrieves trailing twelve months cash flow statement data for a specific stock symbol
func (c *Client) CashFlowStatementTTM(params CashFlowStatementTTMParams) ([]CashFlowStatementTTMResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	if params.Limit != nil {
		if *params.Limit > 1000 {
			return nil, fmt.Errorf("limit cannot exceed 1000")
		}
		urlParams["limit"] = fmt.Sprintf("%d", *params.Limit)
	}

	var result []CashFlowStatementTTMResponse
	err := c.doRequest(c.BaseURL+"/cash-flow-statement-ttm", urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
