package script

import (
	"fmt"

	qlova "github.com/qlova/script"
)

type time struct {
	Script
}

//Time is a script interface to a time value.
type Time struct {
	Q Script
	qlova.Native
}

//Now returns the current time.
func (q time) Now() Time {
	return Time{q.Script, q.Value("(new Date())").Native()}
}

func (time Time) String() String {
	return time.Q.Value(time.LanguageType().Raw() + ".toString()").String()
}

//After executes a function after the given number of milliseconds have passed.
func (q Script) After(time float64, f func()) {
	q.Javascript("setTimeout(function() {")
	f()
	q.Javascript("}, " + fmt.Sprint(time) + ");")
}

//Every executes a function every number of milliseconds.
func (q Script) Every(time float64, f func()) {
	q.Javascript("setInterval(function() {")
	f()
	q.Javascript("}, " + fmt.Sprint(time) + ");")
}
