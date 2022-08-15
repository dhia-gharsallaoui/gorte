package gorte

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

type AuthToken struct {
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	ExpiresIn   int       `json:"expires_in"`
	ExpiryDate  time.Time `json:"expiry_date"`
}

func (c *Client) authenticate(req *retryablehttp.Request) error {
	if time.Now().After(c.token.ExpiryDate) {
		var err error
		c.token, err = c.newToken()
		if err != nil {
			return err
		}
		c.logger.Info(fmt.Sprintf("New token generated! expires at %s", c.token.ExpiryDate.Format("2006-01-02 15:04:05")))
	}
	req.Header.Set("Authorization", c.token.TokenType+" "+c.token.AccessToken)
	return nil
}

func (c *Client) newToken() (*AuthToken, error) {
	authURL, err := URLGenerator(c.baseURL, "token/oauth/")
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", authURL.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+c.config.Key)
	resp, err := c.client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var token AuthToken
	if err := json.Unmarshal(body, &token); err != nil {
		return nil, err
	}
	token.ExpiryDate = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	c.logger.Debug(fmt.Sprintf("token was successfully generated! expires in %s", token.ExpiryDate.Format("2006-01-02 15:04:05")))
	return &token, nil
}
