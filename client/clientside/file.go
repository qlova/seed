package clientside

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/js"
)

//File is a (potentially zipped) file in client memory.
type File struct {
	Name string

	address Address
	Memory  Memory
}

//Variable implements Variable
func (f *File) Variable() (Address, Memory) {
	if f.address == "" {
		if f.Name != "" {
			f.address = Address(f.Name)
		} else {
			f.address = NewAddress()
		}
	}
	return f.address, f.Memory
}

//GetFile implements client.File
func (f *File) GetFile() js.Value {
	address, memory := f.Variable()

	return js.Call(js.Func("q.getvar"),
		client.NewString(string(address)), client.NewString(string(memory)))
}

//GetValue implements client.Value
func (f *File) GetValue() js.Value {
	return f.GetFile().GetValue()
}

//GetBool implements client.Value
func (f *File) GetBool() js.Bool {
	return f.GetFile().GetValue().GetBool()
}

//GetDefaultValue implements Variable
func (f *File) GetDefaultValue() client.Value {
	return js.Null()
}

//SetToRaw returns a script that sets the file to the given client value.
func (f *File) SetToRaw(value client.Value) client.Script {
	address, memory := f.Variable()
	return js.Run(js.Func("q.setvar"), client.NewString(string(address)),
		client.NewString(string(memory)), value)
}

//OnChange runs the given script when the value of this string is changed.
func (f *File) OnChange(do ...client.Script) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var data data
		c.Read(&data)
		data.hooks = append(data.hooks, hook{
			variable: f,
			do:       client.NewScript(do...),
		})
		c.Write(data)
	})

}
