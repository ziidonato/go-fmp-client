package go_fmp

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// FMPArticlesResponse represents the response from the FMP articles API
type FMPArticlesResponse struct {
	Title   string `json:"title"`
	Date    string `json:"date"`
	Content string `json:"content"`
	Tickers string `json:"tickers"`
	Image   string `json:"image"`
	Link    string `json:"link"`
	Author  string `json:"author"`
	Site    string `json:"site"`
}

// GetFMPArticles retrieves the latest articles from Financial Modeling Prep
func (c *Client) GetFMPArticles(page, limit int) ([]FMPArticlesResponse, error) {
	if page < 0 {
		return nil, fmt.Errorf("page must be non-negative")
	}
	if limit <= 0 {
		return nil, fmt.Errorf("limit must be positive")
	}

	url := "https://financialmodelingprep.com/stable/fmp-articles"

	resp, err := c.doRequest(url, map[string]string{
		"page":  strconv.Itoa(page),
		"limit": strconv.Itoa(limit),
	})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()


	var result []FMPArticlesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
