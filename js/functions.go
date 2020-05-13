package js

import (
	"strings"
)

type Function struct {
	Value
}

type AnyFunction interface {
	AnyValue
	GetFunction() Function
}

func NewFunction(do Script) Function {
	var s strings.Builder

	s.WriteString(`async function() {`)

	NewCtx(&s)(do)

	s.WriteString(`}`)

	return Function{
		Value: NewValue(s.String()),
	}
}

//GetFunction impliments AnyFunction.
func (f Function) GetFunction() Function {
	return f
}

func (f Function) Call() Value {
	return NewValue("await " + f.string + "()")
}

func (q Ctx) Await(v AnyValue) {
	var val = v.GetValue()
	q("await " + val.string + ";")

}

func Await(v AnyValue) Value {
	var val = v.GetValue()
	val.string = "await " + val.string
	return val
}

//Call calls the given function expressionw with the given arguments.
func Call(fname AnyFunction, args ...AnyValue) Value {
	var b strings.Builder

	b.WriteString(fname.GetFunction().String())
	b.WriteRune('(')
	for i, arg := range args {
		b.WriteString(arg.GetValue().String())
		if i < len(args)-1 {
			b.WriteRune(',')
		}
	}
	b.WriteRune(')')

	return NewValue(b.String())
}

//Run runs the given function expression with the given arguments.
func Run(fname AnyFunction, args ...AnyValue) Script {
	return func(q Ctx) {
		q.Run(fname, args...)
	}
}

//Run runs the given function expression with the given arguments.
func (q Ctx) Run(fname AnyFunction, args ...AnyValue) {
	q(fname.GetFunction().String())
	q('(')
	for i, arg := range args {
		if arg == nil {
			q("null")
		} else {
			q(arg.GetValue().String())
		}
		if i < len(args)-1 {
			q(',')
		}
	}
	q(");")
}

//Return returns the value from the current function.
func Return(v AnyValue) func(q Ctx) {
	return func(q Ctx) {
		q("return ")
		if v != nil {
			q(v.GetValue().String())
		}
		q(";")
	}
}

//Return returns the value from the current function.
func (q Ctx) Return(v AnyValue) {
	q("return ")
	q(v.GetValue().String())
	q(";")
}
