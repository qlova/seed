package webgl

import qlova "github.com/qlova/script"
import "github.com/qlova/seed/script"

type BitField string

type Context struct {
	GL string
	Q script.Script
	
	ColorBufferBit BitField
}

func NewContext(canvas script.Seed) Context {
	var unique = script.Unique()
	canvas.Q.Javascript(`let `+unique+` = `+canvas.Element()+`.getContext("webgl");`)
	return Context{
		GL: unique,
		Q: canvas.Q,
		
		ColorBufferBit: "COLOR_BUFFER_BIT",
	}
}

func (ctx *Context) ClearColor(red, green, blue, alpha qlova.Float) {
	ctx.Q.JS().Run(ctx.GL+`.clearColor`, red, green, blue, alpha)
}

func (ctx *Context) Clear(mask BitField) {
	ctx.Q.Javascript(ctx.GL+`.clear(`+ctx.GL+"."+string(mask)+")")
}
