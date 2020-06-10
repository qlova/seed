package popup

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/qlova/seed"
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/s/html/div"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/style"
)

func ID(p Popup) string {
	return strings.Replace(reflect.TypeOf(p).String(), ".", "_", -1)
}

type Manager struct {
	seed.Seed
}

type Popup interface {
	Popup(Manager) Seed
}

type Seed struct {
	seed.Seed
}

func New(options ...seed.Option) Seed {
	var Popup = Seed{div.New(html.SetTag("div"),

		css.SetDisplay(css.Flex),
		css.SetFlexDirection(css.Column),

		style.SetSize(100, 100),
		style.SetLayer(1),
		css.SetPosition(css.Absolute),
	)}

	for _, option := range options {
		option.AddTo(Popup)
	}

	return Popup
}

func ManagerOf(c seed.Seed) Manager {
	return Manager{c}
}

type data struct {
	seed.Data

	popups map[reflect.Type]Popup
}

//Show shows the provided popup.
func (c Manager) Show(p Popup) script.Script {
	return func(q script.Ctx) {

		//Sort out script arguments of the page.
		popup, args := parseArgs(p)

		var data data
		c.Read(&data)
		if data.popups == nil {
			data.popups = make(map[reflect.Type]Popup)
			c.Write(data)
		}

		data.popups[reflect.TypeOf(p)] = popup

		fmt.Fprintf(q, `seed.show("%v", %v);`, ID(p), args.GetObject().String())
	}
}

//Wrap shows the provided popup while the provided script is running.
func (c Manager) Wrap(p Popup, s script.Script) script.Script {
	return func(q script.Ctx) {

		//Sort out script arguments of the page.
		popup, args := parseArgs(p)

		var data data
		c.Read(&data)
		if data.popups == nil {
			data.popups = make(map[reflect.Type]Popup)
			c.Write(data)
		}

		data.popups[reflect.TypeOf(p)] = popup

		fmt.Fprintf(q, `seed.show("%v", %v); try {`, ID(p), args.GetObject().String())
		s(q)
		fmt.Fprintf(q, `seed.hide("%[1]v"); } catch(e) { seed.hide("%[1]v"); throw e;  }`, ID(p))
	}
}

//Hide hides the provided popup.
func (c Manager) Hide(p Popup) script.Script {
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
