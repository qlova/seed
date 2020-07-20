package client

import (
	"qlova.org/seed"
	"qlova.org/seed/script"
)

//On is called when the seed has the given event triggered.
func On(event string, do ...Script) seed.Option {
	return script.On(event, NewScript(do...).GetScript())
}

//OnClick is called when the seed is clicked by the client.
func OnClick(do ...Script) seed.Option {
	return script.OnClick(NewScript(do...).GetScript())
}

//OnInput is called when the seed recieves input from the client.
func OnInput(do ...Script) seed.Option {
	return script.OnInput(NewScript(do...).GetScript())
}

//OnChange is called when the seed is changed by client interaction.
func OnChange(do ...Script) seed.Option {
	return script.OnChange(NewScript(do...).GetScript())
}

//OnEnterKey is called when the client presses the Enter key whilst focused on this seed.
func OnEnterKey(do ...Script) seed.Option {
	return script.OnEnter(NewScript(do...).GetScript())
}

//OnLoad is called when the seed is loaded by the client.
//this will only be called once per app-launch, or in the case of dynamic content, once when the seed is created.
func OnLoad(do ...Script) seed.Option {
	return script.OnReady(NewScript(do...).GetScript())
}
