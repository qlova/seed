package state

import (
	"encoding/base64"
	"fmt"
	"math/big"
	"strconv"

	"qlova.org/seed"
	"qlova.org/seed/js"
	"qlova.org/seed/script"
	"qlova.org/seed/user"
)

//Option can be used to configure values.
type Option func(*Value)

type AnyValue interface {
	value() Value
}

//Value is the backing of any state value.
type Value struct {
	key, raw string
	fallback string
	storage  string

	ro bool

	dependencies *[]Value
}

//Get the value.
func (v Value) Null() bool {
	if v.raw == "" && v.key == "" && v.storage == "" {
		return true
	}
	return false
}

//Raw returns a raw value from a JS expression.
func Raw(expr string, options ...Option) Value {
	id++
	v := Value{
		key: "state." + base64.RawURLEncoding.EncodeToString(big.NewInt(id).Bytes()),
		ro:  true,
		raw: expr,
	}
	for _, o := range options {
		o(&v)
	}
	return v
}

var id int64

//NewValue returns a new state value.
func newValue(fallback string, options ...Option) Value {
	id++
	v := Value{
		key:      "state." + base64.RawURLEncoding.EncodeToString(big.NewInt(id).Bytes()),
		fallback: fallback,
		storage:  "localStorage",
	}
	for _, o := range options {
		o(&v)
	}
	return v
}

func (v Value) value() Value {
	return v
}

func (v Value) setFor(u user.Ctx, value string) {
	if !v.ro {
		var f = js.Function{js.NewValue(v.storage + `.setItem`)}
		u.Execute(js.Run(f, js.NewString(v.key), js.NewString(value)))
	}
	u.Execute(func(q js.Ctx) {
		q(fmt.Sprintf(`if (seed.state["%[1]v"]) await seed.state["%[1]v"].changed(q);`, v.key))
	})
}

//Set the value.
func (v Value) set(q script.Ctx, value script.String) {
	if !v.ro {
		q(fmt.Sprintf(`%v.setItem("%v", %v);`, v.storage, v.key, value))
	}
	q(fmt.Sprintf(`if (seed.state["%[1]v"]) await seed.state["%[1]v"].changed(q);`, v.key))
}

//Get the value.
func (v Value) get() script.String {
	return js.String{v.getter()}
}

//Get the value.
func (v Value) getter() js.Value {
	if v.raw != "" {
		return js.NewValue(v.raw)
	}
	return js.NewValue(fmt.Sprintf(`(%v.getItem("%v") || %v)`, v.storage, v.key, strconv.Quote(v.fallback)))
}

func SetProperty(property string, value AnyValue) seed.Option {
	var s = value.value()
	return seed.NewOption(func(c seed.Seed) {
		switch c.(type) {
		case script.Seed, script.Undo:
			panic("state.String.SetText must not be called on a script.Seed")
		}

		OnRefresh(func(q script.Ctx) {
			fmt.Fprintf(q, `%[1]v.`+property+` = %[2]v;`, script.Scope(c, q).Element(), s.get())
		}).AddTo(c)

		if s.key != "" {
			var data data
			c.Read(&data)

			if data.change == nil {
				data.change = make(map[Value]script.Script)
			}

			data.change[s] = data.change[s].Append(Refresh(c))

			if s.dependencies != nil {
				for _, dep := range *s.dependencies {
					data.change[dep] = data.change[dep].Append(Refresh(c))
				}
			}

			c.Write(data)
		}
	})
}
