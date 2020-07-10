package clientside

import (
	"qlova.org/seed"
	"qlova.org/seed/script"
)

//Catch copies any raised errors into the provided String.
func Catch(into *String) seed.Option {
	return script.OnError(func(q script.Ctx, err script.Error) {
		q(into.SetTo(err.String))
	})
}
