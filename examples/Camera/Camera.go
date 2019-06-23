package main

import "github.com/qlova/seed"
import "github.com/qlova/seeds/camera"
import "github.com/qlova/seeds/button"
import "github.com/qlova/seeds/image"

func main() {
	var App = seed.NewApp("Camera")
	App.SetSize(100, 100)
	App.SetColor(seed.RGB(0, 0, 0))

	var Camera = camera.AddTo(App)
	Camera.SetSize(100, 100)
	
	var LastSnapshot = image.AddTo(App)
		LastSnapshot.SetHidden()
		LastSnapshot.SetAttach(seed.Bottom+seed.Right)
		LastSnapshot.SetHeight(10)
	
	var Button = button.AddTo(App, "Take Picture")
		Button.SetAttach(seed.Bottom+seed.Left)
		Button.OnClick(func(q seed.Script) {
			var Camera = Camera.Script(q)
			var LastSnapshot = LastSnapshot.Script(q)

			LastSnapshot.SetSource(Camera.Capture().Source())
			LastSnapshot.SetVisible()
		})

	App.Launch()
}
