package js

import (
	"fmt"
	"io"
	"reflect"

	"github.com/qlova/seed"
)

//Script is any js script.
type Script func(Ctx)

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

	ctx = func(in interface{}) {
		switch arg := in.(type) {
		case Script:
			arg(ctx)
		case func(Ctx):
			arg(ctx)
		case func(Ctx) Value:
			arg(ctx)
		case rune:
			fmt.Fprint(w, string(arg))
		case string:
			fmt.Fprint(w, arg)
		case []byte:
			w.Write(arg)
		case AnyValue:
			fmt.Fprint(w, arg.GetValue())
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

var unique int

//Unique returns a unique string suitable for variable names.
func (Ctx) Unique() string {
	unique++
	return fmt.Sprint("unique_", unique)
}
