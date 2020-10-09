package clientside

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
)

//Catch copies any raised errors into the provided String.
func Catch(into *String) seed.Option {
	return client.OnError(func(err client.String) client.Script {
		return into.SetTo(err)
	})
}
