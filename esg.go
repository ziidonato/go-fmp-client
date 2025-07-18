package fmp_client

// ESGDisclosuresParams represents parameters for ESG disclosures endpoint.
type ESGDisclosuresParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
}

// ESGRatingsParams represents parameters for ESG ratings endpoint.
type ESGRatingsParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
}

// ESGBenchmarkParams represents parameters for ESG benchmark endpoint.
type ESGBenchmarkParams struct {
	// Year is the fiscal year for benchmark data (optional).
	Year *string `json:"year,omitempty"`
}

// ESGDisclosureResponse represents ESG disclosure data.
type ESGDisclosureResponse struct {
	// Date is the date of the disclosure.
	Date string `json:"date"`
	// AcceptedDate is when the filing was accepted.
	AcceptedDate string `json:"acceptedDate"`
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// CIK is the Central Index Key.
	CIK string `json:"cik"`
	// CompanyName is the full company name.
	CompanyName string `json:"companyName"`
	// FormType is the SEC form type.
	FormType string `json:"formType"`
	// EnvironmentalScore is the environmental performance score.
	EnvironmentalScore float64 `json:"environmentalScore"`
	// SocialScore is the social responsibility score.
	SocialScore float64 `json:"socialScore"`
	// GovernanceScore is the corporate governance score.
	GovernanceScore float64 `json:"governanceScore"`
	// ESGScore is the overall ESG score.
	ESGScore float64 `json:"ESGScore"`
	// URL is the link to the SEC filing.
	URL string `json:"url"`
}

// ESGRatingResponse represents ESG rating information.
type ESGRatingResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// CIK is the Central Index Key.
	CIK string `json:"cik"`
	// CompanyName is the full company name.
	CompanyName string `json:"companyName"`
	// Industry is the industry classification.
	Industry string `json:"industry"`
	// FiscalYear is the fiscal year of the rating.
	FiscalYear int `json:"fiscalYear"`
	// ESGRiskRating is the letter grade rating.
	ESGRiskRating string `json:"ESGRiskRating"`
	// IndustryRank is the ranking within the industry.
	IndustryRank string `json:"industryRank"`
}

// ESGBenchmarkResponse represents ESG benchmark data by sector.
type ESGBenchmarkResponse struct {
	// FiscalYear is the fiscal year of the benchmark.
	FiscalYear int `json:"fiscalYear"`
	// Sector is the business sector.
	Sector string `json:"sector"`
	// EnvironmentalScore is the average environmental score for the sector.
	EnvironmentalScore float64 `json:"environmentalScore"`
	// SocialScore is the average social score for the sector.
	SocialScore float64 `json:"socialScore"`
	// GovernanceScore is the average governance score for the sector.
	GovernanceScore float64 `json:"governanceScore"`
	// ESGScore is the average overall ESG score for the sector.
	ESGScore float64 `json:"ESGScore"`
}

// ESGDisclosures retrieves ESG disclosure data for a specific company.
// Returns environmental, social, and governance scores from SEC filings.
//
// Endpoint: /esg-disclosures
func (c *Client) ESGDisclosures(params ESGDisclosuresParams) ([]ESGDisclosureResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/esg-disclosures"

	var result []ESGDisclosureResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// ESGRatings retrieves ESG ratings and risk assessments for a company.
// Returns letter grades and industry rankings for ESG performance.
//
// Endpoint: /esg-ratings
func (c *Client) ESGRatings(params ESGRatingsParams) ([]ESGRatingResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/esg-ratings"

	var result []ESGRatingResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// ESGBenchmark retrieves ESG benchmark scores by sector.
// Returns average ESG scores for comparing companies within industries.
//
// Endpoint: /esg-benchmark
func (c *Client) ESGBenchmark(params ESGBenchmarkParams) ([]ESGBenchmarkResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/esg-benchmark"

	var result []ESGBenchmarkResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}