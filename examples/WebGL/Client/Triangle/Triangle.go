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
		var canvas = Canvas.Ctx(q)
		var gl = canvas.OpenGL()

		//gl.Viewport(q.Float(0), q.Float(0), canvas.Width(), canvas.Height())

		var VertexShader = gl.CreateShader(gl.VertexShader)
		gl.ShaderSource(VertexShader, q.String(`attribute vec3 c;void main(void){gl_Position=vec4(c, 1.0);}`))
		gl.CompileShader(VertexShader)

		var FragmentShader = gl.CreateShader(gl.FragmentShader)
		gl.ShaderSource(FragmentShader, q.String(`void main(void){gl_FragColor=vec4(0,1,1,1);}`))
		gl.CompileShader(FragmentShader)

		var Program = gl.CreateProgram()
		gl.AttachShader(Program, VertexShader)
		gl.AttachShader(Program, FragmentShader)
		gl.LinkProgram(Program)
		gl.UseProgram(Program)

		gl.ClearColor(q.Float(1), q.Float(0), q.Float(1), q.Float(1))
		gl.Clear(gl.ColorBufferBit)

		var VertexBuffer = gl.CreateBuffer()
		gl.BindBuffer(gl.ArrayBuffer, VertexBuffer)
		gl.BufferData(gl.ArrayBuffer, q.List(
			q.Float(-0.5), q.Float(0.5), q.Float(0),
			q.Float(-0.5), q.Float(-0.5), q.Float(0),
			q.Float(0.5), q.Float(-0.5), q.Float(0),
		), gl.StaticDraw)

		var Coordinate = gl.GetAttribLocation(Program, q.String("c"))
		gl.VertexAttribPointer(Coordinate, q.Int(3), gl.Float, q.Bool(false), q.Int(0), q.Int(0))
		gl.EnableVertexAttribArray(Coordinate)

		gl.DrawArrays(gl.Triangles, q.Int(0), q.Int(3))
	})

	App.Launch()
}
