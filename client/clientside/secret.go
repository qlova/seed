package clientside

import (
	"github.com/google/uuid"
	"qlova.org/seed/client"
	"qlova.org/seed/js"
	"qlova.org/seed/script"
)

//Secret is a 'secret' string variable in client memory.
//Only the hash is available to read from Go.
//Use this for passwords.
type Secret struct {
	Name string

	address Address
	Memory  Memory

	//Pepper should be set to a unique random string.
	Pepper string

	//TODO support
	//Salt *clientside.String

	//Tweak these values to adjust the hashing difficulty. CPU time and RAM is in KiB.
	//Default is CPU: 1, RAM: 1024
	CPU, RAM int

	//Length is the length of the readable hash of this secret. 32 by default.
	Length int
}

var _ Variable = new(Secret)

//GetDefaultValue implements Variable
func (s *Secret) GetDefaultValue() client.Value {
	return client.NewString("")
}

//Variable implements Variable
func (s *Secret) Variable() (Address, Memory) {
	if s.address == "" {
		if s.Name != "" {
			s.address = Address(s.Name)
		} else {
			id, _ := uuid.NewRandom()
			s.address = Address(id.String())
		}
	}
	if (s.Pepper) == "" {
		panic("pepper is required for clientside.Secret")
	}
	for len(s.Pepper) < 8 {
		s.Pepper = s.Pepper + s.Pepper
	}
	if s.CPU == 0 {
		s.CPU = 1
	}
	if s.RAM == 0 {
		s.RAM = 1024
	}
	if s.Length == 0 {
		s.Length = 32
	}
	return s.address, s.Memory
}

//GetString implements script.AnyString
func (s *Secret) GetString() script.String {
	address, memory := s.Variable()
	return js.String{Value: js.NewValue(`"'#(import "/argon2_exec.js")(await argon2.hash({
		pass: q.getvar(%v, %v),
		salt: %v,

		time: %v,
		mem: %v,
		hashLen: %v,
		parallelism: 1,
		type: argon2.ArgonType.Argon2di
	})).hashHex`, client.NewString(string(address)), client.NewString(string(memory)),
		client.NewString(s.Pepper),
		client.NewInt(int(s.CPU)),
		client.NewInt(int(s.RAM)),
		client.NewInt(int(s.Length)),
	)}
}

//GetBool implements script.AnyBool
func (s *Secret) GetBool() script.Bool {
	address, memory := s.Variable()
	return js.Bool{Value: js.NewValue(`(q.getvar(%v, %v) != "")`, client.NewString(string(address)), client.NewString(string(memory)))}
}

//GetValue implements script.AnyValue
func (s *Secret) GetValue() script.Value {
	return s.GetString().Value
}

//SetTo returns a script that sets the secret to the given client string.
func (s *Secret) SetTo(value client.String) client.Script {
	address, memory := s.Variable()
	return js.Run(js.Func("q.setvar"), client.NewString(string(address)),
		client.NewString(string(memory)), value)
}

func (s *Secret) Equals(b client.String) client.Bool {
	address, memory := s.Variable()
	if sec2, ok := b.(*Secret); ok {
		address2, memory2 := sec2.Variable()
		return js.NewValue(`(q.getvar(%v, %v) == q.getvar(%v, %v))`,
			client.NewString(string(address)),
			client.NewString(string(memory)),
			client.NewString(string(address2)),
			client.NewString(string(memory2)),
		)
	}

	return js.NewValue(`(q.getvar(%v, %v) == %v)`,
		client.NewString(string(address)),
		client.NewString(string(memory)),
		b,
	)
}
