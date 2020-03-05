package state

import (
	"encoding/base64"
	"math/big"

	"github.com/qlova/script/language"
	"github.com/qlova/seed/script"
)

//Variable is any global variable.
type Variable interface {
	Ref() string
}

//All globals have a unique id.
var id int64 = 1

//Reference is a global variable reference.
type Reference struct {
	string
	initial string
}

//Ref returns the raw reference.
func (ref Reference) Ref() string {
	return ref.string
}

//Set is a set method that should be called whenever the parent value is set.
func (ref Reference) Set(q script.Ctx) {
	q.Javascript(`if (seed.dynamic["` + ref.string + `"]) seed.dynamic["` + ref.string + `"]();`)
}

//NewVariable returns a new globl variable reference.
func NewVariable(initial string) Reference {

	//global identification is compressed to base64 and prefixed with g_.
	var result = "g_" + base64.RawURLEncoding.EncodeToString(big.NewInt(id).Bytes())

	id++

	return Reference{result, initial}
}

//Bool is a global Boolean.
type Bool struct {
	Reference
	Expression string
}

//NewBool returns a reference to a new global boolean.
func NewBool() Bool {
	return Bool{NewVariable("false"), ""}
}

func (b Bool) BoolFromCtx(q script.AnyCtx) script.Bool {
	return b.Get(script.CtxFrom(q))
}

//Get the script.Bool for the global.Bool
func (b Bool) Get(q script.Ctx) script.Bool {
	if b.Expression != "" {
		return q.Value(b.Expression).Bool()
	}
	return script.Bool{language.Expression(q, `(window.localStorage.getItem("`+b.string+`") == "true")`)}
}

//Set the global.Bool to be script.Bool
func (b Bool) Set(q script.Ctx, value script.Bool) {
	q.Javascript(`window.localStorage.setItem("` + b.string + `", ` + q.Raw(value) + `);`)
	b.Reference.Set(q)
}
