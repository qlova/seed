package global

import (
	"encoding/base64"
	"math/big"

	"github.com/qlova/script/language"
	"github.com/qlova/seed/script"

	Javascript "github.com/qlova/script/language/javascript"
)

//All globals have a unique id.
var id int64 = 1

//Reference is a global variable reference.
type Reference struct {
	string
}

func (ref Reference) String() string {
	return ref.string
}

//Set is a set method that should be called whenever the parent value is set.
func (ref Reference) Set(q script.Ctx) {
	q.Javascript(`if (dynamic["` + ref.string + `"]) dynamic["` + ref.string + `"]();`)
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

//Int is a global Integer.
type Int Reference

//NewInt returns a reference to a new global integer.
func NewInt(name ...string) Int {
	return Int(New(name...))
}

//Get the script.Int for the global.Int
func (i Int) Get(q script.Ctx) script.Int {
	return q.IntFromLanguageType(Javascript.Integer{
		Expression: language.Statement(`(parseInt(window.localStorage.getItem("` + i.string + `") || "0"))`),
	})
}

//Set the global.Int to be script.Int
func (i Int) Set(q script.Ctx, value script.Int) {
	q.Javascript(`window.localStorage.setItem("` + i.string + `", ` + value.LanguageType().Raw() + `.toString());`)
	Reference(i).Set(q)
}

//String is a global Integer.
type String Reference

//NewString returns a reference to a new global string.
func NewString(name ...string) String {
	return String(New(name...))
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
	Reference(s).Set(q)
}

//Bool is a global Boolean.
type Bool Reference

//NewBool returns a reference to a new global boolean.
func NewBool(name ...string) Bool {
	return Bool(New(name...))
}

//Get the script.Bool for the global.Bool
func (b Bool) Get(q script.Ctx) script.Bool {
	return q.BoolFromLanguageType(Javascript.Bit{
		Expression: language.Statement(`(window.localStorage.getItem("` + b.string + `") == "true")`),
	})
}

//Set the global.Bool to be script.Bool
func (b Bool) Set(q script.Ctx, value script.Bool) {
	q.Javascript(`window.localStorage.setItem("` + b.string + `", ` + value.LanguageType().Raw() + `);`)
	Reference(b).Set(q)
}

//Array is a global Array.
type Array Reference

//NewArray returns a reference to a new global array.
func NewArray(name ...string) Array {
	return Array(New(name...))
}

//Get the script.Array for the global.Array
func (a Array) Get(q script.Ctx) script.Array {
	return q.Value(`JSON.parse(window.localStorage.getItem("` + a.string + `") || "[]")`).Array()
}

//Set the global.Array to be script.Array
func (a Array) Set(q script.Ctx, value script.Array) {
	q.Javascript(`window.localStorage.setItem("` + a.string + `", JSON.stringify(` + value.LanguageType().Raw() + `));`)
	Reference(a).Set(q)
}

//Object is a global Object.
type Object Reference

//NewObject returns a reference to a new global object.
func NewObject(name ...string) Object {
	return Object(New(name...))
}

//Get the script.Object for the global.Object
func (o Object) Get(q script.Ctx) script.Object {
	return q.Value(`JSON.parse(window.localStorage.getItem("` + o.string + `") || "{}")`).Object()
}

//Set the global.Object to be script.Object
func (o Object) Set(q script.Ctx, value script.Object) {
	q.Javascript(`window.localStorage.setItem("` + o.string + `", JSON.stringify(` + value.LanguageType().Raw() + `));`)
	Reference(o).Set(q)
}
