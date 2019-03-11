package main

import "github.com/qlova/seed"
import "github.com/qlova/seed/widgets/button"

func main() {
	var App = seed.NewApp("Dialogues")
	
	var Alert = button.AddTo(App, "Alert")
		Alert.OnClick(func(q seed.Script) {
			q.Alert(q.String("This is an alert!"))
		})
		
	var Confirm = button.AddTo(App, "Confirm")
		Confirm.OnClick(func(q seed.Script) {
			q.If(q.Confirm(q.String("This is a confirmation box!")), func() {
				q.Alert(q.String("You selected true!"))
			}, q.Else(func() {
				q.Alert(q.String("You selected false!"))
			}))
		})
	
	var Prompt = button.AddTo(App, "Prompt")
		Prompt.OnClick(func(q seed.Script) {
			q.Alert(q.String("You said ").Add(q.Prompt(q.String("This is a prompt!"))))
		})
		
	App.Launch()
}
