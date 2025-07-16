package go_fmp

import (
	"encoding/json"
	"fmt"
)

// InstitutionalOwnershipExtractParams represents the parameters for the Institutional Ownership Extract API
type InstitutionalOwnershipExtractParams struct {
	CIK     string `json:"cik"`     // Required: CIK number (e.g., "0001388838")
	Year    string `json:"year"`    // Required: Year (e.g., "2023")
	Quarter string `json:"quarter"` // Required: Quarter (e.g., "3")
}

// InstitutionalOwnershipExtractResponse represents the response from the Institutional Ownership Extract API
type InstitutionalOwnershipExtractResponse struct {
	Date          string `json:"date"`
	FilingDate    string `json:"filingDate"`
	AcceptedDate  string `json:"acceptedDate"`
	CIK           string `json:"cik"`
	SecurityCusip string `json:"securityCusip"`
	Symbol        string `json:"symbol"`
	NameOfIssuer  string `json:"nameOfIssuer"`
	Shares        int64  `json:"shares"`
	TitleOfClass  string `json:"titleOfClass"`
	SharesType    string `json:"sharesType"`
	PutCallShare  string `json:"putCallShare"`
	Value         int64  `json:"value"`
	Link          string `json:"link"`
	FinalLink     string `json:"finalLink"`
}

// InstitutionalOwnershipExtract retrieves detailed data from SEC filings for institutional ownership
func (c *Client) InstitutionalOwnershipExtract(params InstitutionalOwnershipExtractParams) ([]InstitutionalOwnershipExtractResponse, error) {
	if params.CIK == "" {
		return nil, fmt.Errorf("cik parameter is required")
	}

	if params.Year == "" {
		return nil, fmt.Errorf("year parameter is required")
	}

	if params.Quarter == "" {
		return nil, fmt.Errorf("quarter parameter is required")
	}

	urlParams := map[string]string{
		"cik":     params.CIK,
		"year":    params.Year,
		"quarter": params.Quarter,
	}

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/institutional-ownership/extract", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []InstitutionalOwnershipExtractResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
