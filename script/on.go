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
			//s.Root().Use()
			if event == "press" {
				q.Javascript(`seed.op(%v, async function() {`, q.Element())
			} else {
				q.Javascript(`%v.on%v = async function() {`, q.Element(), event)
			}
			do(q.Ctx)
			if event == "press" {
				q.Javascript(`});`)
			} else {
				q.Javascript(`};`)
			}
		case Undo:
			//s.Root().Use()
			if event == "press" {
				q.Javascript(`seed.op(%v, async function() {`, q.Element())
			} else {
				q.Javascript(`%v.on%v = async function() {`, q.Element(), event)
			}
			q.Ctx.Write(ToJavascript(d.on[event]))
			if event == "press" {
				q.Javascript(`});`)
			} else {
				q.Javascript(`};`)
			}
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
