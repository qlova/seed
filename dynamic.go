package seed

import "github.com/qlova/seed/script"

//Set the text content of the seed which will be dynamic at runtime.
func (seed Seed) SetDynamicText(global script.GlobalString) {
	seed.dynamicText = global
}
