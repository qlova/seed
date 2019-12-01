package style

import (
	"bytes"
)

//Sheet is a CSS stylesheet that produces optimally compressed CSS for Qlovaseed.
type Sheet struct {
	Rules
	Portrait, Landscape Rules
}

//NewSheet returns a new sheet.
func NewSheet() Sheet {
	return Sheet{
		Rules:     NewRules(),
		Portrait:  NewRules(),
		Landscape: NewRules(),
	}
}

//AddGroup adds a new style group to the sheet.
func (sheet Sheet) AddGroup(selector string, group Group) {
	sheet.Rules.AddStyle(selector, group.Style)
	sheet.Portrait.AddStyle(selector, group.Portrait)
	sheet.Landscape.AddStyle(selector, group.Landscape)
}

//Bytes returns the sheet as CSS.
func (sheet Sheet) Bytes() []byte {
	var buffer bytes.Buffer

	buffer.Write(sheet.Rules.Bytes())

	if len(sheet.Portrait.Standard) > 0 {
		buffer.WriteString(`@media screen and (orientation: portrait) {`)
		buffer.Write(sheet.Portrait.Bytes())
		buffer.WriteString(`}`)
	}

	if len(sheet.Landscape.Standard) > 0 {
		buffer.WriteString(`@media screen and (orientation: landscape) {`)
		buffer.Write(sheet.Landscape.Bytes())
		buffer.WriteString(`}`)
	}

	return buffer.Bytes()
}
