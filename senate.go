package fmp_client

// CongressTradingLatestParams represents parameters for latest congress trading endpoints.
type CongressTradingLatestParams struct {
	// Page is the page number for pagination (optional).
	Page *int `json:"page,omitempty"`
	// Limit is the maximum number of results per page (optional).
	Limit *int `json:"limit,omitempty"`
}

// CongressTradingBySymbolParams represents parameters for trading by symbol.
type CongressTradingBySymbolParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
}

// CongressTradingByNameParams represents parameters for trading by name.
type CongressTradingByNameParams struct {
	// Name is the congress member's name to search for (required).
	Name string `json:"name"`
}

// CongressTradingResponse represents a congressional trading disclosure.
type CongressTradingResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// DisclosureDate is when the disclosure was filed.
	DisclosureDate string `json:"disclosureDate"`
	// TransactionDate is when the transaction occurred.
	TransactionDate string `json:"transactionDate"`
	// FirstName is the congress member's first name.
	FirstName string `json:"firstName"`
	// LastName is the congress member's last name.
	LastName string `json:"lastName"`
	// Office is the office or position held.
	Office string `json:"office"`
	// District is the state/district represented.
	District string `json:"district"`
	// Owner is who owns the asset (Self, Spouse, etc).
	Owner string `json:"owner"`
	// AssetDescription is the description of the asset.
	AssetDescription string `json:"assetDescription"`
	// AssetType is the type of asset (Stock, Bond, etc).
	AssetType string `json:"assetType"`
	// Type is the transaction type (Purchase, Sale, etc).
	Type string `json:"type"`
	// Amount is the transaction amount range.
	Amount string `json:"amount"`
	// CapitalGainsOver200USD indicates if capital gains exceed $200.
	CapitalGainsOver200USD string `json:"capitalGainsOver200USD"`
	// Comment is any additional comment on the transaction.
	Comment string `json:"comment"`
	// Link is the URL to the disclosure document.
	Link string `json:"link"`
}

// LatestSenateTrading retrieves the most recent Senate trading disclosures.
// Returns financial transactions by U.S. Senators and their families.
//
// Endpoint: /senate-latest
func (c *Client) LatestSenateTrading(params CongressTradingLatestParams) ([]CongressTradingResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/senate-latest"

	var result []CongressTradingResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// LatestHouseTrading retrieves the most recent House trading disclosures.
// Returns financial transactions by U.S. House members and their families.
//
// Endpoint: /house-latest
func (c *Client) LatestHouseTrading(params CongressTradingLatestParams) ([]CongressTradingResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/house-latest"

	var result []CongressTradingResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// SenateTradingBySymbol retrieves Senate trading activity for a specific stock.
// Returns all Senate transactions involving the specified symbol.
//
// Endpoint: /senate-trades
func (c *Client) SenateTradingBySymbol(params CongressTradingBySymbolParams) ([]CongressTradingResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/senate-trades"

	var result []CongressTradingResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// SenateTradingByName retrieves trading activity by Senator name.
// Returns all transactions by Senators matching the name search.
//
// Endpoint: /senate-trades-by-name
func (c *Client) SenateTradingByName(params CongressTradingByNameParams) ([]CongressTradingResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/senate-trades-by-name"

	var result []CongressTradingResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// HouseTradingBySymbol retrieves House trading activity for a specific stock.
// Returns all House member transactions involving the specified symbol.
//
// Endpoint: /house-trades
func (c *Client) HouseTradingBySymbol(params CongressTradingBySymbolParams) ([]CongressTradingResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/house-trades"

	var result []CongressTradingResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// HouseTradingByName retrieves trading activity by House member name.
// Returns all transactions by House members matching the name search.
//
// Endpoint: /house-trades-by-name
func (c *Client) HouseTradingByName(params CongressTradingByNameParams) ([]CongressTradingResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/house-trades-by-name"

	var result []CongressTradingResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}