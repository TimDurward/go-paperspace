package paperspace

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	clientVersion  = "0.1.0"
	defaultBaseURL = "https://api.paperspace.io/"
	userAgent      = "paperspace/" + clientVersion
	mediaType      = "application/json"

	headerRateLimit     = "RateLimit-Limit"
	headerRateRemaining = "RateLimit-Remaining"
	headerRateReset     = "RateLimit-Reset"
)

// Config manages Authorization with Paperspace API.
type Config struct {
	// API Key for client authorization.
	APIKey string
}

// Client manages communication with Paperspace API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// UserAgent for client.
	UserAgent string

	// Config for client.
	Config *Config

	// Services used for talking to different parts of the Paperspace API.
	Scripts  *ScriptsService
	Machines *MachinesService
}

// Response is a Paperspace response. This wraps the standard http.Response returned from Paperspace.
type Response struct {
	*http.Response
}

// ErrorResponse reports the error caused by an API request.
type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	// Error message
	Message string `json:"message"`
}

// addOptions adds the parameters in opt as URL query parameters to s. opt
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// NewClient returns a new Paperspace API client.
func NewClient(config *Config, httpClient *http.Client) (*Client, error) {
	// If client is not provided (nil) pass in default HTTP client.
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	// Parse defaultBaseURL into URL structure.
	parsedBaseURL, err := url.Parse(defaultBaseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		client:    http.DefaultClient,
		BaseURL:   parsedBaseURL,
		UserAgent: userAgent,
		Config:    config,
	}
	c.Scripts = &ScriptsService{client: c}
	c.Machines = &MachinesService{client: c}

	return c, nil
}

// NewRequest creates a new Paperspace API request.
// A relative URL can be provided in urlStr, in which case it is resolved relative to the baseURL of the Client.
// If specified, the value pointed to by body is JSON encoded and included as the request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	rel, err := url.Parse(urlStr)
	if err != nil {
		fmt.Print("errorr")
		return nil, err
	}
	// Relative URLs should be specified without a preceding slash since baseURL will have the trailing slash
	rel.Path = strings.TrimLeft(rel.Path, "/")

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	// Set request Headers
	req.Header.Add("x-api-key", c.Config.APIKey)
	req.Header.Add("User-Agent", c.UserAgent)
	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", mediaType)

	return req, nil
}

// Do sends an API request and returns the API response.
// The API response is JSON decoded and stored in the value pointed to by v, or returned as an error if an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	httpResp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer httpResp.Body.Close()

	resp := newResponse(httpResp)
	err = CheckResponse(httpResp)
	if err != nil {
		return nil, err
	}
	if v != nil {
		// Open a NewDecoder and defer closing the reader only if there is a provided interface to decode to
		err = json.NewDecoder(httpResp.Body).Decode(v)
		if err != nil {
			return nil, err
		}
	}
	return resp, err
}

// CheckResponse checks the Paperpspace API response for a non-2xx status code.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	err := fmt.Errorf("Request failed. Please analyze the request body for more details. Status code: %d", r.StatusCode)
	return err
}

// newResponse creates a new Paperspace API Response for http.Response
func newResponse(r *http.Response) *Response {
	response := Response{Response: r}
	return &response
}
