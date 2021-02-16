package client

import (
	"time"

	"qlova.org/seed/use/js"
)

type date struct {
	js.Value
}

func (d date) GetTime() js.Value {
	return d.Value
}

type duration struct {
	js.Value
}

func (d duration) GetDuration() js.Value {
	return d.Value
}

//Time is a client time.Time representation.
type Time interface {
	Value

	GetTime() js.Value
}

//NewTime returns a new Time from the given Go time.
func NewTime(literal time.Time) Time {
	ecmaEpoch := literal.Sub(time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC))
	return date{NewFloat64(float64(ecmaEpoch.Milliseconds())).GetValue()}
}

//Now returns the current time on the client.
func Now() Time {
	return date{js.NewValue("Date.now()")}
}

//Duration in milliseconds.
type Duration interface {
	Value

	GetDuration() js.Value
}

//NewDuration returns a new Time from the given Go time.
func NewDuration(literal time.Duration) Duration {
	return duration{NewFloat64(float64(literal.Milliseconds())).GetValue()}
}
