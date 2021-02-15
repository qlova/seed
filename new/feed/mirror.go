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
//The field must be a field previously initialised by Feed.Into
func (f *Feed) String(field string) client.String {
	return js.String{Value: js.NewValue("%v"+f.mirror.Path(field), f.Data)}
}

//Int uses the mirror package to identify a field from its int value.
//The field must be a field previously initialised by Feed.Into
func (f *Feed) Int(field int) client.Int {
	return js.Number{Value: js.NewValue("%v"+f.mirror.Path(field), f.Data)}
}

//Int8 uses the mirror package to identify a field from its int8 value.
//The field must be a field previously initialised by Feed.Into
func (f *Feed) Int8(field int8) client.Int {
	return js.Number{Value: js.NewValue("%v"+f.mirror.Path(field), f.Data)}
}

//Int16 uses the mirror package to identify a field from its int16 value.
//The field must be a field previously initialised by Feed.Into
func (f *Feed) Int16(field int16) client.Int {
	return js.Number{Value: js.NewValue("%v"+f.mirror.Path(field), f.Data)}
}

//Int32 uses the mirror package to identify a field from its int32 value.
//The field must be a field previously initialised by Feed.Into
func (f *Feed) Int32(field int32) client.Int {
	return js.Number{Value: js.NewValue("%v"+f.mirror.Path(field), f.Data)}
}

//Float32 uses the mirror package to identify a field from its float32 value.
//The field must be a field previously initialised by Feed.Into
func (f *Feed) Float32(field float32) client.Float {
	return js.Number{Value: js.NewValue("%v"+f.mirror.Path(field), f.Data)}
}

//Float64 uses the mirror package to identify a field from its float64 value.
//The field must be a field previously initialised by Feed.Into
func (f *Feed) Float64(field float32) client.Float {
	return js.Number{Value: js.NewValue("%v"+f.mirror.Path(field), f.Data)}
}

//Bool uses the mirror package to identify a field from its bool value.
//The field must be a field previously initialised by Feed.Into
func (f *Feed) Bool(field bool) client.Bool {
	return js.Bool{Value: js.NewValue("%v"+f.mirror.Path(field), f.Data)}
}
