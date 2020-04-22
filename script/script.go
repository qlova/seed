package script

import (
	"fmt"

	"github.com/qlova/seed"
)

type data struct {
	seed.Data

	id string

	on map[string]Script
}

var seeds = make(map[seed.Seed]data)

var unique int

//Unique returns a unique string suitable for variable names.
func Unique() string {
	unique++
	return fmt.Sprint("unique_", unique)
}

//Download downloads a requested resource with given name and path.
func Download(name, path string) Script {
	return func(q Ctx) {
		q.Run("await seed.download", q.String(name), q.String(path))
	}
}
