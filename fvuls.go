package fvuls

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

const (
	defaultBaseURL    = "https://rest.vuls.biz/"
	defaultUserAgent  = "futurevuls-sdk-go"
	apiRequestTimeout = 30 * time.Second
)

// Client api client for fvuls
type Client struct {
	BaseURL           *url.URL
	APIKey            string
	Verbose           bool
	UserAgent         string
	AdditionalHeaders http.Header
	HTTPClient        *http.Client
}

// NewClient returns new fvuls.Client
func NewClient(apikey string) *Client {
	u, _ := url.Parse(defaultBaseURL)
	return &Client{u, apikey, false, defaultUserAgent, http.Header{
		"Authorization": []string{apikey},
	}, &http.Client{}}
}

// NewClientWithOptions returns new fvuls.Client
func NewClientWithOptions(apikey string, rawurl string, verbose bool) (*Client, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	return &Client{u, apikey, verbose, defaultUserAgent, http.Header{
		"Authorization": []string{apikey},
	}, &http.Client{}}, nil
}

func (c *Client) urlFor(path string) *url.URL {
	newURL, err := url.Parse(c.BaseURL.String())
	if err != nil {
		panic("invalid url passed")
	}
	newURL.Path = path
	return newURL
}

func (c *Client) buildReq(req *http.Request) *http.Request {
	for header, values := range c.AdditionalHeaders {
		for _, v := range values {
			req.Header.Add(header, v)
		}
	}
	req.Header.Set("Authorization", c.APIKey)
	req.Header.Set("User-Agent", c.UserAgent)
	return req
}

// Request request to fvuls and receive response
func (c *Client) Request(req *http.Request) (resp *http.Response, err error) {
	req = c.buildReq(req)

	if c.Verbose {
		dump, err := httputil.DumpRequest(req, true)
		if err == nil {
			log.Printf("%s", dump)
		}
	}

	client := c.HTTPClient
	client.Timeout = apiRequestTimeout
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	if c.Verbose {
		dump, err := httputil.DumpResponse(resp, true)
		if err == nil {
			log.Printf("%s", dump)
		}
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		message := extractErrorMessage(resp.Body)
		defer resp.Body.Close()
		if message != "" {
			return nil, &APIError{StatusCode: resp.StatusCode, Message: message}
		}
		return nil, &APIError{StatusCode: resp.StatusCode, Message: resp.Status}
	}
	return resp, nil
}

// PostJSON shortcut method for posting json
func (c *Client) PostJSON(path string, payload interface{}) (*http.Response, error) {
	return c.requestJSON("POST", path, payload)
}

// PutJSON shortcut method for putting json
func (c *Client) PutJSON(path string, payload interface{}) (*http.Response, error) {
	return c.requestJSON("PUT", path, payload)

}

func (c *Client) toJSON(src interface{}) string {
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(src)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(body.String())
}

func (c *Client) requestJSON(method string, path string, payload interface{}) (*http.Response, error) {
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, c.urlFor(path).String(), &body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	return c.Request(req)
}

func closeResponse(resp *http.Response) {
	if resp != nil {
		resp.Body.Close()
	}
}
