package popup

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
)

func ID(p Popup) string {
	return strings.Replace(reflect.TypeOf(p).String(), ".", "_", -1)
}

type Popup interface {
	Popup(Seed)
}

type Seed struct {
	seed.Seed
}

func Scope(c seed.Seed) Seed {
	return Seed{c}
}

type data struct {
	seed.Data

	popups map[reflect.Type]Popup
}

//Show shows the provided popup.
func (c Seed) Show(p Popup) script.Script {
	return func(q script.Ctx) {

		var data data
		c.Read(&data)
		if data.popups == nil {
			data.popups = make(map[reflect.Type]Popup)
			c.Write(data)
		}

		data.popups[reflect.TypeOf(p)] = p

		fmt.Fprintf(q, `seed.show("%v");`, ID(p))
	}
}

//Wrap shows the provided popup while the provided script is running.
func (c Seed) Wrap(p Popup, s script.Script) script.Script {
	return func(q script.Ctx) {

		var data data
		c.Read(&data)
		if data.popups == nil {
			data.popups = make(map[reflect.Type]Popup)
			c.Write(data)
		}

		data.popups[reflect.TypeOf(p)] = p

		fmt.Fprintf(q, `seed.show("%v"); try {`, ID(p))
		s(q)
		fmt.Fprintf(q, `seed.hide("%[1]v"); } catch(e) { seed.hide("%[1]v"); throw e; }`, ID(p))
	}
}

//Hide hides the provided popup.
func (c Seed) Hide(p Popup) script.Script {
	return func(q script.Ctx) {
		fmt.Fprintf(q, `seed.hide("%v");`, ID(p))
	}
}

func OnShow(f script.Script) seed.Option {
	return script.On("show", f)
}

func OnHide(f script.Script) seed.Option {
	return script.On("hide", f)
}
