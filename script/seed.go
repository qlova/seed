package script

import (
	"encoding/base64"
	"fmt"
	"math/big"
	"strings"

	"qlova.org/seed"
)

func SetID(id string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		switch c.(type) {
		case Seed, Undo:
			panic("script.SetID must not be called on a script.Seed")
		}

		var data Data
		c.Read(&data)
		data.id = id
		c.Write(data)
	})
}

//ID returns the script ID of this seed.
func ID(c seed.Seed) string {
	c.Use()
	var data Data
	c.Read(&data)

	if data.id != "" {
		return data.id
	}

	id := base64.RawURLEncoding.EncodeToString(big.NewInt(int64(c.ID())).Bytes())

	if id[0] >= '0' && id[0] <= '9' {
		id = "_" + id
	}

	id = strings.Replace(id, "-", "__", -1)

	return id
}

type Undo struct {
	Seed
}

func (c Undo) AddTo(other seed.Seed) {
	c.Q(fmt.Sprintf(`%v.style.display = "none";  if (%[1]v.onhidden) %[1]v.onhidden();`, c.Element()))
}

func (c Undo) With(options ...seed.Option) seed.Seed {
	for _, o := range options {
		if other, ok := o.(seed.Seed); ok {
			o = Undo{Scope(other, c.Q)}
		}
		o.AddTo(c)
	}
	return c
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
	return fmt.Sprintf(`q.get("%v")`, ID(c))
}

func (c Seed) Undo(options ...seed.Option) {
	Undo{c}.With(options...)
}

func (c Seed) Javascript(format string, args ...interface{}) {
	c.Q(fmt.Sprintf(format, args...))
}

func (c Seed) With(options ...seed.Option) seed.Seed {
	for _, o := range options {
		o.AddTo(c)
	}
	return c
}

var p = 0

func (c Seed) AddTo(other seed.Seed) {
	c.Q(fmt.Sprintf(`%v.style.display = ""; if (%[1]v.onvisible) %[1]v.onvisible();`, c.Element()))
	seed.Add(c, other)
}
