package script

import (
	"fmt"

	"github.com/qlova/script"
	"github.com/qlova/script/language"
)

func ToJavascript(s Script) []byte {
	return language.Javascript(func(q script.Ctx) {
		s(Ctx{q})
	})
}

//Javascript inserts raw js into the script.
func (q Ctx) Javascript(js string, args ...interface{}) {
	var converted = make([]interface{}, len(args))
	for i := range args {
		if v, ok := args[i].(script.Value); ok {
			converted[i] = q.Raw(v)
		} else {
			converted[i] = args[i]
		}
	}

	if len(args) > 0 {
		q.Write([]byte(fmt.Sprintf(js, converted...)))
	} else {
		q.Write([]byte(fmt.Sprint(js)))
	}
}

//Value is any script value.
type value struct {
	q   Ctx
	raw string
}

//Bool returns the value as a bool.
func (v value) Bool() Bool {
	return Bool{language.Expression(v.q, v.raw)}
}

//String returns the value as a string.
func (v value) String() String {
	return String{language.Expression(v.q, v.raw)}
}

//Bool returns the value as a bool.
func (v value) Int() Int {
	return Int{language.Expression(v.q, v.raw)}
}

//Promise returns the value as a promise.
func (v value) Promise() Promise {

	return Promise{Native{language.Expression(v.q, v.raw)}, v.q}
}

//Promise returns the value as a promise.
func (v value) File() File {
	return File{v.Interface()}
}

//Interface returns the value as a interface.
func (v value) Interface() Interface {
	return Interface{v.q, Native{language.Expression(v.q, v.raw)}}
}

//Value wraps a JS string as a value that can be cast to script.Type.
func (q Ctx) Value(format string, args ...interface{}) value {

	var converted = make([]interface{}, len(args))
	for i := range args {
		if v, ok := args[i].(script.Value); ok {
			converted[i] = q.Raw(v)
		} else {
			converted[i] = args[i]
		}
	}

	if len(args) > 0 {
		return value{q, fmt.Sprintf(format, converted...)}
	}
	return value{q, format}
}
