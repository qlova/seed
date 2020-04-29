package state

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/signal"
	"github.com/qlova/seed/user"
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
		return signal.Raw("state.unset." + state.key)
	}
	return signal.Raw("state.set." + state.key)
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
	var reference = state.key
	if state.not {
		if !state.ro {
			state.set(q, js.False)
		}
		q(`if (seed.state["` + reference + `"])`)
		q(`await seed.state["` + reference + `"].unset();`)

	} else {
		if !state.ro {
			state.set(q, js.True)
		}
		q(`if (seed.state["` + reference + `"])`)
		q(`await seed.state["` + reference + `"].set();`)
	}
}

//Unset sets the state to not be active.
func (state State) Unset(q script.Ctx) {
	var reference = state.key
	if state.not {
		state.set(q, js.True)
		q(`if (seed.state["` + reference + `"])`)
		q(`await seed.state["` + reference + `"].set();`)
	} else {
		state.set(q, js.False)
		q(`if (seed.state["` + reference + `"])`)
		q(`await seed.state["` + reference + `"].unset();`)

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
	var state = s.s
	if state.not {
		state.setFor(s.u, "false")
	} else {
		state.setFor(s.u, "true")
	}

}

type data struct {
	seed.Data

	set, unset map[Bool]script.Script
	change     map[Value]script.Script

	onrefresh script.Script
	refresh   bool
}

var seeds = make(map[seed.Seed]data)

//If only applies its options if the state is active.
func (state State) If(options ...seed.Option) seed.Option {

	return seed.NewOption(func(c seed.Seed) {
		switch c.(type) {
		case script.Seed, script.Undo:
			panic("state.State.If must not be called on a script.Seed")
		}

		//Add any children seeds to the parent seed.
		//Hacky fix.
		for _, o := range options {
			switch child := o.(type) {
			case seed.Seed:
				c.Add(child)
			}
		}

		var data data
		c.Read(&data)

		if data.set == nil {
			data.set = make(map[Bool]script.Script)
			c.Write(data)
		}

		if data.unset == nil {
			data.unset = make(map[Bool]script.Script)
			c.Write(data)
		}

		if state.not {
			data.unset[state.Bool] = data.unset[state.Bool].Append(func(q script.Ctx) {
				for _, option := range options {
					if other, ok := option.(seed.Seed); ok {
						script.Scope(other, q).AddTo(script.Scope(c, q))
					} else {
						option.AddTo(script.Scope(c, q))
					}
				}
			})
			data.set[state.Bool] = data.set[state.Bool].Append(func(q script.Ctx) {
				for _, option := range options {
					if other, ok := option.(seed.Seed); ok {
						script.Scope(c, q).Undo(script.Scope(other, q))
					} else {
						script.Scope(c, q).Undo(option)
					}
				}
			})
		} else {
			data.set[state.Bool] = data.set[state.Bool].Append(func(q script.Ctx) {
				for _, option := range options {
					if other, ok := option.(seed.Seed); ok {
						script.Scope(other, q).AddTo(script.Scope(c, q))
					} else {
						option.AddTo(script.Scope(c, q))
					}
				}
			})
			data.unset[state.Bool] = data.unset[state.Bool].Append(func(q script.Ctx) {
				for _, option := range options {
					if other, ok := option.(seed.Seed); ok {
						script.Scope(c, q).Undo(script.Scope(other, q))
					} else {
						script.Scope(c, q).Undo(option)
					}
				}
			})
		}
	})
}
