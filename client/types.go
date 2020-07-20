package client

import (
	"qlova.org/seed/js"
	"qlova.org/seed/script"
)

type Script = js.AnyScript

func NewScript(do ...Script) Script {
	var s js.Script
	for _, scriptable := range do {
		if scriptable == nil {
			continue
		}
		s = s.Append(scriptable.GetScript())
	}
	if s == nil {
		return js.Script(func(q js.Ctx) {})
	}
	return s
}

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

func NewFloat64(from float64) Float {
	return js.NewNumber(from)
}

func NewFunction(from script.Script) Function {
	return js.NewFunction(from)
}
