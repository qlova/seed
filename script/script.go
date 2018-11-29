package script

import qlova "github.com/qlova/script"
import "github.com/qlova/script/language"
import "github.com/qlova/script/language/javascript"

type String = qlova.String
type Expression string

type Seed struct {
	ID string
	Qlovascript qlova.Script
}

func (seed Seed) Javascript(js string) {
	seed.Qlovascript.Raw("Javascript", language.Statement(js))
}

type File Expression

func (seed Seed) SetText(s qlova.String) {
	seed.Javascript(`get("`+seed.ID+`").textContent = `+s.Raw()+`;`)
}

func (seed Seed) SetLeft(s qlova.String) {
	seed.Javascript(`get("`+seed.ID+`").style.left = `+s.Raw()+`;`)
}

func (seed Seed) SetDisplay(s qlova.String) {
	seed.Javascript(`get("`+seed.ID+`").style.display = `+s.Raw()+`;`)
}

func (seed Seed) SetVisible() {
	seed.Javascript(`get("`+seed.ID+`").style.display = "block";`)
}

func (seed Seed) SetHidden() {
	seed.Javascript(`get("`+seed.ID+`").style.display = "none";`)
}

func (seed Seed) Click() {
	seed.Javascript(`get("`+seed.ID+`").click();`)
}

func (seed Seed) Left() qlova.String {
	return seed.Qlovascript.Wrap(Javascript.String(`get("`+seed.ID+`").style.left`)).(qlova.String)
}

func (seed Seed) Width() qlova.String {
	return seed.Qlovascript.Wrap(Javascript.String(`getComputedStyle(get("`+seed.ID+`")).width`)).(qlova.String)
}

func (seed Seed) Value() qlova.String {
	return seed.Qlovascript.Wrap(Javascript.String(`get("`+seed.ID+`").value`)).(qlova.String)
}
