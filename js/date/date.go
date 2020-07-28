package date

import (
	"qlova.org/seed/client"
	"qlova.org/seed/js"
)

//Date is the JS Date type.
type Date struct {
	js.Value
}

//FromString creates a JavaScript Date instance that represents a single moment
//in time in a platform-independent format. Date objects contain a Number
//that represents milliseconds since 1 January 1970 UTC.
func FromString(dateString js.AnyString) Date {
	return Date{Value: js.New(js.NewValue("Date"), dateString)}
}

func Now() js.Number {
	return js.Number{js.NewValue(`Date.now()`)}
}

//MonthString returns the localised month string of the date, ie January, Febuary.
func (d Date) MonthString() js.String {
	return js.String{Value: d.Call("toLocaleString", client.NewString("default"), js.NewObject{
		"month": js.NewString("long"),
	})}
}

//TimeString returns the localised time-string of the date, ie 1:15:30 AM
func (d Date) TimeString() js.String {
	return js.String{Value: d.Call("toLocaleTimeString", client.NewString("default"))}
}

//Day returns the day of the month.
func (d Date) Day() js.Number {
	return js.Number{Value: d.Call("getDate")}
}
