package client

import (
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
)

type Value = js.AnyValue

type Object = js.AnyObject

type String = js.AnyString

type Bool = js.AnyBool

type Int = js.AnyNumber

type Float = js.AnyNumber

type Function = js.AnyFunction

func NewString(from string) String {
	return js.NewString(from)
}

func NewInt(from int) Int {
	return js.NewNumber(float64(from))
}

func NewBool(from bool) Bool {
	return js.NewBool(from)
}

func NewFunction(from script.Script) Function {
	return js.NewFunction(from)
}
