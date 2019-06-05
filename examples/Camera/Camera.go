package main

import "github.com/qlova/seed"
import "github.com/qlova/seed/widgets/camera"

func main() {
	var App = seed.NewApp("Camera")
	App.SetSize(100, 100)
	App.SetColor(seed.RGB(0, 0, 0))

	var Camera = camera.AddTo(App)
	Camera.SetSize(100, 100)

	App.Launch()
}
