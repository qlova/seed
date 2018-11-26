package css

import "bytes"

//A CSS stylesheet interface.
type StylableSheet interface {
	Set(selector string, style Stylable)
	Get(selector string) (style Stylable) //Returns nil if not found.
	Bytes() []byte
}

//A complete CSS Stylesheet with each selector having a style.
type StyleSheet map[string]Stylable

//Add a new CSS style for the specified CSS selector.
func (sheet StyleSheet) Add(selector string, style Stylable) {
	sheet[selector] = style
}

//Retrieve the CSS style for the specified CSS selector.
func (sheet StyleSheet) Get(selector string) Stylable {
	return sheet[selector]
}

//This returns the CSS style string for the style formatted as bytes.
func (sheet StyleSheet) Bytes() []byte {
	var buffer bytes.Buffer

	for selector, style := range sheet {
		buffer.WriteString(selector)
		buffer.WriteByte('{')
		buffer.Write(style.Bytes())
		buffer.WriteByte('}')
	}

	return buffer.Bytes()
}
