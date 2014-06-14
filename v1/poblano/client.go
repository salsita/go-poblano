// Copyright (c) 2014 The go-poblano AUTHORS
//
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package poblano

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	LibraryVersion = "0.0.1"

	defaultUserAgent = "go-poblano/" + LibraryVersion
)

type Client struct {
	// Poblano access token to be used to authenticate API requests.
	token string

	// Optional Basic auth credentials to use when connecting to Poblano.
	credentials *Credentials

	// HTTP client to be used for communication with the Poblano API.
	client *http.Client

	// Base URL of the Poblano API that is to be used to form API requests.
	baseURL *url.URL

	// User-Agent header to use when connecting to the Poblano API.
	UserAgent string

	// GitHub service encapsulates all the functionality connected to GitHub.
	GitHub *GitHubService
}

type Credentials struct {
	Username string
	Password string
}

func NewClient(baseURL, apiToken string, cred *Credentials) (*Client, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		token:       apiToken,
		credentials: cred,
		client:      http.DefaultClient,
		baseURL:     base,
		UserAgent:   defaultUserAgent,
	}
	c.GitHub = newGitHubService(c)

	return c, nil
}

func (c *Client) NewRequest(method, urlPath string, body interface{}) (*http.Request, error) {
	relativePath, err := url.Parse(urlPath)
	if err != nil {
		return nil, err
	}

	u := c.baseURL.ResolveReference(relativePath)

	buf := new(bytes.Buffer)
	if body != nil {
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if cred := c.credentials; cred != nil {
		req.SetBasicAuth(cred.Username, cred.Password)
	}

	req.Header.Add("User-Agent", c.UserAgent)
	req.Header.Add("X-PoblanoToken", c.token)
	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return resp, &ErrHTTP{resp}
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, err
}
