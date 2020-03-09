package state

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
)

type State struct {
	Bool
	not bool
}

func New(options ...Option) State {
	return State{NewBool(options...), false}
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
	if state.ro {
		return
	}

	var reference = state.key
	if state.not {
		state.set(q, q.False)
		q.Javascript(`if (seed.state["` + reference + `"])`)
		q.Javascript(`await seed.state["` + reference + `"].unset();`)

	} else {
		state.set(q, q.True)
		q.Javascript(`if (seed.state["` + reference + `"])`)
		q.Javascript(`await seed.state["` + reference + `"].set();`)
	}
}

//Unset sets the state to not be active.
func (state State) Unset(q script.Ctx) {
	if state.ro {
		return
	}

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

type data struct {
	set, unset map[State]script.Script
	change     map[Value]script.Script
}

var seeds = make(map[seed.Seed]data)

//If only applies its options if the state is active.
func (state State) If(options ...seed.Option) seed.Option {
	return seed.NewOption(func(s seed.Any) {
		data := seeds[s.Root()]
		if data.set == nil {
			data.set = make(map[State]script.Script)
			data.unset = make(map[State]script.Script)
		}
		data.set[state] = data.set[state].Then(func(q script.Ctx) {
			for _, option := range options {
				option.Apply(s.Root().Ctx(q))
			}
		})
		data.unset[state] = data.unset[state].Then(func(q script.Ctx) {
			for _, option := range options {
				option.Reset(s.Root().Ctx(q))
			}
		})
		seeds[s.Root()] = data
	}, func(s seed.Ctx) {
		panic("nested state.If is not allowed")
	}, func(s seed.Ctx) {
		panic("nested state.If is not allowed")
	})
}
