package state

import (
	"fmt"
	"strconv"

	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/user"
)

//String is a global String.
type String struct {
	Value
}

//NewString returns a reference to a new global string.
func NewString(initial string, options ...Option) String {
	return String{newValue(strconv.Quote(initial), options...)}
}

//GetString implements script.AnyString
func (s String) GetString() script.String {
	return s.get()
}

//GetValue implements script.AnyValue
func (s String) GetValue() script.Value {
	return s.get().Value
}

//Set allows setting the value of a String in the given script ctx.
func (s String) Set(value script.String) script.Script {
	return func(q script.Ctx) {
		s.set(q, value)
	}
}

//SetL allows setting the value of a String to a literal in the given script ctx.
func (s String) SetL(literal string) script.Script {
	return func(q script.Ctx) {
		s.set(q, q.String(literal))
	}
}

type RemoteString struct {
	u user.Ctx
	s String
}

func (s String) For(u user.Ctx) RemoteString {
	return RemoteString{u, s}
}

func (s RemoteString) Set(value string) {
	s.s.setFor(s.u, value)
}

func (s String) SetText() seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		switch c.(type) {
		case script.Seed, script.Undo:
			panic("state.String.SetText must not be called on a script.Seed")
		}

		if s.raw != "" {
			c.Add(script.OnReady(func(q script.Ctx) {
				fmt.Fprintf(q, `%v.innerText = %v;`, script.Scope(c, q).Element(), s.get())
			}))
		}

		if s.key != "" {
			var data data
			c.Read(&data)

			if data.change == nil {
				data.change = make(map[Value]script.Script)
			}

			data.change[s.Value] = data.change[s.Value].Append(func(q script.Ctx) {
				q(fmt.Sprintf(`%v.innerText = %v;`, script.Scope(c, q).Element(), s.get()))
			})

			if s.dependencies != nil {
				for _, dep := range *s.dependencies {
					data.change[dep] = data.change[dep].Append(func(q script.Ctx) {
						q(fmt.Sprintf(`seed.state["%v"].changed();`, s.key))
					})
				}
			}

			c.Write(data)
		}
	})
}
