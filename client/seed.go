package client

import (
	"encoding/base64"
	"fmt"
	"math/big"
	"strings"

	"qlova.org/seed"
	"qlova.org/seed/web/js"
)

//Element returns the js.Element of the given seed.
func Element(c seed.Seed) string {
	c.Use()
	return fmt.Sprintf(`q.get("%v")`, ID(c))
}

//Mode on the client, when processing a seed option.
type Mode int8

//SeedTypes
const (
	AddTo Mode = 1
	Undo  Mode = 2
)

//Seed returns the mode and ctx of the given seed.
func Seed(c seed.Seed) (Mode, js.Ctx) {
	var d data
	c.Load(&d)
	return d.mode, d.ctx
}

//Option converts an option into a client option.
func Option(o seed.Option, q js.Ctx) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		c.Save(data{
			mode: AddTo,
			ctx:  q,
		})
		o.AddTo(c)
		c.Save(data{})
	})
}

//Reverse converts an option into a reversed client option.
func Reverse(o seed.Option, q js.Ctx) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		c.Save(data{
			mode: Undo,
			ctx:  q,
		})
		o.AddTo(c)
		c.Save(data{})
	})
}

type data struct {
	mode Mode

	ctx js.Ctx
}

func SetID(id string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		switch mode, _ := Seed(c); mode {
		case AddTo, Undo:
			panic("script.SetID must not be called on a script.Seed")
		}

		var data Data
		c.Load(&data)
		data.id = id
		c.Save(data)
	})
}

//ID returns the client ID of this seed.
func ID(c seed.Seed) string {
	c.Use()
	var data Data
	c.Load(&data)

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

/*type Undo struct {
	Seed

	Option int
}

func (c Undo) AddTo(other seed.Seed) {
	c.Q(fmt.Sprintf(`%v.style.display = "none";  if (%[1]v.onhidden) %[1]v.onhidden();`, c.Element()))
}

func (c Undo) With(options ...seed.Option) seed.Seed {
	for _, o := range options {
		if other, ok := o.(seed.Seed); ok {
			o = Undo{Seed{other, c.Q, other.ID()}, other.ID()}
		}
		o.AddTo(c)
	}
	return c
}*/

//Seed is the script Ctx of a seed.
/*type Seed struct {
	seed.Seed
	Q js.Ctx

	Option int
}

func (c Seed) Undo(options ...seed.Option) {
	Undo{c, c.Option}.With(options...)
}

func (c Seed) Javascript(format string, args ...interface{}) {
	c.Q(fmt.Sprintf(format, args...))
}

var p = 0

func (c Seed) AddTo(other seed.Seed) {
	c.Q(fmt.Sprintf(`%v.style.display = ""; if (%[1]v.onvisible) %[1]v.onvisible();`, c.Element()))
	seed.Add(c, other)
}*/
