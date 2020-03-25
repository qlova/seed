package state

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/user"
)

type State struct {
	Bool
	not bool
}

func New(options ...Option) State {
	return State{NewBool(options...), false}
}

func (state State) Not() State {
	state.not = !state.not
	return state
}

func (state State) Toggle() script.Script {
	return func(q script.Ctx) {
		q.If(state, func() {
			state.Unset(q)
		}).Else(func() {
			state.Set(q)
		}).End()
	}
}

//Set sets the state to be active.
func (state State) Set(q script.Ctx) {
	var reference = state.key
	if state.not {
		if !state.ro {
			state.set(q, q.False)
		}
		q.Javascript(`if (seed.state["` + reference + `"])`)
		q.Javascript(`await seed.state["` + reference + `"].unset();`)

	} else {
		if !state.ro {
			state.set(q, q.True)
		}
		q.Javascript(`if (seed.state["` + reference + `"])`)
		q.Javascript(`await seed.state["` + reference + `"].set();`)
	}
}

//Unset sets the state to not be active.
func (state State) Unset(q script.Ctx) {
	var reference = state.key
	if state.not {
		state.set(q, q.True)
		q.Javascript(`if (seed.state["` + reference + `"])`)
		q.Javascript(`await seed.state["` + reference + `"].set();`)
	} else {
		state.set(q, q.False)
		q.Javascript(`if (seed.state["` + reference + `"])`)
		q.Javascript(`await seed.state["` + reference + `"].unset();`)

	}
}

type RemoteState struct {
	u user.Ctx
	s State
}

func (s State) For(u user.Ctx) RemoteState {
	return RemoteState{u, s}
}

func (s RemoteState) Set() {
	s.s.setFor(s.u, "true")
}

type data struct {
	seed.Data

	set, unset map[State]script.Script
	change     map[Value]script.Script
}

var seeds = make(map[seed.Seed]data)

//If only applies its options if the state is active.
func (state State) If(options ...seed.Option) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		switch c.(type) {
		case script.Seed, script.Undo:
			panic("state.State.If must not be called on a script.Seed")
		}

		var data data
		c.Read(&data)

		if data.set == nil {
			data.set = make(map[State]script.Script)
			data.unset = make(map[State]script.Script)
		}

		data.set[state] = data.set[state].Then(func(q script.Ctx) {
			for _, option := range options {
				if other, ok := option.(seed.Seed); ok {
					q.Scope(other).AddTo(q.Scope(c))
				} else {
					option.AddTo(q.Scope(c))
				}
			}
		})
		data.unset[state] = data.unset[state].Then(func(q script.Ctx) {
			for _, option := range options {
				if other, ok := option.(seed.Seed); ok {
					q.Scope(c).Undo(q.Scope(other))
				} else {
					q.Scope(c).Undo(option)
				}
			}
		})

		c.Write(data)
	})
}
