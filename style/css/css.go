package css

import (
	"bytes"
)

type Property = *Style

//this is the internal stringable interface used by properties.go
type stringable interface {
	String() string
}

//A stylable can be styled with raw CSS properties being set to raw CSS values.
type Stylable interface {
	Set(property, value string)
	Get(property string) (value string)
	Bytes() []byte
}

// A CSS style for a single element.
type Style struct {
	Stylable
}

func (style Style) CSS() Style {
	return style
}

//This is the internal set method used by properties.go
func (style Style) set(property string, value stringable) {
	style.Stylable.Set(property, value.String())
}

//This returns the CSS style string for the style.
///panics if style implementation is not renderable.
func (style Style) String() string {
	return string(style.Bytes())
}

//The default style implementation.
type Implementation map[string]string

//The raw set implementation.
func (impl Implementation) Set(property, value string) {
	impl[property] = value
}

//The raw get implementation.
func (impl Implementation) Get(property string) string {
	return impl[property]
}

//Returns a new style using the default style implementation.
func NewStyle() Style {
	return Style{Stylable: make(Implementation)}
}

//Returns a CSS style string formatted as bytes.
func (impl Implementation) Bytes() []byte {
	var data bytes.Buffer

	for property, value := range impl {
		data.WriteString(property)
		data.WriteByte(':')
		data.WriteString(value)
		data.WriteByte(';')
	}

	return data.Bytes()
}

//A style implementation that catches properties.
type propertyCatcher string

//The raw set implementation.
func (impl *propertyCatcher) Set(property, value string) {
	*impl = propertyCatcher(property)
}

//The raw get implementation.
func (impl *propertyCatcher) Get(property string) string {
	*impl = propertyCatcher(property)
	return ""
}

//Returns a CSS style string formatted as bytes.
func (impl propertyCatcher) Bytes() []byte {
	return nil
}
