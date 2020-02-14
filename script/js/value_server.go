//+build !wasm

package js

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/qlova/seed"
	"github.com/qlova/seed/user"
)

//Value is a generic js.Value.
type Value struct {
	string
	writer *bytes.Buffer

	u user.Ctx
}

//ValueFromUpdate returns a value from a seed.Update
func ValueFromUpdate(u seed.Update) Value {
	return Value{
		string: `document.getElementById("` + u.ID() + `")`,
		writer: bytes.NewBuffer(nil),
	}
}

func (v Value) Global() Value {
	return Value{
		string: "window",
		writer: v.writer,
	}
}

func (v Value) String(s string) Value {
	return Value{
		string: strconv.Quote(s),
		writer: v.writer,
	}
}

func (v Value) Number(n float64) Value {
	return Value{
		string: fmt.Sprint(n),
		writer: v.writer,
	}
}

func (v Value) Run(method string, args ...Value) {
	fmt.Fprintf(v.writer, `%v[%v](`, v.string, strconv.Quote(method))
	for i, arg := range args {
		fmt.Fprintf(v.writer, "%v", arg.string)
		if i < len(args)-1 {
			v.writer.WriteByte(',')
		}
	}
	fmt.Fprintf(v.writer, ");")
	v.u.Execute(v.writer.String())
}

func (v Value) Call(method string, args ...Value) Value {
	var buffer bytes.Buffer

	fmt.Fprintf(&buffer, `%v[%v](`, v.string, strconv.Quote(method))
	for i, arg := range args {
		fmt.Fprintf(&buffer, "%v", arg.string)
		if i < len(args)-1 {
			buffer.WriteByte(',')
		}
	}
	fmt.Fprintf(&buffer, ")")

	return Value{buffer.String(), v.writer, v.u}
}

func (v Value) Set(property string, value Value) {
	fmt.Fprintf(v.writer, "%v[%v] = %v;", v.string, strconv.Quote(property), value.string)
	v.u.Execute(v.writer.String())
}

func (v Value) Get(property string) Value {
	return Value{
		string: fmt.Sprintf("%v[%v]", v.string, strconv.Quote(property)),
		writer: v.writer,
	}
}

func (v Value) Var(name string) Value {
	fmt.Fprintf(v.writer, "let %v = %v;", name, v.string)
	return Value{
		name, v.writer, v.u,
	}
}
