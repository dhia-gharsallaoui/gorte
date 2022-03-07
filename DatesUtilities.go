package main

import (
	"errors"
	"time"
)

type Horizon struct {
	startDate time.Time
	endDate   time.Time
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
