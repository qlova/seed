package seed

import "github.com/qlova/seed/script"

type State struct {
	script.BoolVar
	not bool
}

func NewState() State {
	return State{script.NewBool(), false}
}

func (state State) Not() State {
	var s = state
	s.not = !s.not
	return s
}

func (seed Seed) When(state State, f func(q Script)) {
	if seed.states == nil {
		seed.states = make(map[State]func(Script))
	}
	seed.states[state] = f
}

func (seed Seed) OnClickToggleState(state State) {
	seed.OnClick(func(q Script) {
		state.Toggle(q)
	})
}

func (state State) Toggle(q Script) {
	q.If(state.Get(q), func() {
		state.Unset(q)
	}, q.Else(func() {
		state.Set(q)
	}))
}

func (state State) Get(q Script) script.Bool {
	return state.Script(q)
}

func (state State) Set(q Script) {
	if state.not {
		state.Unset(q)
		return
	}
	q.Javascript(string(state.BoolVar.Variable) + `_set();`)
}

func (state State) Unset(q Script) {
	if state.not {
		state.Set(q)
		return
	}
	q.Javascript(string(state.BoolVar.Variable) + `_unset();`)
}
