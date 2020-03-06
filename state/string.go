package state

import (
	"fmt"
	"strconv"

	"github.com/qlova/script/language"
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/user"
)

//String is a global String.
type String struct {
	Reference
	Expression string

	dependencies []Reference
}

//NewString returns a reference to a new global string.
func NewString(initial string) String {
	return String{NewVariable(strconv.Quote(initial)), "", nil}
}

func (s String) StringFromCtx(q script.AnyCtx) script.String {
	return s.get(script.CtxFrom(q))
}

func (s String) ValueFromCtx(q script.AnyCtx) script.Value {
	return s.get(script.CtxFrom(q))
}

//Get the script.Bool for the global.String
func (s String) get(q script.Ctx) script.String {
	if s.Expression != "" {
		return q.Value(s.Expression).String()
	}
	return script.String{language.Expression(q, `(window.localStorage.getItem("`+s.string+`") || `+s.initial+`)`)}
}

//SetL allows setting the value of a String to a literal in the given script ctx.
func (s String) SetL(literal string) script.Script {
	return func(q script.Ctx) {
		s.set(q, q.String(literal))
	}
}

//Set the global.Bool to be script.Bool
func (s String) set(q script.Ctx, value script.String) {
	q.Javascript(`window.localStorage.setItem("` + s.string + `", ` + q.Raw(value) + `); seed.state["` + s.string + `"].changed();`)
	s.Reference.Set(q)
}

type RemoteString struct {
	u user.Ctx
	s String
}

func (s String) For(u user.Ctx) RemoteString {
	return RemoteString{u, s}
}

func (s RemoteString) Set(value string) {
	s.u.Execute(fmt.Sprintf(`window.localStorage.setItem("%v", %v); seed.state["%[1]v"].changed();`, s.s.string, strconv.Quote(value)))
}

func (s String) SetText() seed.Option {
	return seed.NewOption(func(any seed.Any) {
		if s.Expression == "" {
			any.Add(script.OnReady(func(q script.Ctx) {
				fmt.Fprintf(q, `%v.innerText = %v;`, any.Root().Ctx(q).Element(), q.Raw(s.get(q)))
			}))
		}

		if s.Ref() != "" {
			data := seeds[any.Root()]

			if data.change == nil {
				data.change = make(map[Reference]script.Script)
			}

			data.change[s.Reference] = data.change[s.Reference].Then(func(q script.Ctx) {
				fmt.Fprintf(q, `%v.innerText = %v;`, any.Root().Ctx(q).Element(), q.Raw(s.get(q)))
			})

			for _, dep := range s.dependencies {
				data.change[dep.GetReference()] = data.change[dep.GetReference()].Then(func(q script.Ctx) {
					fmt.Fprintf(q, `seed.state["%v"].changed();`, s.Ref())
				})
			}

			seeds[any.Root()] = data
		}
	}, func(seed seed.Ctx) {
		panic(".Var seeds not allowed in conditional")
	}, func(seed seed.Ctx) {
		panic(".Var seeds not allowed in conditional")
	})
}
