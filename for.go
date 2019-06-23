package seed

import "github.com/qlova/seed/user"
import qlova "github.com/qlova/script"

func Go(function interface{}, args ...qlova.Type) func(q Script) {
	return func(q Script) {
		q.Go(function)
	}
}

type Update struct {
	id string
	user.User
}

func (seed Seed) For(u User) Update {
	return Update{seed.id, u}
}

func (update Update) SetText(text string) {
	update.Document["#"+update.id+".innerText"] = text
}
