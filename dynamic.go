package seed

import "github.com/qlova/seed/script/global"

type dynamic struct {
	Text, Source global.String
}

//SetDynamicText sets the text content of the seed which will be dynamic at runtime.
func (seed Seed) SetDynamicText(s global.String) {
	seed.dynamic.Text = s
}

//SetDynamicSource sets the source of the seed which will be dynamic at runtime.
func (seed Seed) SetDynamicSource(s global.String) {
	seed.dynamic.Source = s
}
