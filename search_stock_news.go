package go_fmp

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// SearchStockNewsResponse represents the response from the search stock news API
type SearchStockNewsResponse struct {
	Symbol        string `json:"symbol"`
	PublishedDate string `json:"publishedDate"`
	Publisher     string `json:"publisher"`
	Title         string `json:"title"`
	Image         string `json:"image"`
	Site          string `json:"site"`
	Text          string `json:"text"`
	URL           string `json:"url"`

// SearchStockNews searches for stock-related news by symbol
func (c *Client) SearchStockNews(symbols string, page, limit int, from, to string) ([]SearchStockNewsResponse, error) {
	if symbols == "" {
		return nil, fmt.Errorf("symbols is required")
	}
	if page < 0 {
		return nil, fmt.Errorf("page must be non-negative")
	}
	if limit <= 0 || limit > 250 {
		return nil, fmt.Errorf("limit must be between 1 and 250")
	}
	if page > 100 {
		return nil, fmt.Errorf("page cannot exceed 100")
	}

	params := map[string]string{
		"symbols": symbols,
		"page":    strconv.Itoa(page),
		"limit":   strconv.Itoa(limit),
	}

	if from != "" {
		params["from"] = from
	}
	if to != "" {
		params["to"] = to
	}

	url := "https://financialmodelingprep.com/stable/news/stock"

	return doRequest[[]SearchStockNewsResponse](c, url, params)
