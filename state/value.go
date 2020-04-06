package state

import (
	"encoding/base64"
	"fmt"
	"math/big"
	"strconv"

	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/user"
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
		u.Execute(fmt.Sprintf(`%v.setItem("%v", %v); seed.state["%[2]v"].changed();`, v.storage, v.key, strconv.Quote(value)))
	}
	u.Execute(fmt.Sprintf(`if (seed.state["%[1]v"]) seed.state["%[1]v"].changed();`, v.key))
}

//Set the value.
func (v Value) set(q script.Ctx, value script.String) {
	if !v.ro {
		q(fmt.Sprintf(`%v.setItem("%v", %v);`, v.storage, v.key, value))
	}
	q(fmt.Sprintf(`if (seed.state["%[1]v"]) seed.state["%[1]v"].changed();`, v.key))
}

//Get the value.
func (v Value) get() script.String {
	if v.raw != "" {
		return js.String{js.NewValue(v.raw)}
	}
	return js.String{v.getter()}
}

//Get the value.
func (v Value) getter() js.Value {
	return js.NewValue(fmt.Sprintf(`(%v.getItem("%v") || %v)`, v.storage, v.key, strconv.Quote(v.fallback)))
}
