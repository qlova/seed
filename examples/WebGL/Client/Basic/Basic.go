//+build ignore

package main

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	"github.com/qlova/seeds/canvas"
)

func main() {
	var App = seed.NewApp()

	var Canvas = canvas.AddTo(App)

	Canvas.OnReady(func(q script.Ctx) {
		var gl = Canvas.Ctx(q).OpenGL()

		gl.ClearColor(q.Float(0), q.Float(1), q.Float(0), q.Float(1))
		gl.Clear(gl.ColorBufferBit)
	})

	App.Launch()
}
