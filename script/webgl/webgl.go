package webgl

import (
	qlova "github.com/qlova/script"
	"github.com/qlova/script/language"
	"github.com/qlova/seed/script"
)

//BitField specifies a buffer bit field type.
type BitField string

//ShaderType is a type of shader.
type ShaderType string

//BindingPoint is a buffer binding point / target.
type BindingPoint string

//UsagePattern specifies the intended usage pattern of the data store for optimization purposes.
type UsagePattern string

//DataType specified the format of buffer data.
type DataType string

//DrawMode specifying the primitive to render.
type DrawMode string

//Shader is a GPU shader.
type Shader struct {
	qlova.Native
}

//Program is a GPU program.
type Program struct {
	qlova.Native
}

//Buffer is a GPU buffer.
type Buffer struct {
	qlova.Native
}

//Attribute is a shader attribute.
type Attribute struct {
	qlova.Native
}

//Context is a gl context for rendering to the screen.
type Context struct {
	GL string
	Q  script.Ctx

	ColorBufferBit                  BitField
	VertexShader, FragmentShader    ShaderType
	ArrayBuffer, ElementArrayBuffer BindingPoint

	StaticDraw UsagePattern

	Byte, Float DataType

	Triangles DrawMode
}

//NewContext returns a gl context from the specified seed.
func NewContext(canvas script.Seed) Context {
	var unique = script.Unique()
	canvas.Q.Javascript(`let ` + unique + ` = ` + canvas.Element() + `.getContext("webgl");`)
	return Context{
		GL: unique,
		Q:  canvas.Q,

		ColorBufferBit: "COLOR_BUFFER_BIT",

		VertexShader:   "VERTEX_SHADER",
		FragmentShader: "FRAGMENT_SHADER",

		ArrayBuffer:        "ARRAY_BUFFER",
		ElementArrayBuffer: "ELEMENT_ARRAY_BUFFER",

		StaticDraw: "STATIC_DRAW",

		Byte:  "BYTE",
		Float: "FLOAT",

		Triangles: "TRIANGLES",
	}
}

//ClearColor clears the color of the buffer.
func (ctx *Context) ClearColor(red, green, blue, alpha qlova.Float) {
	ctx.Q.JS().Run(ctx.GL+`.clearColor`, red, green, blue, alpha)
}

//Clear clears the mask.
func (ctx *Context) Clear(mask BitField) {
	ctx.Q.Javascript(ctx.GL + `.clear(` + ctx.GL + "." + string(mask) + ");")
}

//Viewport sets the viewport, which specifies the affine transformation of x and y from normalized device coordinates to window coordinates.
func (ctx *Context) Viewport(x, y, w, h qlova.Float) {
	ctx.Q.JS().Run(ctx.GL+`.viewport`, x, y, w, h)
}

//CreateShader returns an empty vertex or fragment shader object based on the type specified.
func (ctx *Context) CreateShader(T ShaderType) Shader {
	return Shader{script.Native{language.Expression(ctx.Q, ctx.GL+`.createShader(`+ctx.GL+"."+string(T)+");")}}
}

//ShaderSource sets and replaces shader source code in a shader object.
func (ctx *Context) ShaderSource(shader Shader, source qlova.String) {
	ctx.Q.JS().Run(ctx.GL+`.shaderSource`, shader, source)
}

//CompileShader compiles the GLSL shader source into binary data used by the GL Program.
func (ctx *Context) CompileShader(shader Shader) {
	ctx.Q.JS().Run(ctx.GL+`.compileShader`, shader)
}

//CreateProgram creates an empty Program  to which shaders can be bound.
func (ctx *Context) CreateProgram() Program {
	return Program{script.Native{language.Expression(ctx.Q, ctx.GL+`.createProgram();`)}}
}

//AttachShader attaches a Shader to a Program.
func (ctx *Context) AttachShader(program Program, shader Shader) {
	ctx.Q.JS().Run(ctx.GL+`.attachShader`, program, shader)
}

//LinkProgram links an attached vertex shader and an attached fragment shader
// to a program so it can be used by the graphics processing unit (GPU).
func (ctx *Context) LinkProgram(program Program) {
	ctx.Q.JS().Run(ctx.GL+`.linkProgram`, program)
}

//UseProgram sets the program object to use for rendering.
func (ctx *Context) UseProgram(program Program) {
	ctx.Q.JS().Run(ctx.GL+`.useProgram`, program)
}

//CreateBuffer creates and initializes a Buffer.
func (ctx *Context) CreateBuffer() Buffer {
	return Buffer{script.Native{language.Expression(ctx.Q, ctx.GL+`.createBuffer();`)}}
}

//BindBuffer associates a buffer with a buffer target.
func (ctx *Context) BindBuffer(target BindingPoint, buffer Buffer) {
	ctx.Q.Javascript(ctx.GL + `.bindBuffer(` + ctx.GL + "." + string(target) + ", " + buffer.Raw() + ");")
}

//BufferData creates a buffer in memory and initializes it with array data.
//If no array is provided, the contents of the buffer is initialized to 0.
// panics if data is not a Float List.
func (ctx *Context) BufferData(target BindingPoint, data qlova.List, usage UsagePattern) {

	ctx.Q.Javascript(ctx.GL + `.bufferData(` + ctx.GL + "." + string(target) + ", new Float32Array(" + data.Raw() + ")," + ctx.GL + "." + string(usage) + ");")
}

//GetAttribLocation returns a named attribute variable.
func (ctx *Context) GetAttribLocation(program Program, attrib qlova.String) Attribute {
	return Attribute{ctx.Q.JS().Call(ctx.GL+`.getAttribLocation`, program, attrib).Native()}
}

//VertexAttribPointer binds the buffer currently bound to gl.ArrayBuffer to a generic vertex attribute of the current vertex buffer object and specifies its layout.
func (ctx *Context) VertexAttribPointer(attribute Attribute, size qlova.Int, datatype DataType, normalized qlova.Bool, stride, offset qlova.Int) {
	ctx.Q.Javascript(ctx.GL + `.vertexAttribPointer(` +
		attribute.Raw() + "," +
		size.Raw() + "," +
		ctx.GL + "." + string(datatype) + "," +
		normalized.Raw() + "," +
		stride.Raw() + "," +
		offset.Raw() + "," +
		");")
}

//EnableVertexAttribArray turns on a vertex attribute at a specific index position in
// a vertex attribute array.
func (ctx *Context) EnableVertexAttribArray(attribute Attribute) {
	ctx.Q.JS().Run(ctx.GL+`.enableVertexAttribArray`, attribute)
}

//DrawArrays renders geometric primitives from bound and enabled vertex data.
func (ctx *Context) DrawArrays(mode DrawMode, first, count qlova.Int) {
	ctx.Q.Javascript(ctx.GL + `.drawArrays(` +
		ctx.GL + "." + string(mode) + "," +
		first.Raw() + "," +
		count.Raw() + "," +
		");")
}
