package client

import (
	"qlova.org/seed"
	"qlova.org/seed/script"
)

//OnRender is called whenever this seed is asked to render itself.
func OnRender(do script.Script) seed.Option {
	return script.On("render", do)
}
