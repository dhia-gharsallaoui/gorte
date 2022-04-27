package gorte

import (
	"encoding/json"
	"io/ioutil"
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
	Logger  Logger
	Key     string
	BaseURL string `default:"https://digital.iservices.rte-france.com/"`
}

type Client struct {
	logger      Logger
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
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	url, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func NewClient(config ClientConfig) (*Client, error) {
	var logger Logger
	if config.Logger != nil {
		logger = config.Logger
	} else {
		logger = NewLogger(&LoggerConfiguration{
			Verbosity: Warn,
		})
	}
	if config.Key == "" {
		logger.Fatal("can't connect without the RTE token in Base 64 format. to get one subscribe to the API")
	}
	c := Client{logger: logger}
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
	c.client.Logger = newLeveledLogger(logger)
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
		logger.Info("client was successfully created!")
	}
	c.token = token
	c.Market = &market{client: &c}
	c.Consumption = &consumption{client: &c}
	c.Partners = &partners{client: &c}
	c.Generation = &generation{client: &c}
	c.Exchanges = &exchanges{client: &c}
	return &c, nil
}
