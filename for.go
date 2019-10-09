package seed

import (
	"fmt"
	"strings"

	qlova "github.com/qlova/script"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/style"
	"github.com/qlova/seed/style/css"
	"github.com/qlova/seed/user"
)

func dashes2camels(s string) string {
	var camel string
	var parts = strings.Split(s, "-")
	for i, part := range parts {
		if i == 0 {
			camel += part
		} else {
			camel += strings.Title(part)
		}
	}
	return camel
}

//Go runs a Go function instead of a script Function.
func Go(function interface{}, args ...qlova.Type) func(q script.Ctx) {
	return func(q script.Ctx) {
		q.Go(function)
	}
}

//Update is a remote update to a seed.
type Update struct {
	style.Style

	id string
	user.User
}

//For returns a new remote update to a seed that can be used to remotely modify the seed.
func (seed Seed) For(u User) Update {
	var update Update
	update.id = seed.id
	update.User = u
	update.Style = style.From(update)
	return update
}

//CSS returns the css.Style to the seed's style.
func (update Update) CSS() css.Style {
	return update.Style.Style
}

//Set sets the CSS style property to value.
func (update Update) Set(property, value string) {
	property = dashes2camels(property)

	fmt.Println(property)

	update.Document["#"+update.id+".style."+property] = value
}

//Get returns the CSS property of this seed.
func (update Update) Get(property string) string {
	property = dashes2camels(property)

	return update.Document["#"+update.id+".style."+property]
}

//SetText sets the seed's text.
func (update Update) SetText(text string) {
	update.Document["#"+update.id+".innerText"] = text
}
