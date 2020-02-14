//+build ignore

package main

import (
	"github.com/qlova/seed"
	"github.com/qlova/seeds/canvas"
)

func main() {
	var App = seed.NewApp()

	var Canvas = canvas.AddTo(App)

	var gl = Canvas.OpenGL()
	go func() {
		gl.ClearColor(0, 1, 0, 1)
		gl.Clear(gl.ColorBufferBit)
	}()

	App.Launch()
}
