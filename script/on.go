package script

import (
	"github.com/qlova/seed"
)

func On(event string, do Script) seed.Option {
	ToJavascript(do) //Catch errors and harvest pages.

	return seed.NewOption(func(c seed.Seed) {
		var d data
		c.Read(&d)

		switch q := c.(type) {
		case Seed:
			c.Use()
			q.Javascript(`seed.on(%v, "%v", async function() {`, q.Element(), event)
			do(q.Ctx)
			q.Javascript(`});`)
		case Undo:
			//s.Root().Use()
			q.Javascript(`seed.on(%v, "%v", async function() {`, q.Element(), event)
			q.Ctx.Write(ToJavascript(d.on[event]))
			q.Javascript(`});`)
		default:
			//s.Root().Use()
			if d.on == nil {
				d.on = make(map[string]Script)
			}
			d.on[event] = d.on[event].Then(do)
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

//OnError calls the provided script when there is an error not handled by this seed or any children seeds.
func OnError(do func(q Ctx, err Error)) seed.Option {
	return On("error", func(q Ctx) {
		do(q, Error{q, q.Value(`arguments[0]`).String(), q.Int(1)})
	})
}
