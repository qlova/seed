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

func (seed Seed) VisibleWhen(state State) {
	seed.When(state, func(q Script) {
		seed.Script(q).SetVisible()
	})
	seed.When(state.Not(), func(q Script) {
		seed.Script(q).SetHidden()
	})

}

func (seed Seed) HiddenWhen(state State) {
	seed.When(state, func(q Script) {
		seed.Script(q).SetHidden()
	})
	seed.When(state.Not(), func(q Script) {
		seed.Script(q).SetVisible()
	})
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

func (seed Seed) OnClickSetState(state State) {
	seed.OnClick(func(q Script) {
		state.Set(q)
	})
}

func (seed Seed) OnClickUnsetState(state State) {
	seed.OnClick(func(q Script) {
		state.Unset(q)
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
	if state.not {
		return q.Not(state.Script(q))
	} else {
		return state.Script(q)
	}
}

func (state State) Set(q Script) {
	if state.not {
		q.Javascript(string(state.BoolVar.Variable) + `_unset();`)
	} else {
		q.Javascript(string(state.BoolVar.Variable) + `_set();`)
	}
}

func (state State) Unset(q Script) {
	if state.not {
		q.Javascript(string(state.BoolVar.Variable) + `_set();`)
	} else {
		q.Javascript(string(state.BoolVar.Variable) + `_unset();`)
	}
}

func (state State) UnsetFor(u User) {
	if state.not {
		u.Update.Evaluations["state"] = append(u.Update.Evaluations["state"], string(state.BoolVar.Variable)+`_set();`)
	}
	u.Update.Evaluations["state"] = append(u.Update.Evaluations["state"], string(state.BoolVar.Variable)+`_unset();`)
}

func (state State) SetFor(u User) {
	if state.not {
		u.Update.Evaluations["state"] = append(u.Update.Evaluations["state"], string(state.BoolVar.Variable)+`_unset();`)
	}
	u.Update.Evaluations["state"] = append(u.Update.Evaluations["state"], string(state.BoolVar.Variable)+`_set();`)
}
