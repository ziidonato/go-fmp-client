package go_fmp

import (
	"encoding/json"
	"fmt"
)

// RatingBulkResponse represents the response from the Rating Bulk API
type RatingBulkResponse struct {
	Symbol                  string `json:"symbol"`
	Date                    string `json:"date"`
	Rating                  string `json:"rating"`
	DiscountedCashFlowScore string `json:"discountedCashFlowScore"`
	ReturnOnEquityScore     string `json:"returnOnEquityScore"`
	ReturnOnAssetsScore     string `json:"returnOnAssetsScore"`
	DebtToEquityScore       string `json:"debtToEquityScore"`
	PriceToEarningsScore    string `json:"priceToEarningsScore"`
	PriceToBookScore        string `json:"priceToBookScore"`

// GetRatingBulk retrieves comprehensive rating data for multiple stocks
func (c *Client) GetRatingBulk() ([]RatingBulkResponse, error) {
	return doRequest[[]RatingBulkResponse](c, "https://financialmodelingprep.com/stable/rating-bulk", nil)
