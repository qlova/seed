package feed

import (
	"qlova.org/seed/client"
	"qlova.org/seed/use/js"
)

//Into uses the mirror package to create an internal mirror to the given structure.
//The structure can then be used for type-safe field access. Panics if the structure is invalid or unsupported.
func (f *Feed) Into(structure interface{}) {
	f.mirror.Reflect(structure)
}

//String uses the mirror package to identify a field from its string value.
//The field must be a field from a type previously initialised by Feed.Into
func (f *Feed) String(field string) client.String {
	return js.String{Value: js.NewValue("%v"+f.mirror.Path(field), f.Data)}
}
