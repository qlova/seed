package seed

import (
	"fmt"
	"strconv"

	"github.com/qlova/seed/script"
	"github.com/qlova/seed/script/global"
)

//SetTextf allows setting the text of a seed with dynamic arguments.
func (seed Seed) SetTextf(format string, args ...global.Variable) {

	var converted = make([]interface{}, len(args))
	for i, arg := range args {
		converted[i] = `"+(window.localStorage.getItem("` + arg.Ref() + `") || "")+"`
	}

	var s = fmt.Sprintf(strconv.Quote(format), converted...)

	for _, arg := range args {
		seed.OnGlobalChanged(arg, func(q script.Ctx) {
			seed.Ctx(q).SetText(q.Value(s).String())
		})
	}
}

type dynamic struct {
	Handlers map[string][]func(q script.Ctx)
}

//OnGlobalChanged is called whenever the global reference is changed.
func (seed Seed) OnGlobalChanged(v global.Variable, f func(script.Ctx)) {
	if seed.dynamic.Handlers == nil {
		seed.dynamic.Handlers = make(map[string][]func(q script.Ctx))
	}
	seed.dynamic.Handlers[v.Ref()] = append(seed.dynamic.Handlers[v.Ref()], f)
}

//SetDynamicText sets the text content of the seed which will be dynamic at runtime.
func (seed Seed) SetDynamicText(s global.String) {
	seed.OnGlobalChanged(s, func(q script.Ctx) {
		seed.Ctx(q).SetText(s.Get(q))
	})
}

//SetDynamicSource sets the source of the seed which will be dynamic at runtime.
func (seed Seed) SetDynamicSource(s global.String) {
	seed.OnGlobalChanged(s, func(q script.Ctx) {
		seed.Ctx(q).SetSource(s.Get(q))
	})
}
