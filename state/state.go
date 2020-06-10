package state

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/signal"
)

type State struct {
	Bool
	not bool
}

func New(options ...Option) State {
	return State{NewBool(options...), false}
}

func (state State) Signal() signal.Type {
	if state.not {
		return signal.Raw(state.key + ".unset")
	}
	return signal.Raw(state.key + ".set")
}

//GetBool implements script.AnyBool
func (state State) GetBool() script.Bool {
	if state.not {
		return state.Bool.GetBool().Not()
	}
	return state.Bool.GetBool()
}

//GetValue implements script.AnyValue
func (state State) GetValue() script.Value {
	return state.GetBool().Value
}

func (state State) Not() State {
	state.not = !state.not
	return state
}

func (state State) Toggle() script.Script {
	return func(q script.Ctx) {
		q.If(state, func(q script.Ctx) {
			state.Unset(q)
		}).Else(func(q script.Ctx) {
			state.Set(q)
		})
	}
}

//Set sets the state to be active.
func (state State) Set(q script.Ctx) {
	if state.not {
		if !state.ro {
			state.set(q, js.False)
		}

	} else {
		if !state.ro {
			state.set(q, js.True)
		}
	}
	q(signal.Emit(state.Signal()))
}

//Unset sets the state to not be active.
func (state State) Unset(q script.Ctx) {
	if state.not {
		state.set(q, js.True)

	} else {
		state.set(q, js.False)
	}
	q(signal.Emit(state.Signal()))
}

type data struct {
	seed.Data

	set, unset map[Bool]script.Script
	change     map[Value]script.Script

	onrefresh script.Script
	refresh   bool
}

var seeds = make(map[seed.Seed]data)

func (state State) Protect(s ...script.Script) script.Script {
	return js.If(state.Not(),
		js.Try(
			script.New(state.Set, script.New(s...), state.Unset),
		).Catch(
			script.New(state.Unset, js.Throw(js.NewValue("e"))),
			"e"),
	)
}

//If only applies its options if the state is active.
func (state State) If(options ...seed.Option) seed.Option {

	if state.Null() {
		return seed.NewOption(func(c seed.Seed) {})
	}

	return seed.NewOption(func(c seed.Seed) {
		switch c.(type) {
		case script.Seed, script.Undo:
			panic("state.State.If must not be called on a script.Seed")
		}

		If(state, options...).AddTo(c)

		var data data
		c.Read(&data)

		if data.change == nil {
			data.change = make(map[Value]script.Script)
		}

		c.Write(data)

		if state.dependencies == nil {
			data.change[state.Value] = data.change[state.Value].Append(Refresh(c))
		} else {
			for _, dep := range *state.dependencies {
				data.change[dep] = data.change[dep].Append(Refresh(c))
			}
		}

	})
}
