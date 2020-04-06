package script

import (
	"encoding/base64"
	"fmt"
	"math/big"

	"github.com/qlova/seed"
)

func SetID(id string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		switch c.(type) {
		case Seed, Undo:
			panic("script.SetID must not be called on a script.Seed")
		}

		var data data
		c.Read(&data)
		data.id = id
		c.Write(data)
	})
}

//ID returns the script ID of this seed.
func ID(c seed.Seed) string {
	var data data
	c.Read(&data)

	if data.id != "" {
		return data.id
	}

	return base64.RawURLEncoding.EncodeToString(big.NewInt(int64(c.ID())).Bytes())
}

type Undo struct {
	Seed
}

func (c Undo) AddTo(other seed.Seed) {
	c.Q(fmt.Sprintf(`%v.style.display = "none";`, c.Element()))
}

func (c Undo) Add(options ...seed.Option) seed.Option {
	for _, o := range options {
		if other, ok := o.(seed.Seed); ok {
			o = Undo{Scope(other, c.Q)}
		}
		o.AddTo(c)
	}
	return seed.NewOption(func(seed.Seed) {})
}

//Seed is the script Ctx of a seed.
type Seed struct {
	seed.Seed
	Q Ctx
}

func Scope(c seed.Seed, q Ctx) Seed {
	return Seed{c, q}
}

func (c Seed) Element() string {
	c.Use()
	return fmt.Sprintf(`seed.get("%v")`, ID(c))
}

func (c Seed) Undo(options ...seed.Option) {
	Undo{c}.Add(options...)
}

func (c Seed) Javascript(format string, args ...interface{}) {
	c.Q(fmt.Sprintf(format, args...))
}

func (c Seed) Add(options ...seed.Option) seed.Option {
	for _, o := range options {
		o.AddTo(c)
	}
	return seed.NewOption(func(seed.Seed) {})
}

var p = 0

func (c Seed) AddTo(other seed.Seed) {
	c.Q(fmt.Sprintf(`%v.style.display = "";`, c.Element()))
	seed.Add(c, other)
}
