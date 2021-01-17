package popup

import (
	"fmt"
	"reflect"
	"strings"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/new/html/div"
	"qlova.org/seed/set"
	"qlova.org/seed/set/transition"
	"qlova.org/seed/use/css"
	"qlova.org/seed/use/css/units/percentage/of"
	"qlova.org/seed/use/html"
	"qlova.org/seed/use/js"
)

func ID(p Popup) string {
	return strings.Replace(reflect.TypeOf(p).String(), ".", "_", -1)
}

type Popup interface {
	Popup() seed.Seed
}

//New returns a new popup seed.
func New(options ...seed.Option) seed.Seed {
	var Popup = div.New(html.SetTag("div"),

		css.SetDisplay(css.Flex),
		css.SetFlexDirection(css.Column),

		set.Size(100%of.Parent, 100%of.Parent),
		set.Layer(1),
		css.SetPosition(css.Absolute),

		transition.SetOnEnter(OnShow),
		transition.SetOnExit(OnHide),
	)

	for _, option := range options {
		option.AddTo(Popup)
	}

	return Popup
}

type data struct {
	popups map[reflect.Type]Popup
}

//Show shows the provided popup.
func Show(p Popup) js.Script {
	return func(q js.Ctx) {
		//Sort out script arguments of the page.
		popup, args := parseArgs(p)

		q(seed.Mutate(func(d *data) {
			if d.popups == nil {
				d.popups = make(map[reflect.Type]Popup)
			}

			d.popups[reflect.TypeOf(p)] = popup
		}))

		fmt.Fprintf(q, `seed.show("%v", %v);`, ID(p), args.GetObject().String())
	}
}

//Wrap shows the provided popup while the provided script is running.
func Wrap(p Popup, s ...client.Script) js.Script {
	return func(q js.Ctx) {

		//Sort out script arguments of the page.
		popup, args := parseArgs(p)

		q(seed.Mutate(func(d *data) {
			if d.popups == nil {
				d.popups = make(map[reflect.Type]Popup)
			}

			d.popups[reflect.TypeOf(p)] = popup
		}))

		fmt.Fprintf(q, `seed.show("%v", %v); try {`, ID(p), args.GetObject().String())
		client.NewScript(s...).GetScript()(q)
		fmt.Fprintf(q, `seed.hide("%[1]v"); } catch(e) { seed.hide("%[1]v"); throw e;  }`, ID(p))
	}
}

//Hide hides the provided popup.
func Hide(p Popup) js.Script {
	return func(q js.Ctx) {
		fmt.Fprintf(q, `seed.hide("%v");`, ID(p))
	}
}

func OnShow(f ...client.Script) seed.Option {
	return client.On("show", f...)
}

func OnHide(f ...client.Script) seed.Option {
	return client.On("hide", f...)
}
