package script

import (
	"fmt"

	"github.com/qlova/seed"
)

type data struct {
	seed.Data

	id string

	on map[string]Script

	requires map[string]string
}

var seeds = make(map[seed.Seed]data)

var unique int

//Unique returns a unique string suitable for variable names.
func Unique() string {
	unique++
	return fmt.Sprint("unique_", unique)
}
