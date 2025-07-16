package go_fmp

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// GeneralNewsResponse represents the response from the general news API
type GeneralNewsResponse struct {
	Symbol        *string `json:"symbol"`
	PublishedDate string  `json:"publishedDate"`
	Publisher     string  `json:"publisher"`
	Title         string  `json:"title"`
	Image         string  `json:"image"`
	Site          string  `json:"site"`
	Text          string  `json:"text"`
	URL           string  `json:"url"`

// GeneralNews retrieves the latest general news articles from various sources
func (c *Client) GeneralNews(page, limit int, from, to string) ([]GeneralNewsResponse, error) {
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
		"page":  strconv.Itoa(page),
		"limit": strconv.Itoa(limit),
	}

	if from != "" {
		params["from"] = from
	}
	if to != "" {
		params["to"] = to
	}

	url := "https://financialmodelingprep.com/stable/news/general-latest"

	return doRequest[[]GeneralNewsResponse](c, url, params)
