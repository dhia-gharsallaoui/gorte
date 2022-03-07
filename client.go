package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type AuthToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type ClientConfig struct {
	authAdress string `default:"https://digital.iservices.rte-france.com/token/oauth/"`
	host       string `default:digital.iservices.rte-france.com`
	token      string
	method     string // "GET" "POST"
	apiAdress  string
}

type Client struct {
	config ClientConfig
	token  AuthToken
	client *http.Client
}

func (c *Client) NewClient() error {
	if c.config.token == "" {
		return errors.New("Can't connect without the RTE token in Base 64 format. To get one subscribe to the API.")
	}

	if c.config.authAdress == "" {
		c.config.authAdress = "https://digital.iservices.rte-france.com/token/oauth/"
	}
	if c.config.host == "" {
		c.config.host = "digital.iservices.rte-france.com"
	}

	c.client = &http.Client{}
	ctype := "application/json"

	auth := "Basic " + c.config.token

	req, err := http.NewRequest("POST", c.config.authAdress, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", ctype)
	req.Header.Set("Authorization", auth)
	resp, err := c.client.Do(req)
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
	if c.config.method == "" {
		return errors.New("No method provided")
	}
	if c.config.apiAdress == "" {
		return errors.New("No API adress provided in the client config")
	}
	return nil

}
