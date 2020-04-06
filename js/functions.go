package js

import "strings"

//Call calls the given function expressionw with the given arguments.
func Call(fname string, args ...AnyValue) Value {
	var b strings.Builder

	b.WriteString(fname)
	b.WriteRune('(')
	for i, arg := range args {
		b.WriteString(arg.GetValue().String())
		if i < len(args)-1 {
			b.WriteRune(',')
		}
	}
	b.WriteRune(')')

	return Value{b.String()}
}

//Run runs the given function expression with the given arguments.
func (q Ctx) Run(fname string, args ...AnyValue) {
	q(fname)
	q('(')
	for i, arg := range args {
		q(arg.GetValue().String())
		if i < len(args)-1 {
			q(',')
		}
	}
	q(')')
}
