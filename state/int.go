package state

import (
	"fmt"
	"strconv"

	"qlova.org/seed"
	"qlova.org/seed/js"
	"qlova.org/seed/script"
	"qlova.org/seed/user"
)

//Int is a global Int.
type Int struct {
	Value
}

//NewInt returns a reference to a new global int.
func NewInt(initial int) Int {
	return Int{newValue(strconv.Itoa(initial))}
}

func (i Int) Increment() script.Script {
	return func(q script.Ctx) {
		i.set(q, i.get().Plus(q.Number(1)))
	}
}

//GetNumber implements script.AnyNumber
func (i Int) GetNumber() script.Number {
	return i.get()
}

//GetValue implements script.AnyValue
func (i Int) GetValue() script.Value {
	return i.get().Value
}

func (i Int) get() script.Number {
	return js.Number{Value: js.NewValue(`parseInt(%v)`, i.Value.get())}
}

//SetL sets the value of the Int with a literal.
func (i Int) SetL(value int) script.Script {
	return func(q script.Ctx) {
		i.set(q, q.Number(float64(value)))
	}
}

func (i Int) set(q script.Ctx, value script.Number) {
	i.Value.set(q, js.String{Value: js.NewValue(`(%v).toString()`, value)})
}

//SetText sets the seed's text to reflect the value of this Int.
func (i Int) SetText() seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		switch c.(type) {
		case script.Seed, script.Undo:
			panic("state.Int.SetText must not be called on a script.Seed")
		}

		if i.raw == "" {
			var d data
			c.Read(&d)

			if d.change == nil {
				d.change = make(map[Value]script.Script)
			}

			d.change[i.Value] = d.change[i.Value].Append(func(q script.Ctx) {
				q(fmt.Sprintf(`%v.innerText = (%v).toString();`, script.Scope(c, q).Element(), i.get().GetValue()))
			})

			c.Write(d)
		} else {
			c.With(script.OnReady(func(q script.Ctx) {
				fmt.Fprintf(q, `%v.innerText = (%v).toString();`, script.Scope(c, q).Element(), i.get().GetValue())
			}))
		}
	})
}

//RemoteInt is a remote reference to an Int.
type RemoteInt struct {
	u user.Ctx
	i Int
}

//For returns this Int as a RemoteInt associated with this user.
func (i Int) For(u user.Ctx) RemoteInt {
	return RemoteInt{u, i}
}

//Set sets the value of the RemoteInt.
func (i RemoteInt) Set(value int) {
	i.i.setFor(i.u, strconv.Itoa(value))
}
