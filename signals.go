package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Signal struct {
	Signals []struct {
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		Type      string    `json:"type"`
		Values    []struct {
			StartDate   time.Time `json:"start_date"`
			EndDate     time.Time `json:"end_date"`
			Value       bool      `json:"value"`
			UpdatedDate time.Time `json:"updated_date"`
		} `json:"values"`
	} `json:"signals"`
}

type Horizon struct {
	startDate time.Time
	endDate   time.Time
}
type GridPeak struct {
	StartDate  time.Time
	EndDate    time.Time
	Type       string // PP1/PP2 or  PP2 or NORMAL
	UpdateDate time.Time
}

func DateFormat(date time.Time) string {
	layout := "2006-01-02T15:04:05-07:00"
	z := date.Format(layout)
	return z
}

func datecheck(date interface{}) (*string, error) {

	switch t := interface{}(date).(type) {
	case time.Time:

		x := t.String()
		return &x, nil
	case string:
		return &t, nil
	default:
		return nil, errors.New("Date must be in string or time.time type.")
	}
}

func dateParser(startDate, endDate interface{}) (*Horizon, error) {

	StartDate, err := datecheck(startDate)
	if err != nil {
		return nil, err
	}
	EndDate, err := datecheck(endDate)
	if err != nil {
		return nil, err
	}

	start, err := time.Parse("2006-01-02 15:04", *StartDate)
	if err != nil {
		return nil, err
	}

	end, err := time.Parse("2006-01-02 15:04", *EndDate)
	if err != nil {
		return nil, err

	}
	return &Horizon{start, end}, nil
}

func (c *Client) GetSignals(startDate, endDate interface{}) (*Signal, error) {
	err := c.ConfigCheck()
	if err != nil {
		return nil, err
	}

	dd, err := dateParser(startDate, endDate)
	if err != nil {
		return nil, err
	}
	dates := *dd
	req, err := http.NewRequest(c.config.method, c.config.apiAdress, nil)
	if err != nil {
		return nil, err
	}
	auth := c.token.TokenType + " " + c.token.AccessToken
	req.Header.Set("Authorization", auth)
	q := req.URL.Query()
	q.Add("start_date", DateFormat(dates.startDate))
	q.Add("end_date", DateFormat(dates.endDate))
	req.URL.RawQuery = q.Encode()
	resp, err := c.client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var respSignal Signal
	if err := json.Unmarshal(body, &respSignal); err != nil {
		return nil, err
	} else {
		log.Println("Signal catched !")
	}
	return &respSignal, nil

}

func (signal *Signal) ToGrid(onlyPP bool) []GridPeak {

	var stype string
	var peak GridPeak
	gridpeak := make([]GridPeak, 0)
	for i, ppsignal := range signal.Signals[0].Values {
		if signal.Signals[0].Values[i].Value && signal.Signals[1].Values[i].Value {
			stype = "PP1"
		} else if signal.Signals[1].Values[i].Value {
			stype = "PP2"
		} else {
			stype = "NORMAL"
		}
		if onlyPP == true {
			if stype != "NORMAL" {
				peak = GridPeak{ppsignal.StartDate, ppsignal.EndDate, stype, ppsignal.UpdatedDate}
				gridpeak = append(gridpeak, peak)

			}

		} else {
			peak = GridPeak{ppsignal.StartDate, ppsignal.EndDate, stype, ppsignal.UpdatedDate}
			gridpeak = append(gridpeak, peak)
		}
	}
	return gridpeak
}
