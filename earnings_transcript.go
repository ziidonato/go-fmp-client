package fmp_client

// LatestEarningsTranscriptsParams represents parameters for latest earnings transcripts.
type LatestEarningsTranscriptsParams struct {
	// Limit is the maximum number of results (optional).
	Limit *int `json:"limit,omitempty"`
	// Page is the page number for pagination (optional).
	Page *int `json:"page,omitempty"`
}

// EarningsTranscriptParams represents parameters for specific earnings transcript.
type EarningsTranscriptParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
	// Year is the fiscal year (required).
	Year string `json:"year"`
	// Quarter is the fiscal quarter (required).
	Quarter string `json:"quarter"`
	// Limit is the maximum number of results (optional).
	Limit *int `json:"limit,omitempty"`
}

// TranscriptDatesBySymbolParams represents parameters for transcript dates.
type TranscriptDatesBySymbolParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
}

// LatestEarningsTranscriptResponse represents latest earnings transcript info.
type LatestEarningsTranscriptResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Period is the fiscal period (e.g., Q1, Q2, Q3, Q4).
	Period string `json:"period"`
	// FiscalYear is the fiscal year.
	FiscalYear int `json:"fiscalYear"`
	// Date is the date of the earnings call.
	Date string `json:"date"`
}

// EarningsTranscriptResponse represents a full earnings transcript.
type EarningsTranscriptResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Period is the fiscal period (e.g., Q1, Q2, Q3, Q4).
	Period string `json:"period"`
	// Year is the fiscal year.
	Year int `json:"year"`
	// Date is the date of the earnings call.
	Date string `json:"date"`
	// Content is the full transcript text.
	Content string `json:"content"`
}

// TranscriptDateResponse represents transcript availability by date.
type TranscriptDateResponse struct {
	// Quarter is the fiscal quarter number.
	Quarter int `json:"quarter"`
	// FiscalYear is the fiscal year.
	FiscalYear int `json:"fiscalYear"`
	// Date is the date of the earnings call.
	Date string `json:"date"`
}

// TranscriptSymbolListResponse represents companies with available transcripts.
type TranscriptSymbolListResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// CompanyName is the full company name.
	CompanyName string `json:"companyName"`
	// NoOfTranscripts is the number of available transcripts.
	NoOfTranscripts string `json:"noOfTranscripts"`
}

// LatestEarningsTranscripts retrieves the most recent earnings transcripts.
// Returns a list of companies with their latest earnings call information.
//
// Endpoint: /earning-call-transcript-latest
func (c *Client) LatestEarningsTranscripts(params LatestEarningsTranscriptsParams) ([]LatestEarningsTranscriptResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/earning-call-transcript-latest"

	var result []LatestEarningsTranscriptResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// EarningsTranscript retrieves the full text of an earnings call transcript.
// Returns the complete transcript for a specific quarter and year.
//
// Endpoint: /earning-call-transcript
func (c *Client) EarningsTranscript(params EarningsTranscriptParams) ([]EarningsTranscriptResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/earning-call-transcript"

	var result []EarningsTranscriptResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// TranscriptDatesBySymbol retrieves available transcript dates for a company.
// Returns all quarters and years with available earnings transcripts.
//
// Endpoint: /earning-call-transcript-dates
func (c *Client) TranscriptDatesBySymbol(params TranscriptDatesBySymbolParams) ([]TranscriptDateResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/earning-call-transcript-dates"

	var result []TranscriptDateResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// AvailableTranscriptSymbols retrieves companies with earnings transcripts.
// Returns symbols and transcript counts for all companies with available transcripts.
//
// Endpoint: /earnings-transcript-list
func (c *Client) AvailableTranscriptSymbols() ([]TranscriptSymbolListResponse, error) {
	pathName := "/earnings-transcript-list"

	var result []TranscriptSymbolListResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}