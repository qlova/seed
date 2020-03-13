//Package page provides global pages that can be swapped in and out.
package page

import (
	"fmt"

	"github.com/qlova/seed"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/script"
)

type Page interface {
	Page(Scope)
}

type Scope struct {
	seed.Seed
	harvester
}

func (scope Scope) Goto(page Page) script.Script {
	return func(q script.Ctx) {
		fmt.Fprintf(q, `seed.goto("%v");`, html.ID(scope.harvester.harvest(page)))
	}
}

func OnEnter(f script.Script) seed.Option {
	return script.On("pageenter", f)
}

func OnExit(f script.Script) seed.Option {
	return script.On("pageexit", f)
}
