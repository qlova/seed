//Package page provides global pages that can be swapped in and out.
package page

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/state"
)

func ID(p Page) string {
	return strings.Replace(reflect.TypeOf(p).String(), ".", "_", -1)
}

type Page interface {
	Page(Seed)
}

type Seed struct {
	seed.Seed
}

func Scope(c seed.Seed) Seed {
	return Seed{c}
}

type data struct {
	seed.Data

	pages []Page
}

var seeds = make(map[seed.Seed]data)

func (c Seed) Goto(page Page) script.Script {
	var data data
	c.Read(&data)
	data.pages = append(data.pages, page)
	c.Write(data)

	return func(q script.Ctx) {
		fmt.Fprintf(q, `seed.goto("%v");`, ID(page))
	}
}

func OnEnter(f script.Script) seed.Option {
	return script.On("pageenter", f)
}

func OnExit(f script.Script) seed.Option {
	return script.On("pageexit", f)
}

func State(p Page) state.State {
	return state.New(state.SetKey("page."+ID(p)), state.ReadOnly())
}
