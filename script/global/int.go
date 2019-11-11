package global

import (
	"github.com/qlova/seed/script"
)

//Int is a global Integer.
type Int struct {
	Reference
}

//NewInt returns a reference to a new global integer.
func NewInt(name ...string) Int {
	return Int{New(name...)}
}

//Get the script.Int for the global.Int
func (i Int) Get(q script.Ctx) script.Int {
	return q.Value(`(parseInt(window.localStorage.getItem("` + i.string + `") || "0"))`).Int()
}

//Set the global.Int to be script.Int
func (i Int) Set(q script.Ctx, value script.Int) {
	q.Javascript(`window.localStorage.setItem("` + i.string + `", (` + value.LanguageType().Raw() + `).toString());`)
	i.Reference.Set(q)
}

//PlusPlus increments the integer by one.
func (i Int) PlusPlus(q script.Ctx) {
	q.Javascript(`window.localStorage.setItem("%v", (+(window.localStorage.getItem("%v") || "0") + 1).toString());`,
		i.string, i.string)
	i.Reference.Set(q)
}
