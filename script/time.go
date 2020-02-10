package script

import (
	"fmt"

	qlova "github.com/qlova/script"
)

type time struct {
	Ctx
}

//A Month specifies a month of the year (January = 1, ...).
type Month struct {
	Q Ctx
	Int
}

func (month Month) String() String {
	month.Q.Require(`var time_months = [ "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December" ];`)
	return month.Q.Value(`(time_months[%v])`, month.LanguageType().Raw()).String()
}

//Time is a script interface to a time value.
type Time struct {
	Q Ctx
	qlova.Native
}

//Now returns the current time.
func (t time) Now() Time {
	return Time{t.Ctx, t.Value("(new Date())").Native()}
}

func (t Time) String() String {
	return t.Q.Value(t.LanguageType().Raw() + ".toString()").String()
}

//Unix returns t as a Unix time, the number of seconds elapsed since January 1, 1970 UTC. The result does not depend on the location associated with t.
func (t Time) Unix() Int {
	return t.Q.Value("Math.floor(%v/1000)", t).Int()
}

//Day returns the day of the month specified by t.
func (t Time) Day() Int {
	return t.Q.Value(t.LanguageType().Raw() + ".getDate()").Int()
}

//Month returns the month of the year specified by t.
func (t Time) Month() Month {
	return Month{t.Q, t.Q.Value(t.LanguageType().Raw() + ".getMonth()").Int()}
}

//After executes a function after the given number of milliseconds have passed.
func (q Ctx) After(time float64, f func()) {
	q.Javascript("setTimeout(async function() {")
	f()
	q.Javascript("}, " + fmt.Sprint(time) + ");")
}

//Every executes a function every number of milliseconds.
func (q Ctx) Every(time float64, f func()) {
	q.Javascript("setInterval(async function() {")
	f()
	q.Javascript("}, " + fmt.Sprint(time) + ");")
}

//ToString turns a number into a string.
func (q Ctx) ToString(i Int) String {
	return q.Value(`(""+%v)`, i).String()
}
