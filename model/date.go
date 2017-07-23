package model

import (
	"encoding/json"
	"time"
)

// Date alias time.Time with JSON date format
type Date time.Time

var format = "2006-01-02"

func (d Date) String() string {
	return time.Time(d).Format(format)
}

// Add delegates to time.Time.AddDate
func (d Date) Add(years int, months int, days int) Date {
	return Date(time.Time(d).AddDate(years, months, days))
}

// After delegates to time.Time.After
func (d Date) After(o Date) bool {
	return time.Time(d).After(time.Time(o))
}

// Before delegates to time.Time.Before
func (d Date) Before(o Date) bool {
	return time.Time(d).Before(time.Time(o))
}

// MarshalJSON calls String
func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// ToDate from JSON string
func ToDate(s string) Date {
	t, _ := time.Parse(format, s)
	return Date(t)
}

// Today Date
func Today() Date {
	return Date(time.Now())
}
