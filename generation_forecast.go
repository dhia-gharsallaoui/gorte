package gorte

import (
	"net/http"
	"time"

	"github.com/dhia-gharsallaoui/gorte/utils"
)

type ForecastsResp struct {
	Forecasts []struct {
		StartDate      time.Time `json:"start_date"`
		EndDate        time.Time `json:"end_date"`
		Type           string    `json:"type"`
		ProductionType string    `json:"production_type"`
		Values         []struct {
			StartDate   time.Time `json:"start_date"`
			EndDate     time.Time `json:"end_date"`
			UpdatedDate time.Time `json:"updated_date"`
			Value       int       `json:"value"`
		} `json:"values"`
		SubType string `json:"sub_type,omitempty"`
	} `json:"forecasts"`
}

type GetForecastsOptions struct {
	StartDate utils.Time `url:"start_date"`
	EndDate   utils.Time `url:"end_date"`
	ProdType  string     `url:"production_type"`
	Type      string     `url:"type"`
}

func (s *generation) GetForecasts(opt *GetForecastsOptions) (*ForecastsResp, *http.Response, error) {
	c := s.client
	req, err := c.NewRequest(http.MethodGet, "open_api/generation_forecast/v2/forecasts", opt)
	if err != nil {
		c.logger.Err(err.Error())
		return nil, nil, err
	}
	sig := &ForecastsResp{}
	resp, err := c.Do(req, sig)
	if err != nil {
		c.logger.Err(err.Error())
		return nil, resp, err
	}
	return sig, resp, err
}
