package seed

import "github.com/qlova/seed/script/global"

//SetDynamicText sets the text content of the seed which will be dynamic at runtime.
func (seed Seed) SetDynamicText(global global.String) {
	seed.dynamicText = global
}
