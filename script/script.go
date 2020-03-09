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
	on map[string]Script

	requires map[string]string
}

var seeds = make(map[seed.Seed]data)

func On(event string, do Script) seed.Option {
	ToJavascript(do) //Catch errors and harvest pages.

	return seed.NewOption(func(s seed.Any) {
		s.Root().Use()
		data := seeds[s.Root()]
		if data.on == nil {
			data.on = make(map[string]Script)
			seeds[s.Root()] = data
		}
		data.on[event] = data.on[event].Then(do)
	}, func(s seed.Ctx) {
		s.Root().Use()
		fmt.Fprintf(s.Ctx, `%v.on%v = async function() {`, s.Element(), event)
		do(Ctx{s.Ctx})
		fmt.Fprint(s.Ctx, `};`)
	}, func(s seed.Ctx) {
		s.Root().Use()
		data := seeds[s.Root()]
		fmt.Fprintf(s.Ctx, `%v.on%v = async function() {`, s.Element(), event)
		s.Ctx.Write(ToJavascript(data.on[event]))
		fmt.Fprint(s.Ctx, `};`)
	})
}

func OnClick(do Script) seed.Option {
	return On("click", do)
}

func OnReady(do Script) seed.Option {
	return On("ready", do)
}

var unique int

//Unique returns a unique string suitable for variable names.
func Unique() string {
	unique++
	return fmt.Sprint("unique_", unique)
}
