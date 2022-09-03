package gorte

import (
	"net/http"
	"time"
)

type SignalEcowatt struct {
	Signals []struct {
		GenerationFichier time.Time `json:"GenerationFichier"`
		Jour              time.Time `json:"jour"`
		Dvalue            int       `json:"dvalue"`
		Message           string    `json:"message"`
		Values            []struct {
			Pas    int `json:"pas"`
			Hvalue int `json:"hvalue"`
		} `json:"values"`
	} `json:"signals"`
}

func (co *consumption) GetSignalEcowatt() (*SignalEcowatt, *http.Response, error) {
	c := co.client
	req, err := c.NewRequest(http.MethodGet, "open_api/ecowatt/v4/signals", nil)
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
