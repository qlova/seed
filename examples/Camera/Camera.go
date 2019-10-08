package main

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	"github.com/qlova/seeds/button"
	"github.com/qlova/seeds/camera"
	"github.com/qlova/seeds/image"
)

func main() {
	var App = seed.NewApp("Camera")
	App.SetSize(100, 100)
	App.SetColor(seed.RGB(0, 0, 0))

	var Camera = camera.AddTo(App)
	Camera.SetSize(100, 100)

	var LastSnapshot = image.AddTo(App)
	LastSnapshot.SetHidden()
	LastSnapshot.AttachToScreen().Bottom().Right()
	LastSnapshot.SetHeight(10)

	var Button = button.AddTo(App, "Take Picture")
	Button.AttachToScreen().Bottom().Left()
	Button.OnClick(func(q script.Ctx) {
		var Camera = Camera.Ctx(q)
		var LastSnapshot = LastSnapshot.Ctx(q)

		LastSnapshot.SetSource(Camera.Capture().Source())
		LastSnapshot.SetVisible()
	})

	App.Launch()
}
