package signal

import (
	"encoding/base64"
	"io/ioutil"
	"math/big"

	"github.com/qlova/seed/js"

	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
)

type data struct {
	seed.Data

	handlers map[Type]script.Script
}

//Type is a type of signal.
type Type struct {
	string
}

func Raw(name string) Type {
	return Type{name}
}

var id int64

//New returns a new signal type.
func New() Type {
	id++
	return Type{base64.RawURLEncoding.EncodeToString(big.NewInt(id).Bytes())}
}

//On runs the script when the signal is emitted.
func On(signal Type, do script.Script) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		js.NewCtx(ioutil.Discard, c)(do)

		var data data
		c.Read(&data)

		if data.handlers == nil {
			data.handlers = make(map[Type]script.Script)
			c.Write(data)
		}

		data.handlers[signal] = data.handlers[signal].Append(func(q script.Ctx) {
			q(js.Global().Get("seed").Set("active", script.Element(c).Value))
			q(do)
		})
	})
}

//Emit returns a script that emits the given signal.
func Emit(signal Type) script.Script {
	return func(q script.Ctx) {
		q.Run("if (seed.signal['" + signal.string + "']) await seed.signal['" + signal.string + "']")
	}
}
