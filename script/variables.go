package script

import (
	//Global ids.
	"encoding/base64"
	"math/big"
)

import qlova "github.com/qlova/script"
import "github.com/qlova/script/language"
import Javascript "github.com/qlova/script/language/javascript"

type setable interface {
	raw() string
	variable()
}

type Variable string

//All globals have a unique id.
var global_id int64 = 1

func NewVariable() Variable {
	//global identification is compressed to base64 and prefixed with g_.
	var result = "g_" + base64.RawURLEncoding.EncodeToString(big.NewInt(global_id).Bytes())

	global_id++

	return Variable(result)
}

func (Variable) variable() {}
func (v Variable) raw() string {
	return string(v)
}

func (q Script) Get(name setable) qlova.String {
	return q.wrap(`window.localStorage.getItem("` + string(name.raw()) + `")`)
}

func (q Script) Set(name setable, value qlova.Type) {
	var v = value.LanguageType()

	var s = value.LanguageType().Raw()

	switch v.(type) {
	case Javascript.String:

	case Javascript.Integer, Javascript.Bit:
		s = s + ".toString()"

	default:
		panic("Unimplemented")
	}

	q.Javascript(`window.localStorage.setItem("` + name.raw() + `", ` + s + `);`)
	q.Javascript(`if (dynamic["` + name.raw() + `"]) dynamic["` + name.raw() + `"](` + s + `);`)
}

type GlobalInt struct {
	Variable
}

func NewInt() GlobalInt {
	return GlobalInt{NewVariable()}
}

func (i GlobalInt) Script(q Script) Int {
	return q.IntFromLanguageType(Javascript.Integer{
		Expression: language.Statement(`(parseInt(window.localStorage.getItem("` + string(i.Variable) + `") || "0"))`),
	})
}

type BoolVar struct {
	Variable
}

func NewBool() BoolVar {
	return BoolVar{NewVariable()}
}

func (b BoolVar) Script(q Script) qlova.Bool {
	var result = q.BoolFromLanguageType(Javascript.Bit{
		Expression: language.Statement(`(window.localStorage.getItem("` + string(b.Variable) + `") == "true")`),
	})

	/*var extension = result.Extend()
	extension.Setter = func(value language.Type) language.Statement {
		return language.Statement(`window.localStorage.setItem("`+name.raw()+`", (`+string(t.Expression)+`).toString());`)
	}*/

	return result
}

type GlobalString struct {
	Variable
}

func NewString() GlobalString {
	return GlobalString{NewVariable()}
}

func (s GlobalString) Script(q Script) qlova.String {
	var result = q.StringFromLanguageType(Javascript.String{
		Expression: language.Statement(`window.localStorage.getItem("` + string(s.Variable) + `")`),
	})
	return result
}
