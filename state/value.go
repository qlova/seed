package state

import (
	"encoding/base64"
	"fmt"
	"math/big"
	"strconv"

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
		q.Javascript(`%v.setItem("%v", %v);`, v.storage, v.key, q.Raw(value))
	}
	q.Javascript(`if (seed.state["%[1]v"]) seed.state["%[1]v"].changed();`, v.key)
}

//Get the value.
func (v Value) get(q script.Ctx) script.String {
	if v.raw != "" {
		return q.Value(v.raw).String()
	}
	return q.Value(v.getter()).String()
}

//Get the value.
func (v Value) getter() string {
	return fmt.Sprintf(`(%v.getItem("%v") || %v)`, v.storage, v.key, v.fallback)
}
