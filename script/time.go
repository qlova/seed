package script

import (
	"fmt"

	qlova "github.com/qlova/script"
)

type time struct {
	Ctx
}

//Time is a script interface to a time value.
type Time struct {
	Q Ctx
	qlova.Native
}

//Now returns the current time.
func (q time) Now() Time {
	return Time{q.Ctx, q.Value("(new Date())").Native()}
}

func (time Time) String() String {
	return time.Q.Value(time.LanguageType().Raw() + ".toString()").String()
}

//After executes a function after the given number of milliseconds have passed.
func (q Ctx) After(time float64, f func()) {
	q.Javascript("setTimeout(function() {")
	f()
	q.Javascript("}, " + fmt.Sprint(time) + ");")
}

//Every executes a function every number of milliseconds.
func (q Ctx) Every(time float64, f func()) {
	q.Javascript("setInterval(function() {")
	f()
	q.Javascript("}, " + fmt.Sprint(time) + ");")
}
