package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/hashicorp/go-retryablehttp"
)

const (
	defaultBaseURL = "https://digital.iservices.rte-france.com/"
)

type AuthToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type ClientConfig struct {
	AuthAdress string `default:"https://digital.iservices.rte-france.com/token/oauth/"`
	Host       string `default:digital.iservices.rte-france.com`
	Token      string
	Method     string // "GET" "POST"
	ApiAdress  string
}

type Client struct {
	client  *retryablehttp.Client
	baseURL *url.URL
	config  ClientConfig
	token   AuthToken
}

func (c *Client) setBaseURL(urlStr string) error {
	// Make sure the given URL end with a slash
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	// Update the base URL of the client.
	c.baseURL = baseURL

	return nil
}

func (c *Client) NewClient(config ClientConfig) error {
	if c.baseURL == nil {
		c.setBaseURL(defaultBaseURL)
	}
	if config.Token == "" {
		return errors.New("Can't connect without the RTE token in Base 64 format. To get one subscribe to the API.")
	}

	if config.AuthAdress == "" {
		config.AuthAdress = "https://digital.iservices.rte-france.com/token/oauth/"
	}
	if config.Host == "" {
		config.Host = "digital.iservices.rte-france.com"
	}
	c.config = config

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10

	c.client = retryClient
	ctype := "application/json"

	auth := "Basic " + config.Token

	req, err := http.NewRequest("POST", config.AuthAdress, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", ctype)
	req.Header.Set("Authorization", auth)
	resp, err := c.client.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	var token AuthToken
	if err := json.Unmarshal(body, &token); err != nil {
		return err
	} else {
		log.Println("Client was successfully created !!!")
	}

	c.token = token
	return nil
}

func (c *Client) ConfigCheck() error {
	if c.config.Method == "" {
		return errors.New("No method provided")
	}
	if c.config.ApiAdress == "" {
		return errors.New("No API adress provided in the client config")
	}
	return nil

}
