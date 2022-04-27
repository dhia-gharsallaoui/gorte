package gorte

import (
	"net/http"
	"time"
)

type SignalEcowatt struct {
	Signals []struct {
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		Values    struct {
			Pdl []string `json:"PDL"`
			Hdf []string `json:"HDF"`
			Naq []string `json:"NAQ"`
			Bfc []string `json:"BFC"`
			Bre []string `json:"BRE"`
			Pac []string `json:"PAC"`
			Ara []string `json:"ARA"`
			Occ []string `json:"OCC"`
			Nor []string `json:"NOR"`
			Idf []string `json:"IDF"`
			Cvl []string `json:"CVL"`
			Ges []string `json:"GES"`
		} `json:"values"`
	} `json:"signals"`
}

func (co *consumption) GetSignalEcowatt(opt *Period) (*SignalEcowatt, *http.Response, error) {
	c := co.client
	req, err := c.NewRequest(http.MethodGet, "open_api/ecowatt/v3/signals", opt)
	if err != nil {
		c.logger.Err(err.Error())
		return nil, nil, err
	}
	var sig *SignalEcowatt
	resp, err := c.Do(req, &sig)
	if err != nil {
		c.logger.Err(err.Error())
		return nil, resp, err
	}
	return sig, resp, err
}
