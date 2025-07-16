package go_fmp

import (
	"encoding/json"
	"fmt"
)

// InstitutionalOwnershipHolderPerformanceSummaryParams represents the parameters for the Institutional Ownership Holder Performance Summary API
type InstitutionalOwnershipHolderPerformanceSummaryParams struct {
	CIK  string `json:"cik"`  // Required: CIK number (e.g., "0001067983")
	Page *int   `json:"page"` // Optional: Page number (default: 0)
}

// InstitutionalOwnershipHolderPerformanceSummaryResponse represents the response from the Institutional Ownership Holder Performance Summary API
type InstitutionalOwnershipHolderPerformanceSummaryResponse struct {
	Date                                               string  `json:"date"`
	CIK                                                string  `json:"cik"`
	InvestorName                                       string  `json:"investorName"`
	PortfolioSize                                      int     `json:"portfolioSize"`
	SecuritiesAdded                                    int     `json:"securitiesAdded"`
	SecuritiesRemoved                                  int     `json:"securitiesRemoved"`
	MarketValue                                        int64   `json:"marketValue"`
	PreviousMarketValue                                int64   `json:"previousMarketValue"`
	ChangeInMarketValue                                int64   `json:"changeInMarketValue"`
	ChangeInMarketValuePercentage                      float64 `json:"changeInMarketValuePercentage"`
	AverageHoldingPeriod                               int     `json:"averageHoldingPeriod"`
	AverageHoldingPeriodTop10                          int     `json:"averageHoldingPeriodTop10"`
	AverageHoldingPeriodTop20                          int     `json:"averageHoldingPeriodTop20"`
	Turnover                                           float64 `json:"turnover"`
	TurnoverAlternateSell                              float64 `json:"turnoverAlternateSell"`
	TurnoverAlternateBuy                               float64 `json:"turnoverAlternateBuy"`
	Performance                                        int64   `json:"performance"`
	PerformancePercentage                              float64 `json:"performancePercentage"`
	LastPerformance                                    int64   `json:"lastPerformance"`
	ChangeInPerformance                                int64   `json:"changeInPerformance"`
	Performance1year                                   int64   `json:"performance1year"`
	PerformancePercentage1year                         float64 `json:"performancePercentage1year"`
	Performance3year                                   int64   `json:"performance3year"`
	PerformancePercentage3year                         float64 `json:"performancePercentage3year"`
	Performance5year                                   int64   `json:"performance5year"`
	PerformancePercentage5year                         float64 `json:"performancePercentage5year"`
	PerformanceSinceInception                          int64   `json:"performanceSinceInception"`
	PerformanceSinceInceptionPercentage                float64 `json:"performanceSinceInceptionPercentage"`
	PerformanceRelativeToSP500Percentage               float64 `json:"performanceRelativeToSP500Percentage"`
	Performance1yearRelativeToSP500Percentage          float64 `json:"performance1yearRelativeToSP500Percentage"`
	Performance3yearRelativeToSP500Percentage          float64 `json:"performance3yearRelativeToSP500Percentage"`
	Performance5yearRelativeToSP500Percentage          float64 `json:"performance5yearRelativeToSP500Percentage"`
	PerformanceSinceInceptionRelativeToSP500Percentage float64 `json:"performanceSinceInceptionRelativeToSP500Percentage"`
}

// InstitutionalOwnershipHolderPerformanceSummary retrieves performance summary for institutional holders
func (c *Client) InstitutionalOwnershipHolderPerformanceSummary(params InstitutionalOwnershipHolderPerformanceSummaryParams) ([]InstitutionalOwnershipHolderPerformanceSummaryResponse, error) {
	if params.CIK == "" {
		return nil, fmt.Errorf("cik parameter is required")
	}

	urlParams := map[string]string{
		"cik": params.CIK,
	}

	if params.Page != nil {
		urlParams["page"] = fmt.Sprintf("%d", *params.Page)
	}

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/institutional-ownership/holder-performance-summary", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []InstitutionalOwnershipHolderPerformanceSummaryResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
