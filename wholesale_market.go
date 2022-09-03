package gorte

import (
	"net/http"
	"time"
)

type FrancePowerExchangesValue struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Value     float64   `json:"value"`
	Price     float64   `json:"price"`
}

type FrancePowerExchanges struct {
	Data []struct {
		StartDate   time.Time                   `json:"start_date"`
		EndDate     time.Time                   `json:"end_date"`
		UpdatedDate time.Time                   `json:"updated_date"`
		Values      []FrancePowerExchangesValue `json:"values"`
	} `json:"france_power_exchanges"`
}

func (s *market) GetFrancePowerExchanges() (*FrancePowerExchanges, *http.Response, error) {
	c := s.client
	req, err := c.NewRequest(http.MethodGet, "open_api/wholesale_market/v2/france_power_exchanges", nil)
	if err != nil {
		c.logger.Err(err.Error())
		return nil, nil, err
	}
	fpe := &FrancePowerExchanges{}
	resp, err := c.Do(req, fpe)
	if err != nil {
		c.logger.Err(err.Error())
		return nil, resp, err
	}
	return fpe, resp, err
}
