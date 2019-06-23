package webgl

import qlova "github.com/qlova/script"
import "github.com/qlova/seed/script"

import "github.com/qlova/script/language"
import "github.com/qlova/script/language/javascript"

type BitField string
type ShaderType string
type BindingPoint string
type UsagePattern string
type DataType string
type DrawMode string

type Shader struct {
	qlova.Native
}

type Program struct {
	qlova.Native
}

type Buffer struct {
	qlova.Native
}

type Attribute struct {
	qlova.Native
}

type Context struct {
	GL string
	Q  script.Script

	ColorBufferBit                  BitField
	VertexShader, FragmentShader    ShaderType
	ArrayBuffer, ElementArrayBuffer BindingPoint

	StaticDraw UsagePattern

	Byte, Float DataType

	Triangles DrawMode
}

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

func (ctx *Context) ClearColor(red, green, blue, alpha qlova.Float) {
	ctx.Q.JS().Run(ctx.GL+`.clearColor`, red, green, blue, alpha)
}

func (ctx *Context) Clear(mask BitField) {
	ctx.Q.Javascript(ctx.GL + `.clear(` + ctx.GL + "." + string(mask) + ");")
}

func (ctx *Context) Viewport(x, y, w, h qlova.Float) {
	ctx.Q.JS().Run(ctx.GL+`.viewport`, x, y, w, h)
}

func (ctx *Context) CreateShader(T ShaderType) Shader {
	return Shader{ctx.Q.NativeFromLanguageType(Javascript.Native{
		Expression: ctx.GL + `.createShader(` + ctx.GL + "." + string(T) + ");"}).Var()}
}

func (ctx *Context) ShaderSource(shader Shader, source qlova.String) {
	ctx.Q.JS().Run(ctx.GL+`.shaderSource`, shader, source)
}

func (ctx *Context) CompileShader(shader Shader) {
	ctx.Q.JS().Run(ctx.GL+`.compileShader`, shader)
}

func (ctx *Context) CreateProgram() Program {
	return Program{ctx.Q.NativeFromLanguageType(Javascript.Native{
		Expression: ctx.GL + `.createProgram();`}).Var()}
}

func (ctx *Context) AttachShader(program Program, shader Shader) {
	ctx.Q.JS().Run(ctx.GL+`.attachShader`, program, shader)
}

func (ctx *Context) LinkProgram(program Program) {
	ctx.Q.JS().Run(ctx.GL+`.linkProgram`, program)
}

func (ctx *Context) UseProgram(program Program) {
	ctx.Q.JS().Run(ctx.GL+`.useProgram`, program)
}

func (ctx *Context) CreateBuffer() Buffer {
	return Buffer{ctx.Q.NativeFromLanguageType(Javascript.Native{
		Expression: ctx.GL + `.createBuffer();`}).Var()}
}

func (ctx *Context) BindBuffer(target BindingPoint, buffer Buffer) {
	ctx.Q.Javascript(ctx.GL + `.bindBuffer(` + ctx.GL + "." + string(target) + ", " + buffer.LanguageType().Raw() + ");")
}

func (ctx *Context) BufferData(target BindingPoint, data qlova.List, usage UsagePattern) {

	if _, ok := data.Subtype().LanguageType().(language.Real); !ok {
		panic("Invalid data type")
	}

	ctx.Q.Javascript(ctx.GL + `.bufferData(` + ctx.GL + "." + string(target) + ", new Float32Array(" + data.LanguageType().Raw() + ")," + ctx.GL + "." + string(usage) + ");")
}

func (ctx *Context) GetAttribLocation(program Program, attrib qlova.String) Attribute {
	return Attribute{ctx.Q.JS().Call(ctx.GL+`.getAttribLocation`, program, attrib).Native().Var()}
}

func (ctx *Context) VertexAttribPointer(attribute Attribute, size qlova.Int, datatype DataType, normalized qlova.Bool, stride, offset qlova.Int) {
	ctx.Q.Javascript(ctx.GL + `.vertexAttribPointer(` +
		attribute.LanguageType().Raw() + "," +
		size.LanguageType().Raw() + "," +
		ctx.GL + "." + string(datatype) + "," +
		normalized.LanguageType().Raw() + "," +
		stride.LanguageType().Raw() + "," +
		offset.LanguageType().Raw() + "," +
		");")
}

func (ctx *Context) EnableVertexAttribArray(attribute Attribute) {
	ctx.Q.JS().Run(ctx.GL+`.enableVertexAttribArray`, attribute)
}

func (ctx *Context) DrawArrays(mode DrawMode, first, count qlova.Int) {
	ctx.Q.Javascript(ctx.GL + `.drawArrays(` +
		ctx.GL + "." + string(mode) + "," +
		first.LanguageType().Raw() + "," +
		count.LanguageType().Raw() + "," +
		");")
}
