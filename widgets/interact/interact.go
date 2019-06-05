package interact

import "github.com/qlova/seed"
import "github.com/qlova/seed/script"

func init() {
	seed.Embed("/interact.js", []byte(Javascript))
}

type Interaction interface {
	AddTo(i Interactable)
}

type Interactable struct {
	register string
	seed     script.Seed
}

func With(seed script.Seed) Interactable {
	var register = script.Unique()
	seed.Javascript("let " + register + " = interact(" + seed.Element() + ");")
	return Interactable{seed: seed, register: register}
}

func (i Interactable) Add(in Interaction) {
	in.AddTo(i)
}

type Draggable struct {
	OnMove func()
}

func (d Draggable) AddTo(i Interactable) {
	i.seed.Javascript(i.register + ".draggable({")
	i.seed.Javascript("onmove: function(event) {")
	d.OnMove()
	i.seed.Javascript("}});")
}
