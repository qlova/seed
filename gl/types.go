//+build ignore

package gl

//TextureUnit is a GPU texture register.
type TextureUnit int32

//BufferTarget is buffer binding point / target.
type BufferTarget uint8

//UsagePattern specifies the intended usage pattern of the data store for optimization purposes.
type UsagePattern uint8

//DataType specified the format of buffer data.
type DataType uint8

//DrawMode specifying the primitive to render.
type DrawMode uint8

//Attribute is a shader attribute.
type Attribute struct {
	ctx *Context

	int32
}

//Load binds the buffer currently bound to gl.ArrayBuffer to a generic vertex attribute of the current vertex buffer object and specifies its layout.
func (attr *Attribute) Load(size byte, T DataType, normalized bool, stride, offset int32) {
	attr.ctx.send(vertexAttribPointer, attr.int32, size, T, normalized, byte(0), stride, offset)
}

//Enable turns on a vertex attribute at a specific index position in
// a vertex attribute array.
func (attr *Attribute) Enable() {
	attr.ctx.send(enableVertexAttribArray, attr.int32)
}

//Program is a GPU program.
type Program struct {
	ctx *Context

	int32
	vertex, fragment Shader

	attributes map[string]Attribute
}

// Attribute returns a named attribute variable.
func (program *Program) Attribute(name string) Attribute {
	program.ctx.attributes++
	program.ctx.send(getAttribLocation, program.int32, name)
	return Attribute{program.ctx, program.ctx.attributes}
}

//Attach attaches a Shader to a Program.
func (program *Program) Attach(shader Shader) {
	if shader.T == program.ctx.VertexShader {
		program.vertex = shader
	}
	if shader.T == program.ctx.FragmentShader {
		program.fragment = shader
	}
	shader.ctx.send(attachShader, program.int32, shader.int32)
}

// Link links an attached vertex shader and an attached fragment shader
// to a program so it can be used by the graphics processing unit (GPU).
func (program *Program) Link() {
	program.ctx.send(linkProgram, program.int32)
}

// Use sets the program object to use for rendering.
func (program *Program) Use() {
	program.ctx.send(useProgram, program.int32)
}

//Shader is a GPU shader.
type Shader struct {
	ctx *Context

	T ShaderType
	int32
	string
}

//SetSource sets and replaces shader source code in a shader object.
func (shader *Shader) SetSource(source string) {
	shader.string = source
	shader.ctx.send(shaderSource, shader.int32, source)
}

//Compile compiles the GLSL shader source into binary data used by the GL Program.
func (shader *Shader) Compile() {
	shader.ctx.send(compileShader, shader.int32)
}

//Buffer is a GPU buffer.
type Buffer struct {
	ctx *Context

	int32
	target BufferTarget
}

//Bind associates a buffer with a buffer target.
func (buffer *Buffer) Bind(target BufferTarget) {
	buffer.ctx.send(bindBuffer, target, byte(0), byte(0), byte(0), buffer.int32)
}

//ShaderType is a type of shader.
type ShaderType uint8

//BitField specifies a buffer bit field type.
type BitField uint8
