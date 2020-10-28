package js

import (
	"fmt"
	"io"
	"reflect"

	"qlova.org/seed"
)

type AnyScript interface {
	AnyFunction
	GetScript() Script
}

//Script is any js script.
type Script func(Ctx)

func (s Script) GetBool() Bool {
	return s.GetValue().GetBool()
}

func (s Script) GetValue() Value {
	return s.GetFunction().Value
}

func (s Script) GetScript() Script {
	return s
}

func (s Script) GetFunction() Function {
	return NewFunction(s)
}

func Append(a, b AnyScript) Script {
	if a == nil {
		if b == nil {
			return func(Ctx) {}
		}
		return b.GetScript()
	}
	if b == nil {
		if a == nil {
			return func(Ctx) {}
		}
		return a.GetScript()
	}
	return func(c Ctx) {
		a.GetScript()(c)
		b.GetScript()(c)
	}
}

//Append appends two scripts to return a new script.
func (s Script) Append(next Script) Script {
	if s == nil {
		return next
	}
	if next == nil {
		return s
	}
	return func(c Ctx) {
		s(c)
		next(c)
	}
}

//Ctx is a script context.
type Ctx func(interface{})

//NewCtx returns a new ctx that writes a script to the given writer.
//This function takes an optional number of seeds, any options passed to this context will be applied to the given seeds.
func NewCtx(w io.Writer, seeds ...seed.Seed) Ctx {
	var ctx func(in interface{})

	mw := newMacroWriter(w, seeds...)

	ctx = func(in interface{}) {
		switch arg := in.(type) {
		case Script:
			if arg != nil {
				arg(ctx)
			}
		case AnyScript:
			if arg != nil {
				arg.GetScript()(ctx)
			}
		case func(Ctx):
			if arg != nil {
				arg(ctx)
			}
		case func(Ctx) Value:
			if arg != nil {
				arg(ctx)
			}
		case error:
			mw.Flush()
		case rune:
			fmt.Fprint(mw, string(arg))
		case string:
			fmt.Fprint(mw, arg)
		case []byte:
			mw.Write(arg)
		case AnyValue:
			fmt.Fprint(mw, arg.GetValue())
		case seed.Option:
			for _, c := range seeds {
				arg.AddTo(c)
			}
		default:
			panic("invalid type: " + reflect.TypeOf(in).String())

		}
	}

	return ctx
}

func (q Ctx) Write(b []byte) (int, error) {
	q(b)
	return len(b), nil
}

func (q Ctx) Flush() {
	q(io.EOF)
}

var unique int

//Unique returns a unique string suitable for variable names.
func (Ctx) Unique() string {
	unique++
	return fmt.Sprint("unique_", unique)
}
