package go_fmp

import (
	"fmt"
	"time"
)

// IncomeStatementTTMParams represents the parameters for the Income Statement TTM API
type IncomeStatementTTMParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Limit  *int   `json:"limit"`  // Optional: Number of results (Maximum 1000 records per request)
}

// IncomeStatementTTMResponse represents the response from the Income Statement TTM API
type IncomeStatementTTMResponse struct {
	Symbol                                  string    `json:"symbol"`
	Date                                    time.Time `json:"date"`
	ReportedCurrency                        string    `json:"reportedCurrency"`
	CIK                                     string    `json:"cik"`
	FilingDate                              time.Time `json:"filingDate"`
	AcceptedDate                            time.Time `json:"acceptedDate"`
	CalendarYear                            string    `json:"calendarYear"`
	Period                                  string    `json:"period"`
	RevenueTTM                              float64   `json:"revenueTTM"`
	CostOfRevenueTTM                        float64   `json:"costOfRevenueTTM"`
	GrossProfitTTM                          float64   `json:"grossProfitTTM"`
	GrossProfitRatioTTM                     float64   `json:"grossProfitRatioTTM"`
	ResearchAndDevelopmentExpensesTTM       float64   `json:"researchAndDevelopmentExpensesTTM"`
	GeneralAndAdministrativeExpensesTTM     float64   `json:"generalAndAdministrativeExpensesTTM"`
	SellingAndMarketingExpensesTTM          float64   `json:"sellingAndMarketingExpensesTTM"`
	SellingGeneralAndAdministrativeExpensesTTM float64 `json:"sellingGeneralAndAdministrativeExpensesTTM"`
	OtherExpensesTTM                        float64   `json:"otherExpensesTTM"`
	OperatingExpensesTTM                    float64   `json:"operatingExpensesTTM"`
	CostAndExpensesTTM                      float64   `json:"costAndExpensesTTM"`
	InterestIncomeTTM                       float64   `json:"interestIncomeTTM"`
	InterestExpenseTTM                      float64   `json:"interestExpenseTTM"`
	DepreciationAndAmortizationTTM          float64   `json:"depreciationAndAmortizationTTM"`
	EBITDATTM                               float64   `json:"ebitdaTTM"`
	EBITDARatioTTM                          float64   `json:"ebitdaratioTTM"`
	OperatingIncomeTTM                      float64   `json:"operatingIncomeTTM"`
	OperatingIncomeRatioTTM                 float64   `json:"operatingIncomeRatioTTM"`
	TotalOtherIncomeExpensesNetTTM          float64   `json:"totalOtherIncomeExpensesNetTTM"`
	IncomeBeforeTaxTTM                      float64   `json:"incomeBeforeTaxTTM"`
	IncomeBeforeTaxRatioTTM                 float64   `json:"incomeBeforeTaxRatioTTM"`
	IncomeTaxExpenseTTM                     float64   `json:"incomeTaxExpenseTTM"`
	NetIncomeTTM                            float64   `json:"netIncomeTTM"`
	NetIncomeRatioTTM                       float64   `json:"netIncomeRatioTTM"`
	EPSTTM                                  float64   `json:"epsTTM"`
	EPSDilutedTTM                           float64   `json:"epsdilutedTTM"`
	WeightedAverageShsOutTTM                float64   `json:"weightedAverageShsOutTTM"`
	WeightedAverageShsOutDilTTM             float64   `json:"weightedAverageShsOutDilTTM"`
	Link                                    string    `json:"link"`
	FinalLink                               string    `json:"finalLink"`
}

// IncomeStatementTTM retrieves trailing twelve months income statement data for a specific stock symbol
func (c *Client) IncomeStatementTTM(params IncomeStatementTTMParams) ([]IncomeStatementTTMResponse, error) {
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

	var result []IncomeStatementTTMResponse
	err := c.doRequest(c.BaseURL+"/income-statement-ttm", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
