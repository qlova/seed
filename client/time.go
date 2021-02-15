package client

import (
	"time"

	"qlova.org/seed/use/js"
)

//Date is a readonly client-typed date.
type Date js.AnyString

//NewDate returns a client-typed Date from the given time.
func NewDate(literal time.Time) String {
	return js.NewString(literal.Format("2006-01-02"))
}
