package script

import qlova "github.com/qlova/script"

//Go calls a Go function with the provided arguments.
func (q Ctx) Go(function interface{}, args ...qlova.AnyValue) Promise {
	return q.rpc(function, "undefined", nil, args...)
}
