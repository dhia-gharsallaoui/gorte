package gorte

import (
	"net/http"
	"time"

	"github.com/dhia-gharsallaoui/gorte/utils"
)

type AnnualForecastsResp struct {
	AnnualForecasts []struct {
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		Values    []struct {
			StartDate                   time.Time `json:"start_date"`
			EndDate                     time.Time `json:"end_date"`
			AverageLoadSaturdayToFriday int       `json:"average_load_saturday_to_friday"`
			AverageLoadMondayToSunday   int       `json:"average_load_monday_to_sunday"`
			WeeklyMinimum               int       `json:"weekly_minimum"`
			WeeklyMaximum               int       `json:"weekly_maximum"`
			AverageLoadUpdatedDate      string    `json:"average_load_updated_date"`
			MarginUpdatedDate           string    `json:"margin_updated_date"`
			ForecastMargin              int       `json:"forecast_margin"`
		} `json:"values"`
	} `json:"annual_forecasts"`
}

func (co *consumption) GetAnnualForecasts(opt *utils.Period) (*AnnualForecastsResp, *http.Response, error) {
	c := co.client
	req, err := c.NewRequest(http.MethodGet, "open_api/consumption/v1/annual_forecasts", opt)
	if err != nil {
		c.logger.Err(err.Error())
		return nil, nil, err
	}
	var sig *AnnualForecastsResp
	resp, err := c.Do(req, &sig)
	if err != nil {
		c.logger.Err(err.Error())
		return nil, resp, err
	}
	return sig, resp, err
}

type WeeklyForecastsResp struct {
	WeeklyForecasts []struct {
		EndDate time.Time `json:"end_date"`
		Values  []struct {
			EndDate   time.Time `json:"end_date"`
			Value     int       `json:"value"`
			StartDate time.Time `json:"start_date"`
		} `json:"values"`
		UpdatedDate time.Time `json:"updated_date"`
		Peak        struct {
			TemperatureDeviation float64   `json:"temperature_deviation"`
			Value                int       `json:"value"`
			PeakHour             time.Time `json:"peak_hour"`
			Temperature          float64   `json:"temperature"`
		} `json:"peak"`
		StartDate time.Time `json:"start_date"`
	} `json:"weekly_forecasts"`
}

func (co *consumption) GetWeeklyForecasts(opt *utils.Period) (*WeeklyForecastsResp, *http.Response, error) {
	c := co.client
	req, err := c.NewRequest(http.MethodGet, "open_api/consumption/v1/weekly_forecasts", opt)
	if err != nil {
		c.logger.Err(err.Error())
		return nil, nil, err
	}
	var sig *WeeklyForecastsResp
	resp, err := c.Do(req, &sig)
	if err != nil {
		c.logger.Err(err.Error())
		return nil, resp, err
	}
	return sig, resp, err
}

type ShortTermResp struct {
	ShortTerm []struct {
		Type      string    `json:"type"`
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		Values    []struct {
			StartDate   time.Time `json:"start_date"`
			EndDate     time.Time `json:"end_date"`
			UpdatedDate time.Time `json:"updated_date"`
			Value       int       `json:"value"`
		} `json:"values"`
	} `json:"short_term"`
}

func (co *consumption) GetShortTerm(opt *utils.Period) (*ShortTermResp, *http.Response, error) {
	c := co.client
	req, err := c.NewRequest(http.MethodGet, "open_api/consumption/v1/short_term", opt)
	if err != nil {
		c.logger.Err(err.Error())
		return nil, nil, err
	}
	var sig *ShortTermResp
	resp, err := c.Do(req, &sig)
	if err != nil {
		c.logger.Err(err.Error())
		return nil, resp, err
	}
	return sig, resp, err
}
