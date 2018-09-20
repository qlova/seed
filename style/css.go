package style

import (
	"bytes"
)

type Css interface {
	Set(string, string)
}

type StaticCss struct {
	Data bytes.Buffer
}

func (css *StaticCss) Set(property, value string) {
	css.Data.WriteString(property)
	css.Data.WriteByte(':')
	css.Data.WriteString(value)
	css.Data.WriteByte(';')
}


func New() Style {
	return Style{Css: &StaticCss{}}
}
