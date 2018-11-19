package style

import (
	"bytes"
)

type Css interface {
	Set(string, string)
}

type StaticCss struct {
	styles map[string]string
}

func (css *StaticCss) Set(property, value string) {
	css.styles[property] = value
}


func New() Style {
	return Style{Css: &StaticCss{styles: make(map[string]string)}}
}

func (style Style) Render() []byte {
	var data bytes.Buffer
	
	for property, value := range style.Css.(*StaticCss).styles {
		data.WriteString(property)
		data.WriteByte(':')
		data.WriteString(value)
		data.WriteByte(';')
	}
	
	return data.Bytes()
}
