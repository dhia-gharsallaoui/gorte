package gorte

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/hashicorp/go-retryablehttp"
)

func URLGenerator(ur *url.URL, path string) (*url.URL, error) {
	u := *ur
	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}

	// Set the encoded path data
	u.RawPath = u.Path + path
	u.Path = u.Path + unescaped
	return &u, nil
}

func (c *Client) NewRequest(method, path string, opt interface{}) (*retryablehttp.Request, error) {
	u, err := URLGenerator(c.baseURL, path)
	if err != nil {
		return nil, err
	}
	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")
	auth := c.token.TokenType + " " + c.token.AccessToken
	reqHeaders.Set("Authorization", auth)
	var body interface{}
	switch {
	case method == http.MethodPost || method == http.MethodPut:
		reqHeaders.Set("Content-Type", "application/json")

		if opt != nil {
			body, err = json.Marshal(opt)
			if err != nil {
				return nil, err
			}
		}
	case opt != nil:
		q, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		u.RawQuery = q.Encode()
	}
	req, err := retryablehttp.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	for k, v := range reqHeaders {
		req.Header[k] = v
	}
	return req, nil
}

func (c *Client) Do(req *retryablehttp.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return resp, err
}
