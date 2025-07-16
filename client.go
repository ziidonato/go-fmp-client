package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HTTPClient is an interface for making HTTP requests.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is the main struct for interacting with the Financial Modeling Prep API.
type Client struct {
	HTTPClient HTTPClient
	APIKey     string
	BaseURL    string
}

// NewClient creates a new Client. If httpClient is nil, http.DefaultClient is used.
func NewClient(httpClient HTTPClient, apiKey string, baseURL string) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	if baseURL == "" {
		baseURL = "https://financialmodelingprep.com/stable"
	}
	newClient := &Client{
		HTTPClient: httpClient,
		APIKey:     apiKey,
		BaseURL:    baseURL,
	}

	err := testClient(newClient)
	if err != nil {
		return nil, fmt.Errorf("failed to test client: %v", err)
	}

	return newClient, nil
}

func testClient(client *Client) error {
	// doRequest expects a result interface, but for a test we can use a dummy variable
	var body map[string]any
	if err := client.doRequest("https://financialmodelingprep.com/api/v3/quote/AAPL", map[string]string{"apikey": client.APIKey}, &body); err != nil {
		return fmt.Errorf("failed to perform test request: %v", err)
	}

	if body["error"] != nil {
		return fmt.Errorf("failed to perform test request: %v", body["error"])
	}

	return nil
}

func editQueryParams(req *http.Request, params map[string]string) {
	query := req.URL.Query()
	for key, value := range params {
		query.Set(key, value)
	}
	req.URL.RawQuery = query.Encode()
}

// Remove the get method, as doRequest replaces it.

// doRequest is a helper that makes a request, checks for errors, and decodes the response into the provided result interface.
func (c *Client) doRequest(url string, params map[string]string, result any) error {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	params["apikey"] = c.APIKey
	editQueryParams(req, params)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	return nil
}
