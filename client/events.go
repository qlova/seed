package client

import (
	"qlova.org/seed"
	"qlova.org/seed/script"
)

//OnClick is called when the seed is clicked.
func OnClick(do ...Script) seed.Option {
	return script.OnClick(script.New(do...))
}
