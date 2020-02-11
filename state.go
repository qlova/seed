package seed

import (
	"fmt"

	"github.com/qlova/seed/script"
	"github.com/qlova/seed/script/global"
)

//State is a boolean state that can propogate effects to seeds.
type State struct {
	global.Bool
	not bool

	readonly bool
}

//Null returns true if this is a null state.
func (state State) Null() bool {
	return state.Bool.Ref() == ""
}

//NewState returns a new globally unique state.
//Name is ignored when built with 'production' flag.
func NewState(name ...string) State {
	if len(name) > 0 && !Production {
		return State{global.NewBool("state_" + name[0]), false, false}
	}
	return State{global.NewBool(), false, false}
}

//Installed is active whenever the app is installed to the device.
var Installed = NewState("installed")

//Not returns the inverse state.
//ie var On = NewState(); var Off = On.Not()
func (state State) Not() State {
	var s = state
	s.not = !s.not
	return s
}

//VisibleWhen sets a seed to be visible when the state is active.
func (seed Seed) VisibleWhen(state State) {
	seed.When(state, func(q script.Ctx) {
		seed.Ctx(q).SetVisible()
	})
	seed.When(state.Not(), func(q script.Ctx) {
		seed.Ctx(q).SetHidden()
	})

}

//HiddenWhen sets a seed to be hidden when the state is active.
func (seed Seed) HiddenWhen(state State) {
	seed.When(state, func(q script.Ctx) {
		seed.Ctx(q).SetHidden()
	})
	seed.When(state.Not(), func(q script.Ctx) {
		seed.Ctx(q).SetVisible()
	})
}

//When runs a script whenever the state becomes active.
func (seed Seed) When(state State, f func(q script.Ctx)) {
	if seed.states == nil {
		seed.states = make(map[State]func(script.Ctx))
	}
	if old, ok := seed.states[state]; ok {
		seed.states[state] = func(q script.Ctx) {
			old(q)
			f(q)
		}
		return
	}
	seed.states[state] = f
}

//OnClickToggleState is shorthand.
func (seed Seed) OnClickToggleState(state State) {
	seed.OnClick(func(q script.Ctx) {
		state.Toggle(q)
	})
}

//OnClickSetState is shorthand.
func (seed Seed) OnClickSetState(state State) {
	seed.OnClick(func(q script.Ctx) {
		state.Set(q)
	})
}

//OnClickUnsetState is shorthand.
func (seed Seed) OnClickUnsetState(state State) {
	seed.OnClick(func(q script.Ctx) {
		state.Unset(q)
	})
}

//Toggle toggles a state.
func (state State) Toggle(q script.Ctx) {
	q.If(state.Get(q), func() {
		state.Unset(q)
	}, q.Else(func() {
		state.Set(q)
	}))
}

//Get gets the state as a bool.
func (state State) Get(q script.Ctx) script.Bool {
	if state.not {
		return q.Not(state.Bool.Get(q))
	}
	return state.Bool.Get(q)
}

//Set sets the state to be active.
func (state State) Set(q script.Ctx) {
	if state.readonly {
		return
	}

	var reference = state.Bool.Ref()
	if state.not {
		state.Bool.Set(q, q.False())
		q.Javascript(`if (window.` + reference + `_unset)`)
		q.Javascript(reference + `_unset();`)
	} else {
		state.Bool.Set(q, q.True())
		q.Javascript(`if (window.` + reference + `_set)`)
		q.Javascript(reference + `_set();`)
	}
}

//Unset sets the state to not be active.
func (state State) Unset(q script.Ctx) {
	if state.readonly {
		return
	}

	var reference = state.Bool.Ref()
	if state.not {
		state.Bool.Set(q, q.True())
		q.Javascript(`if (window.` + reference + `_set)`)
		q.Javascript(reference + `_set();`)
	} else {
		state.Bool.Set(q, q.False())
		q.Javascript(`if (window.` + reference + `_unset)`)
		q.Javascript(reference + `_unset();`)
	}
}

//UnsetFor unsets a state for tthe specified user.
func (state State) UnsetFor(u User) {
	if state.readonly {
		return
	}

	var reference = state.Bool.Ref()
	if state.not {
		u.Execute(fmt.Sprintf(`if (window.%v_set) %v_set();`, reference, reference))
		return
	}
	u.Execute(fmt.Sprintf(`if (window.%v_unset) %v_unset();`, reference, reference))
}

//SetFor sets a state for tthe specified user.
func (state State) SetFor(u User) {
	if state.readonly {
		return
	}

	var reference = state.Bool.Ref()
	if state.not {
		u.Execute(fmt.Sprintf(`if (window.%v_unset) %v_unset();`, reference, reference))
		return
	}
	u.Execute(fmt.Sprintf(`if (window.%v_set) %v_set();`, reference, reference))
}
