package gorte

import (
	"net/url"
	"strings"

	"github.com/dhia-gharsallaoui/gorte/log"
	"github.com/hashicorp/go-retryablehttp"
)

const (
	defaultBaseURL = "https://digital.iservices.rte-france.com/"
)

type ClientConfig struct {
	Logger  log.Logger
	Key     string
	BaseURL string `default:"https://digital.iservices.rte-france.com/"`
}

type Client struct {
	logger      log.Logger
	client      *retryablehttp.Client
	baseURL     *url.URL
	config      ClientConfig
	token       *AuthToken
	Market      *market
	Consumption *consumption
	Partners    *partners
	Generation  *generation
	Exchanges   *exchanges
}

func NewClient(config ClientConfig) (*Client, error) {
	var logger log.Logger
	if config.Logger != nil {
		logger = config.Logger
	} else {
		logger = log.NewLogger(&log.LoggerConfiguration{
			Verbosity: log.Debug,
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
	c.client.Logger = log.NewHTTPLogger(logger)
	c.client.RetryMax = 10
	token, err := c.newToken()
	if err != nil {
		return nil, err
	}
	c.token = token
	c.Market = &market{client: &c}
	c.Consumption = &consumption{client: &c}
	c.Partners = &partners{client: &c}
	c.Generation = &generation{client: &c}
	c.Exchanges = &exchanges{client: &c}
	logger.Info("client was successfully created!")
	return &c, nil
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
