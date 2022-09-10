package gorte

import (
	"net/http"
	"time"

	"github.com/dhia-gharsallaoui/gorte/utils"
)

type ConsolidatedPowerConsumptionResp struct {
	ConsolidatedPowerConsumption []struct {
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		Values    []struct {
			StartDate   time.Time `json:"start_date"`
			EndDate     time.Time `json:"end_date"`
			Value       int       `json:"value"`
			UpdatedDate string    `json:"updated_date"`
			Status      string    `json:"status"`
		} `json:"values"`
	} `json:"consolidated_power_consumption"`
}

func (co *consumption) GetConsolidatedPowerConsumption(opt utils.Period) (*ConsolidatedPowerConsumptionResp, *http.Response, error) {
	c := co.client
	req, err := c.NewRequest(http.MethodGet, "open_api/consolidated_consumption/v1/consolidated_power_consumption", opt)
	if err != nil {
		c.logger.Err(err.Error())
		return nil, nil, err
	}
	sig := &ConsolidatedPowerConsumptionResp{}
	resp, err := c.Do(req, sig)
	if err != nil {
		c.logger.Err(err.Error())
		return nil, resp, err
	}
	return sig, resp, err
}

type ConsolidatedEnergyConsumptionResp struct {
	ConsolidatedEnergyConsumption []struct {
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		Values    []struct {
			StartDate   time.Time `json:"start_date"`
			EndDate     time.Time `json:"end_date"`
			Value       int       `json:"value"`
			UpdatedDate string    `json:"updated_date"`
			Status      string    `json:"status"`
		} `json:"values"`
	} `json:"consolidated_energy_consumption"`
}

func (co *consumption) GetConsolidatedEnergyConsumption(opt utils.Period) (*ConsolidatedEnergyConsumptionResp, *http.Response, error) {
	c := co.client
	req, err := c.NewRequest(http.MethodGet, "open_api/consolidated_consumption/v1/consolidated_energy_consumption", opt)
	if err != nil {
		c.logger.Err(err.Error())
		return nil, nil, err
	}
	sig := &ConsolidatedEnergyConsumptionResp{}
	resp, err := c.Do(req, sig)
	if err != nil {
		c.logger.Err(err.Error())
		return nil, resp, err
	}
	return sig, resp, err
}
