package secrets

import (
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/state"
)

func init() {
	var argon2, _ = Asset("argon2.js")
	var wasm, _ = Asset("argon2.wasm")

	js.NewImport("/argon2-browser.js", Lib)
	js.NewImport("/argon2.js", string(argon2))
	js.NewImport("/argon2.wasm", string(wasm))
}

//State is a client-state secret. If the value is read, the hash is returned.
//Use this for passwords.
type State struct {
	Salt string
	state.String
}

//NewState returns a new Secret with the given salt.
func NewState(salt string, options ...state.Option) State {
	if len(salt) < 8 {
		salt = salt + salt
	}
	return State{salt, state.NewString("", options...)}
}

//GetString implements script.AnyString
func (s State) GetString() script.String {
	return js.String{Value: js.NewValue(`"'#(import "/argon2-browser.js")(await argon2.hash({
		pass: %v,
		salt: %v,

		time: 1,
		mem: 1024*64,
		hashLen: 32,
		parallelism: 1,
		type: argon2.ArgonType.Argon2di
	})).hashHex`, s.String.GetValue(), js.NewString(s.Salt))}
}

//GetBool implements script.AnyBool
func (s State) GetBool() script.Bool {
	return s.String.GetBool()
}

//GetValue implements script.AnyValue
func (s State) GetValue() script.Value {
	return s.GetString().Value
}
