package state

import (
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
)

type Secret struct {
	salt string
	Value
}

func NewSecret(salt string, options ...Option) Secret {
	return Secret{salt, newValue("", options...)}
}

//StringFromCtx implements script.AnyString
func (s Secret) StringFromCtx(q script.Ctx) script.String {
	//script.CtxFrom(q).Require(`/argon2-browser.js`)
	return js.String{Value: js.NewValue(`(await argon2.hash({
		pass: %v,
		salt: %v,

		time: 1,
		mem: 1024*64,
		hashLen: 32,
		parallelism: 1,
		type: argon2.ArgonType.Argon2di
	})).hashHex`, s.get(), js.NewString(s.salt))}
}

//GetValue implements script.AnyValue
func (s Secret) GetValue() script.Value {
	return s.get().Value
}

//Set allows setting the value of a String in the given script ctx.
func (s Secret) Set(value script.String) script.Script {
	return func(q script.Ctx) {
		s.set(q, value)
	}
}

//SetL allows setting the value of a String to a literal in the given script ctx.
func (s Secret) SetL(literal string) script.Script {
	return func(q script.Ctx) {
		s.set(q, q.String(literal))
	}
}
