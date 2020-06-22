package script

import (
	"fmt"
	"time"

	"qlova.org/seed"
	"qlova.org/seed/js"
	"qlova.org/seed/js/console"
)

type Data struct {
	seed.Data

	id string

	On map[string]Script
}

var seeds = make(map[seed.Seed]Data)

var unique int

//Unique returns a unique string suitable for variable names.
func Unique() string {
	unique++
	return fmt.Sprint("unique_", unique)
}

//Download downloads a requested resource with given name and path.
func Download(name, path string) Script {
	return func(q Ctx) {
		q.Run(js.Function{js.NewValue("await download")}, q.String(name), q.String(path))
	}
}

//New creates a new script out of multiple scripts.
func New(scripts ...Script) Script {
	return func(q Ctx) {
		for _, s := range scripts {
			if s != nil {
				q(s)
			}
		}
	}
}

func After(duration time.Duration, do Script) Script {
	return js.Global().Run("setTimeout", js.NewFunction(do), js.NewNumber(duration.Seconds()*1000))
}

func Trace(s Script) Script {
	return New(
		console.Error(js.NewString("trace")),
		s,
	)
}
