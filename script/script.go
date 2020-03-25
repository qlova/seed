package script

import (
	"fmt"

	"github.com/qlova/script"
	"github.com/qlova/seed"
)

type Script func(Ctx)

func (s Script) Then(other Script) Script {
	return func(q Ctx) {
		if s != nil {
			s(q)
		}
		if other != nil {
			other(q)
		}
	}
}

type Ctx struct {
	script.Ctx
}

func CtxFrom(ctx script.AnyCtx) Ctx {
	return Ctx{ctx.RootCtx()}
}

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
