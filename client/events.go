package client

import (
	"fmt"
	"io/ioutil"

	"qlova.org/seed"
	"qlova.org/seed/use/js"
)

func on(event string, do Script) seed.Option {
	if do == nil {
		return seed.NewOption(func(c seed.Seed) {})
	}
	return seed.NewOption(func(c seed.Seed) {
		do.GetScript()(js.NewCtx(ioutil.Discard, c)) //Catch errors and harvest pages.

		var d Data
		c.Load(&d)

		switch mode, q := Seed(c); mode {
		case AddTo:
			c.Use()
			q(fmt.Sprintf(`seed.on(%v, "%v", async function() {`, Element(c), event))
			do.GetScript()(q)
			q(fmt.Sprintf(`}, %v);`, c.ID())) //add dynamic event handler by ID.
		case Undo:
			q(fmt.Sprintf(`seed.off(%v, "%v", '%v');`, Element(c), event, c.ID())) //remove id.
		default:
			//s.Root().Use()
			if d.On == nil {
				d.On = make(map[string]js.Script)
			}
			if event == "error" {
				d.On[event] = d.On[event].Append(do.GetScript())
			} else {
				d.On[event] = d.On[event].Append(func(q js.Ctx) {
					q("try {")
					q(do.GetScript())
					fmt.Fprintf(q, "} catch(e) { seed.report(e, %v) }", Element(c))
				})
			}
		}

		c.Save(d)
	})
}

//On is called when the seed has the given event triggered.
func On(event string, do ...Script) seed.Option {
	return on(event, NewScript(do...).GetScript())
}

//ClickCursor is an option that is added to all OnClicks.
var ClickCursor seed.Option

//OnClick is called when the seed is clicked by the client.
func OnClick(do ...Script) seed.Option {
	return seed.Options{
		On("click", do...),
		ClickCursor,
	}
}

//OnInput is called when the seed recieves input from the client.
func OnInput(do ...Script) seed.Option {
	return On("input", do...)
}

//OnChange is called when the seed is changed by client interaction.
func OnChange(do ...Script) seed.Option {
	return On("change", do...)
}

//OnEnterKey is called when the client presses the Enter key whilst focused on this seed.
func OnEnterKey(do ...Script) seed.Option {
	return On("enter", do...)
}

//OnLoad is called when the seed is loaded by the client.
//this will only be called once per app-launch, or in the case of dynamic content, once when the seed is created.
func OnLoad(do ...Script) seed.Option {
	return On("ready", do...)
}

//OnError calls the provided script when there is an error not handled by this seed or any children seeds.
func OnError(do func(err String) Script) seed.Option {
	return On("error", NewScript(
		do(js.String{Value: js.NewValue(`arguments[0]`)}),
	))
}

//OnRender is called whenever this seed is asked to render itself.
func OnRender(do ...Script) seed.Option {
	return On("render", do...)
}
