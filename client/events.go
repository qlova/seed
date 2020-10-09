package client

import (
	"fmt"
	"io/ioutil"

	"qlova.org/seed"
	"qlova.org/seed/js"
)

func on(event string, do Script) seed.Option {
	if do == nil {
		return seed.NewOption(func(c seed.Seed) {})
	}
	return seed.NewOption(func(c seed.Seed) {
		do.GetScript()(js.NewCtx(ioutil.Discard, c)) //Catch errors and harvest pages.

		var d Data
		c.Read(&d)

		switch data := c.(type) {
		case Seed:
			c.Use()
			data.Q(fmt.Sprintf(`seed.on(%v, "%v", async function() {`, data.Element(), event))
			do.GetScript()(data.Q)
			data.Q(`});`)
		case Undo:
			//s.Root().Use()
			data.Q(fmt.Sprintf(`seed.on(%v, "%v", async function() {`, data.Element(), event))
			if d.On[event] != nil {
				d.On[event].GetScript()(js.NewCtx(data.Q))
			}
			data.Q(`});`)
		default:
			//s.Root().Use()
			if d.On == nil {
				d.On = make(map[string]js.Script)
			}
			d.On[event] = d.On[event].Append(do.GetScript())
		}

		c.Write(d)
	})
}

//On is called when the seed has the given event triggered.
func On(event string, do ...Script) seed.Option {
	return on(event, NewScript(do...).GetScript())
}

//OnClick is called when the seed is clicked by the client.
func OnClick(do ...Script) seed.Option {
	return On("click", do...)
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
