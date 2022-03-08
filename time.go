package main

import (
	"net/url"
	"time"
)

type Time time.Time

func NewTime(t time.Time) Time {
	return Time(t)
}

func (t Time) EncodeValues(key string, v *url.Values) error {
	layout := "2006-01-02T15:04:05-07:00"
	v.Set(key, time.Time(t).Format(layout))
	return nil
}
