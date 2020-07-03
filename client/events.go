package client

import (
	"qlova.org/seed"
	"qlova.org/seed/script"
)

//OnClick is called when the seed is clicked by the client.
func OnClick(do ...Script) seed.Option {
	return script.OnClick(script.New(do...))
}

//OnLoad is called when the seed is loaded by the client.
//this will only be called once per app-launch, or in the case of dynamic content, once when the seed is created.
func OnLoad(do ...Script) seed.Option {
	return script.OnReady(script.New(do...))
}
