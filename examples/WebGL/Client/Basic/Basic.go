package main

import "github.com/qlova/seed"
import "github.com/qlova/seeds/canvas"

func main() {
	var App = seed.NewApp()

	var Canvas = canvas.AddTo(App)

	Canvas.OnReady(func(q seed.Script) {
		var gl = Canvas.Script(q).OpenGL()

		gl.ClearColor(q.Float(0), q.Float(1), q.Float(0), q.Float(1))
		gl.Clear(gl.ColorBufferBit)
	})

	App.Launch()
}
