package gorte

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
	Key     string
	BaseURL string `default:"https://digital.iservices.rte-france.com/"`
}

type Client struct {
	client      *retryablehttp.Client
	baseURL     *url.URL
	config      ClientConfig
	token       AuthToken
	Market      *market
	Consumption *consumption
	Partners    *partners
	Generation  *generation
	Exchanges   *exchanges
}

func setURL(urlStr string) (*url.URL, error) {
	// Make sure the given URL end with a slash
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	URL, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	return URL, nil
}

func NewClient(config ClientConfig) (*Client, error) {
	if config.Key == "" {
		return nil, errors.New("can't connect without the RTE token in Base 64 format. to get one subscribe to the API")
	}
	c := Client{}
	var err error
	if config.BaseURL == "" {
		config.BaseURL = defaultBaseURL
	}
	c.baseURL, err = setURL(config.BaseURL)
	if err != nil {
		return nil, err
	}
	c.config = config
	c.client = retryablehttp.NewClient()
	c.client.RetryMax = 10
	authURL, err := URLGenerator(c.baseURL, "token/oauth/")
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", authURL.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+config.Key)
	resp, err := c.client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	var token AuthToken
	if err := json.Unmarshal(body, &token); err != nil {
		return nil, err
	} else {
		log.Println("Client was successfully created !!!")
	}
	c.token = token
	c.Market = &market{client: &c}
	c.Consumption = &consumption{client: &c}
	c.Partners = &partners{client: &c}
	c.Generation = &generation{client: &c}
	c.Exchanges = &exchanges{client: &c}

	return &c, nil
}
