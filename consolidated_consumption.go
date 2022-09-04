package gorte

import (
	"net/http"
	"time"

	"github.com/dhia-gharsallaoui/gorte/utils"
)

type ConsolidatedPowerConsumption struct {
	ConsolidatedPowerConsumptions []struct {
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

func (co *consumption) GetConsolidatedPowerConsumption(opt utils.Period) (*ConsolidatedPowerConsumption, *http.Response, error) {
	c := co.client
	req, err := c.NewRequest(http.MethodGet, "open_api/consolidated_consumption/v1/consolidated_power_consumption", opt)
	if err != nil {
		c.logger.Err(err.Error())
		return nil, nil, err
	}
	sig := &ConsolidatedPowerConsumption{}
	resp, err := c.Do(req, sig)
	if err != nil {
		c.logger.Err(err.Error())
		return nil, resp, err
	}
	return sig, resp, err
}

type ConsolidatedEnergyConsumption struct {
	ConsolidatedEnergyConsumptions []struct {
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

func (co *consumption) GetConsolidatedEnergyConsumption(opt utils.Period) (*ConsolidatedEnergyConsumption, *http.Response, error) {
	c := co.client
	req, err := c.NewRequest(http.MethodGet, "open_api/consolidated_consumption/v1/consolidated_energy_consumption", opt)
	if err != nil {
		c.logger.Err(err.Error())
		return nil, nil, err
	}
	sig := &ConsolidatedEnergyConsumption{}
	resp, err := c.Do(req, sig)
	if err != nil {
		c.logger.Err(err.Error())
		return nil, resp, err
	}
	return sig, resp, err
}
