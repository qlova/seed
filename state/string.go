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
}

//NewString returns a reference to a new global string.
func NewString(initial string) String {
	return String{NewVariable(strconv.Quote(initial)), ""}
}

func (s String) StringFromCtx(q script.AnyCtx) script.String {
	return s.Get(script.CtxFrom(q))
}

func (s String) ValueFromCtx(q script.AnyCtx) script.Value {
	return s.Get(script.CtxFrom(q))
}

//Get the script.Bool for the global.String
func (s String) Get(q script.Ctx) script.String {
	if s.Expression != "" {
		return q.Value(s.Expression).String()
	}
	return script.String{language.Expression(q, `(window.localStorage.getItem("`+s.string+`") || `+s.initial+`)`)}
}

//Set the global.Bool to be script.Bool
func (s String) Set(q script.Ctx, value script.String) {
	q.Javascript(`window.localStorage.setItem("` + s.string + `", ` + q.Raw(value) + `);`)
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
			data := seeds[any.Root()]

			if data.change == nil {
				data.change = make(map[Reference]script.Script)
			}

			data.change[s.Reference] = data.change[s.Reference].Then(func(q script.Ctx) {
				fmt.Fprintf(q, `%v.innerText = %v;`, any.Root().Ctx(q).Element(), q.Raw(s.Get(q)))
			})
			seeds[any.Root()] = data
		} else {
			any.Add(script.OnReady(func(q script.Ctx) {
				fmt.Fprintf(q, `%v.innerText = %v;`, any.Root().Ctx(q).Element(), q.Raw(s.Get(q)))
			}))
		}
	}, func(seed seed.Ctx) {
		panic(".Var seeds not allowed in conditional")
	}, func(seed seed.Ctx) {
		panic(".Var seeds not allowed in conditional")
	})
}
