package clientside

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/use/js"
)

//Bool is an bool variable in client memory.
type Bool struct {
	Name string

	address Address
	Memory  Memory

	Value client.Bool

	not bool
}

//Variable implements Variable
func (b *Bool) Variable() (Address, Memory) {
	if b.address == "" {
		if b.Name != "" {
			b.address = Address(b.Name)
		} else {
			b.address = NewAddress()
		}
	}
	return b.address, b.Memory
}

//GetBool implements client.Bool
func (b *Bool) GetBool() js.Bool {
	address, memory := b.Variable()

	var bool = js.Bool{Value: js.Call(js.Func("q.getvar"),
		client.NewString(string(address)), client.NewString(string(memory)))}

	if b.not {
		return bool.Not()
	}
	return bool
}

//GetValue implements client.Value
func (b *Bool) GetValue() js.Value {
	return b.GetBool().Value
}

//GetDefaultValue implements Variable
func (b *Bool) GetDefaultValue() client.Value {
	if b.Value == nil {
		return client.NewBool(false)
	}
	return b.Value
}

//Set returns a script that sets the bool to the given literal.
func (b *Bool) Set(literal bool) client.Script {
	address, memory := b.Variable()

	var bool = client.NewBool(literal).GetBool()

	if b.not {
		bool = bool.Not()
	}

	return js.Run(js.Func("q.setvar"), client.NewString(string(address)),
		client.NewString(string(memory)), bool)
}

//SetTo returns a script that sets the bool to the given client.Bool.
func (b *Bool) SetTo(value client.Bool) client.Script {
	address, memory := b.Variable()
	return js.Run(js.Func("q.setvar"), client.NewString(string(address)),
		client.NewString(string(memory)), value)
}

//Toggle returns a script that toggles the boolean between true and false.
func (b *Bool) Toggle() client.Script {
	return b.SetTo(b.Not())
}

//OnChange runs the given script when the value of this string is changed.
func (b *Bool) OnChange(do ...client.Script) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var data data
		c.Load(&data)
		data.hooks = append(data.hooks, hook{
			variable: b,
			do:       client.NewScript(do...),
		})
		c.Save(data)
	})

}

//Not returns a clientside bool that is the inverse of b.
//Setting the returned bool has the inverse effect.
func (b *Bool) Not() *Bool {
	b.Variable()

	var not = *b
	not.not = true
	return &not
}

//Protect ensures that the given script will only have one running instance.
func (b *Bool) Protect(do ...client.Script) js.Script {
	return js.If(b.Not(),
		js.Try(
			client.NewScript(b.Set(true), client.NewScript(do...), b.Set(false)).GetScript(),
		).Catch(
			client.NewScript(b.Set(false), js.Throw(js.NewValue("e"))).GetScript(),
			"e"),
	)
}
