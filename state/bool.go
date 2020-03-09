package state

import (
	"github.com/qlova/seed/script"
)

//Bool is a global Boolean.
type Bool struct {
	Value
}

//NewBool returns a reference to a new global boolean.
func NewBool(options ...Option) Bool {
	return Bool{newValue("false", options...)}
}

//BoolFromCtx implements script.AnyBool
func (b Bool) BoolFromCtx(q script.AnyCtx) script.Bool {
	return b.get(script.CtxFrom(q))
}

//Get the script.Bool for the global.Bool
func (b Bool) get(q script.Ctx) script.Bool {
	return q.Value(`(%v == "true")`, b.Value.get(q)).Bool()
}

//Set the global.Bool to be script.Bool
func (b Bool) set(q script.Ctx, value script.Bool) {
	b.Value.set(q, q.Value(`(%v).toString();`).String())
}
