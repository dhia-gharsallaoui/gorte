package gorte

import (
	"net/http"
	"time"

	"github.com/dhia-gharsallaoui/gorte/utils"
)

type SignalValue struct {
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Value       bool      `json:"value"`
	UpdatedDate time.Time `json:"updated_date"`
}

type Signal []struct {
	StartDate time.Time     `json:"start_date"`
	EndDate   time.Time     `json:"end_date"`
	Type      string        `json:"type"`
	Values    []SignalValue `json:"values"`
}

type Signals struct {
	Signal `json:"signals"`
}

func (s *market) GetSignals(opt utils.Period) (*Signals, *http.Response, error) {
	c := s.client
	req, err := c.NewRequest(http.MethodGet, "open_api/signal/v1/signals", opt)
	if err != nil {
		c.logger.Err(err.Error())
		return nil, nil, err
	}
	sig := &Signals{}
	resp, err := c.Do(req, sig)
	if err != nil {
		c.logger.Err(err.Error())
		return nil, resp, err
	}
	return sig, resp, err
}
