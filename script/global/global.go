package global

import (
	"encoding/base64"
	"math/big"
	"strconv"

	"github.com/qlova/script/language"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/user"

	Javascript "github.com/qlova/script/language/javascript"
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
}

//Ref returns the raw reference.
func (ref Reference) Ref() string {
	return ref.string
}

//Set is a set method that should be called whenever the parent value is set.
func (ref Reference) Set(q script.Ctx) {
	q.Javascript(`if (dynamic["` + ref.string + `"]) dynamic["` + ref.string + `"]();`)
}

//SetFor is a set method that should be called whenever the parent value is set.
func (ref Reference) SetFor(u user.Ctx) {
	u.Execute(`if (dynamic["` + ref.string + `"]) dynamic["` + ref.string + `"]();`)
}

//New returns a new globl variable reference.
func New(name ...string) Reference {
	if len(name) > 0 {
		return Reference{"global_" + name[0]}
	}

	//global identification is compressed to base64 and prefixed with g_.
	var result = "g_" + base64.RawURLEncoding.EncodeToString(big.NewInt(id).Bytes())

	id++

	return Reference{result}
}

//String is a global Integer.
type String struct {
	Reference
}

//NewString returns a reference to a new global string.
func NewString(name ...string) String {
	return String{New(name...)}
}

//Get the script.String for the global.String
func (s String) Get(q script.Ctx) script.String {
	return q.StringFromLanguageType(Javascript.String{
		Expression: language.Statement(`(window.localStorage.getItem("` + s.string + `"))`),
	})
}

//Set the global.String to be script.String
func (s String) Set(q script.Ctx, value script.String) {
	q.Javascript(`window.localStorage.setItem("` + s.string + `", ` + value.LanguageType().Raw() + `);`)
	s.Reference.Set(q)
}

//SetFor the global.String to be  the given value.
func (s String) SetFor(u user.Ctx, value string) {
	u.Execute(`window.localStorage.setItem("` + s.string + `", ` + strconv.Quote(value) + `);`)
	s.Reference.SetFor(u)
}

//Bool is a global Boolean.
type Bool struct {
	Reference
	Expression string
}

//NewBool returns a reference to a new global boolean.
func NewBool(name ...string) Bool {
	return Bool{New(name...), ""}
}

//Get the script.Bool for the global.Bool
func (b Bool) Get(q script.Ctx) script.Bool {
	if b.Expression != "" {
		return q.Value(b.Expression).Bool()
	}
	return q.BoolFromLanguageType(Javascript.Bit{
		Expression: language.Statement(`(window.localStorage.getItem("` + b.string + `") == "true")`),
	})
}

//Set the global.Bool to be script.Bool
func (b Bool) Set(q script.Ctx, value script.Bool) {
	q.Javascript(`window.localStorage.setItem("` + b.string + `", ` + value.LanguageType().Raw() + `);`)
	b.Reference.Set(q)
}

//Array is a global Array.
type Array struct {
	Reference
}

//NewArray returns a reference to a new global array.
func NewArray(name ...string) Array {
	return Array{New(name...)}
}

//Get the script.Array for the global.Array
func (a Array) Get(q script.Ctx) script.Array {
	return q.Value(`JSON.parse(window.localStorage.getItem("` + a.string + `") || "[]")`).Array()
}

//Set the global.Array to be script.Array
func (a Array) Set(q script.Ctx, value script.Array) {
	q.Javascript(`window.localStorage.setItem("` + a.string + `", JSON.stringify(` + value.LanguageType().Raw() + `));`)
	a.Reference.Set(q)
}

//Object is a global Object.
type Object struct {
	Reference
}

//NewObject returns a reference to a new global object.
func NewObject(name ...string) Object {
	return Object{New(name...)}
}

//Get the script.Object for the global.Object
func (o Object) Get(q script.Ctx) script.Object {
	return q.Value(`JSON.parse(window.localStorage.getItem("` + o.string + `") || "{}")`).Object()
}

//Set the global.Object to be script.Object
func (o Object) Set(q script.Ctx, value script.Object) {
	q.Javascript(`window.localStorage.setItem("` + o.string + `", JSON.stringify(` + value.LanguageType().Raw() + `));`)
	o.Reference.Set(q)
}
