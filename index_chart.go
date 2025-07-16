package go_fmp

import (
	"fmt"
)

type IndexChartParams struct {
	Symbol string
	From   string
	To     string
}

type Index1MinChartResponse struct {
	Date   string  `json:"date"`
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
}

type Index5MinChartResponse = Index1MinChartResponse

type Index1HourChartResponse = Index1MinChartResponse

func (c *Client) GetIndex1MinChart(symbol, from, to string) ([]Index1MinChartResponse, error) {
	return c.getIndexChartInterval(symbol, from, to, "1min")
}

func (c *Client) GetIndex5MinChart(symbol, from, to string) ([]Index5MinChartResponse, error) {
	return c.getIndexChartInterval(symbol, from, to, "5min")
}

func (c *Client) GetIndex1HourChart(symbol, from, to string) ([]Index1HourChartResponse, error) {
	return c.getIndexChartInterval(symbol, from, to, "1hour")
}

func (c *Client) getIndexChartInterval(symbol, from, to, interval string) ([]Index1MinChartResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}
	params := map[string]string{
		"symbol": symbol,
	}
	if from != "" {
		params["from"] = from
	}
	if to != "" {
		params["to"] = to
	}
	url := c.BaseURL + "/historical-chart/" + interval
	var result []Index1MinChartResponse
	err := c.doRequest(url, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
