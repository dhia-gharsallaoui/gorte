package main

import (
	"net/http"
	"time"
)

type TempoLikeCalendarsResp struct {
	TempoLikeCalendars struct {
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		Values    []struct {
			StartDate   time.Time `json:"start_date"`
			EndDate     time.Time `json:"end_date"`
			Value       string    `json:"value"`
			UpdatedDate time.Time `json:"updated_date"`
		} `json:"values"`
	} `json:"tempo_like_calendars"`
}

func (co *Consumption) GetTempoLikeCalendars(opt *Period) (*TempoLikeCalendarsResp, *http.Response, error) {
	c := co.client
	req, err := c.NewRequest(http.MethodGet, "open_api/tempo_like_supply_contract/v1/tempo_like_calendars", opt)
	if err != nil {
		return nil, nil, err
	}
	var sig *TempoLikeCalendarsResp
	resp, err := c.Do(req, &sig)
	if err != nil {
		return nil, resp, err
	}
	return sig, resp, err
}
