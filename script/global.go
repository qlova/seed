package script

import (
	"strconv"

	"github.com/qlova/seed/js"
)

type Global struct {
	id int
	js.Value
}

var id int

func (g Global) Set(v js.AnyValue) js.Script {
	return js.Global().Get("seed").Get("globals").Set(strconv.Itoa(g.id), v)
}

func NewGlobal(initial js.AnyValue) Global {
	id++
	if initial == nil {
		return Global{id, js.NewValue(`seed.globals[` + strconv.Itoa(id) + `]`)}
	}
	return Global{id, js.NewValue(`seed.globals[`+strconv.Itoa(id)+`] || %v`, initial)}
}
