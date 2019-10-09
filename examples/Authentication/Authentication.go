package main

import (
	"github.com/qlova/seed/style/css"
	"github.com/qlova/seed/user"
	"github.com/qlova/seed/user/password"

	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	"github.com/qlova/seeds/passwordbox"
	"github.com/qlova/seeds/text"
	"github.com/qlova/seeds/textbox"
)

//User is a dummy user account.
type User struct {
	Name, Password string
}

//Users is a dummy database of users.
var Users = make(map[string]User)

func main() {
	var App = seed.NewApp("Authentication")
	App.ItemSpacing().Center()

	var UserNameBox = textbox.AddTo(App)
	var PasswordBox = passwordbox.AddTo(App)
	var Message = text.AddTo(App, "\n") //We use a newline to keep the spacing of page when a message arrives.
	PasswordBox.OnEnter(func(q script.Ctx) {
		Message.Ctx(q).SetText(q.String("\n")) //Reset the message.

		PasswordBox.Ctx(q).With(script.Args{
			"username": UserNameBox.Ctx(q).Value(),
			//We need to send a username along with the hash, this could be an email or phone number.
		}).HashAndGo(func(user user.User, hash password.Hash) {

			var username = user.Args("username").String()

			//Check if the user already exists and try logging them in.
			if record, ok := Users[username]; ok || username == "" {
				if !hash.VerifyFor(username, record.Password) {
					Message.For(user).SetText("username taken and/or invalid username and password.")
					Message.For(user).CSS().SetColor(css.Red)
					return
				}
				Message.For(user).SetText("Authenticated!")
				Message.For(user).CSS().SetColor(css.Green)
				return
			}

			//This is a password that you can safely store in your database.
			var password, err = hash.PasswordFor(username)
			if err != nil {
				Message.For(user).SetText(err.Error())
				Message.For(user).CSS().SetColor(css.Red)
				return
			}

			//Add this user to the dummy database.
			Users[username] = User{
				Name:     username,
				Password: password,
			}
			Message.For(user).SetText("You have been registered.")
			Message.For(user).CSS().SetColor(css.Black)
		})
	})

	App.Launch()
}
