package console

import (
	"github.com/qlova/seed/js"
)

//Log logs the string to the console.
func Log(s js.AnyValue) js.Script {
	return func(q js.Ctx) {
		q.Run(js.Function{js.NewValue("console.log")}, s)
	}
}

//Error errors the string to the console.
func Error(s js.AnyValue) js.Script {
	return func(q js.Ctx) {
		q.Run(js.Function{js.NewValue("console.error")}, s)
	}
}
