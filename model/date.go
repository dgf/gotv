package model

import (
	"encoding/json"
	"time"
)

// time.Time alias with JSON date format
type Date time.Time

var format = "2006-01-02"

func (d Date) String() string {
	return time.Time(d).Format(format)
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func ToDate(s string) Date {
	t, _ := time.Parse(format, s)
	return Date(t)
}

func Today() Date {
	return Date(time.Now())
}
