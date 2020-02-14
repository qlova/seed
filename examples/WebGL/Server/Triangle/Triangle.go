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
		var VertexShader = gl.CreateShader(gl.VertexShader)
		VertexShader.SetSource(`attribute vec3 c;void main(void){gl_Position=vec4(c, 1.0);}`)
		VertexShader.Compile()

		var FragmentShader = gl.CreateShader(gl.FragmentShader)
		FragmentShader.SetSource(`void main(void){gl_FragColor=vec4(0,1,1,1);}`)
		FragmentShader.Compile()

		var Program = gl.CreateProgram()
		Program.Attach(VertexShader)
		Program.Attach(FragmentShader)
		Program.Link()
		Program.Use()

		var VertexBuffer = gl.CreateBuffer()
		var c = Program.Attribute("c")

		gl.ClearColor(1, 0, 1, 1)
		gl.Clear(gl.ColorBufferBit)

		VertexBuffer.Bind(gl.ArrayBuffer)

		gl.BufferData(gl.ArrayBuffer, []float32{
			-0.5, 0.5, 0,
			-0.5, -0.5, 0,
			0.5, -0.5, 0,
		}, gl.StaticDraw)

		c.Load(3, gl.Float, false, 0, 0)
		c.Enable()

		gl.DrawArrays(gl.Triangles, 0, 3)
	}()

	App.Launch()
}
