package state

import (
	"fmt"
	"strconv"

	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/user"
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
		i.set(q, i.get(q).Plus(q.Int(1)))
	}
}

//IntFromCtx implements script.AnyInt
func (i Int) IntFromCtx(q script.AnyCtx) script.Int {
	return i.get(script.CtxFrom(q))
}

//ValueFromCtx implements script.AnyValue
func (i Int) ValueFromCtx(q script.AnyCtx) script.Value {
	return i.get(script.CtxFrom(q))
}

func (i Int) get(q script.Ctx) script.Int {
	return q.Value(`parseInt(%v)`, i.Value.get(q)).Int()
}

//SetL sets the value of the Int with a literal.
func (i Int) SetL(value int) script.Script {
	return func(q script.Ctx) {
		i.set(q, q.Int(value))
	}
}

func (i Int) set(q script.Ctx, value script.Int) {
	i.Value.set(q, q.Value(`(%v).toString();`).String())
}

//SetText sets the seed's text to reflect the value of this Int.
func (i Int) SetText() seed.Option {
	return seed.NewOption(func(any seed.Any) {
		if i.raw == "" {
			data := seeds[any.Root()]

			if data.change == nil {
				data.change = make(map[Value]script.Script)
			}

			data.change[i.Value] = data.change[i.Value].Then(func(q script.Ctx) {
				fmt.Fprintf(q, `%v.innerText = (%v).toString();`, any.Root().Ctx(q).Element(), q.Raw(i.get(q)))
			})
			seeds[any.Root()] = data
		} else {
			any.Add(script.OnReady(func(q script.Ctx) {
				fmt.Fprintf(q, `%v.innerText = (%v).toString();`, any.Root().Ctx(q).Element(), q.Raw(i.get(q)))
			}))
		}
	}, func(seed seed.Ctx) {
		panic(".Var seeds not allowed in conditional")
	}, func(seed seed.Ctx) {
		panic(".Var seeds not allowed in conditional")
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
