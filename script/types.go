package script

import "github.com/qlova/script"

type Bool = script.Bool
type AnyCtx = script.AnyCtx
type String = script.String
type Native = script.Native
type Int = script.Int
type Value = script.Value

type Interface struct {
	Q Ctx
	Native
}
