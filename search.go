package fmp_client

type StockSymbolSearchParams struct {
	Query  string `json:"query"`
	Limit  *int   `json:"limit,omitempty"`
	Number *int   `json:"number,omitempty"`
}

type StockSymbolSearchResponse struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	ExchangeFullName string `json:"exchangeFullName"`
	Exchange         string `json:"exchange"`
}

func (c *Client) StockSymbolSearch(params StockSymbolSearchParams) ([]StockSymbolSearchResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/search-symbol"

	var result []StockSymbolSearchResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

type CompnayNameSearchParams struct {
	Query    string  `json:"query"`
	Limit    *int    `json:"limit,omitempty"`
	Number   *int    `json:"number,omitempty"`
	Exchange *string `json:"exchange,omitempty"`
}

type CompnayNameSearchResponse struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	ExchangeFullName string `json:"exchangeFullName"`
	Exchange         string `json:"exchange"`
}

func (c *Client) CompanyNameSearch(params CompnayNameSearchParams) ([]CompnayNameSearchResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/search-name"

	var result []CompnayNameSearchResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

type CIKParams struct {
	CIK string `json:"cik"`
}

type CIKResponse struct {
	CIK string `json:"cik"`
}

func (c *Client) CIK(params CIKParams) ([]CIKResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/search-cik"

	var result []CIKResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

type CUSIPParams struct {
	CUSIP string `json:"cusip"`
}

type CUSIPResponse struct {
	Symbol      string `json:"AAPL"`
	CUSIP       string `json:"cusip"`
	CompanyName string `json:"companyName"`
	MarketCap   int    `json:"marketCap"`
}

func (c *Client) CUSIP(params CUSIPParams) ([]CUSIPResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/search-cusip"

	var result []CUSIPResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

type ISINParams struct {
	ISIN string `json:"isin"`
}

type ISINResponse struct {
	Symbol    string `json:"symbol"`
	Name      string `json:"companyName"`
	ISIN      string `json:"isin"`
	MarketCap int    `json:"marketCap"`
}

func (c *Client) ISIN(params ISINParams) ([]ISINResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/search-isin"

	var result []ISINResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}
