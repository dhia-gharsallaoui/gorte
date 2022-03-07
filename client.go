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
}

type Client struct {
	config ClientConfig
	token  AuthToken
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

	ctype := "application/json"

	auth := "Basic " + c.config.token

	client := &http.Client{}

	req, err := http.NewRequest("POST", c.config.authAdress, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", ctype)
	req.Header.Set("Authorization", auth)
	resp, err := client.Do(req)
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
