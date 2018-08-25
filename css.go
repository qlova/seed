package app

import (
	"bytes"
	"strings"
)

type Css interface {
	Set(string, string)
}

type StaticCss struct {
	data bytes.Buffer
}

func (css *StaticCss) Set(property, value string) {
	css.data.WriteString(property)
	css.data.WriteByte(':')
	css.data.WriteString(value)
	css.data.WriteByte(';')
}

type ScriptCss struct {
	script *Script
	app *ScriptApp
}

func (css *ScriptCss) Set(property, value string) {
	
	if strings.Contains(property, "-") {
		splits := strings.Split(property, "-")
		
		property = splits[0] + strings.Title(splits[1])
	}
	
	css.script.data.WriteString(`document.getElementById("`)
	css.script.data.WriteString(css.app.ID())
	css.script.data.WriteString(`").style.`)
	css.script.data.WriteString(property)
	css.script.data.WriteByte('=')
	css.script.data.WriteByte('"')
	css.script.data.WriteString(value)
	css.script.data.WriteByte('"')
	css.script.data.WriteByte(';')
}
