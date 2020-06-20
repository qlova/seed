package client

import (
	"qlova.org/seed/js"
	"qlova.org/seed/script"
)

type Object = js.AnyObject

type Bool = js.AnyBool

type Int = js.AnyNumber

type Float = js.AnyNumber

type Function = js.AnyFunction

func NewInt(from int) Int {
	return js.NewNumber(float64(from))
}

func NewBool(from bool) Bool {
	return js.NewBool(from)
}

func NewFunction(from script.Script) Function {
	return js.NewFunction(from)
}
