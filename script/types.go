package script

import (
	"qlova.org/seed/js"
)

type Native = js.Value

type Bool = js.Bool
type String = js.String
type Number = js.Number
type Value = js.Value

type AnyValue = js.AnyValue

type Ctx = js.Ctx
type Script = js.Script

type Interface struct {
	Q Ctx
	Native
}
