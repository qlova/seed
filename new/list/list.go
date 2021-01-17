package list

import (
	"qlova.org/seed"
	"qlova.org/seed/new/column"
)

type data struct {
	Items []seed.Seed

	Options []Option
}

//New returns a new list with the given options.
func New(options ...seed.Option) seed.Seed {
	var col = column.New(options...)

	var data data
	col.Load(&data)

	for _, op := range data.Options {
		op(col, &data.Items)
	}

	for _, item := range data.Items {
		item.AddTo(col)
	}

	return col
}

//Set sets the list's items to the given seed slice.
func Set(items ...seed.Seed) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var data data
		c.Load(&data)

		data.Items = items

		c.Save(data)
	})
}

//Option allows the list to be modified by an options.
type Option func(c seed.Seed, items *[]seed.Seed)

//AddTo implements seed.Option
func (o Option) AddTo(c seed.Seed) {
	var data data
	c.Load(&data)

	data.Options = append(data.Options, o)

	c.Save(data)
}
