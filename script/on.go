package script

import (
	"fmt"
	"io/ioutil"

	"github.com/qlova/seed"
	"github.com/qlova/seed/js"
)

func On(event string, do Script) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		do(js.NewCtx(ioutil.Discard, c)) //Catch errors and harvest pages.

		var d data
		c.Read(&d)

		switch data := c.(type) {
		case Seed:
			c.Use()
			data.Q(fmt.Sprintf(`seed.on(%v, "%v", async function() {`, data.Element(), event))
			do(data.Q)
			data.Q(`});`)
		case Undo:
			//s.Root().Use()
			data.Q(fmt.Sprintf(`seed.on(%v, "%v", async function() {`, data.Element(), event))
			d.on[event](js.NewCtx(data.Q))
			data.Q(`});`)
		default:
			//s.Root().Use()
			if d.on == nil {
				d.on = make(map[string]Script)
			}
			d.on[event] = d.on[event].Append(do)
		}

		c.Write(d)
	})
}

func OnPress(do Script) seed.Option {
	return On("press", do)
}

func OnClick(do Script) seed.Option {
	return On("click", do)
}

func OnReady(do Script) seed.Option {
	return On("ready", do)
}

func OnInput(do Script) seed.Option {
	return On("input", do)
}

func OnEnter(do Script) seed.Option {
	return On("enter", do)
}

//OnError calls the provided script when there is an error not handled by this seed or any children seeds.
func OnError(do func(q Ctx, err Error)) seed.Option {
	return On("error", func(q Ctx) {
		do(q, Error{q, js.String{js.NewValue(`arguments[0]`)}, q.Number(1)})
	})
}
