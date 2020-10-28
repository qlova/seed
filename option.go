package seed

import (
	"reflect"

	"github.com/qlova/seed"
)

//Option can be used to modify a seed.
type Option interface {
	AddTo(Seed)
}

//Options is a multiple of Option.
type Options []Option

//AddTo implements Option, Options are a type of Option.
func (options Options) AddTo(c Seed) {
	for _, o := range options {
		if o != nil {
			o.AddTo(c)
		}
	}
}

//OptionFunc can be used to create an Option.
type OptionFunc func(c Seed)

//NewOption can be used to create options.
type NewOption = OptionFunc

//AddTo implements Option.AddTo
func (o OptionFunc) AddTo(c Seed) {
	o(c)
}

//And implements Option.And
func (o OptionFunc) And(more ...Option) Option {
	return And(o, more...)
}

//And implements Option.And
func And(o Option, more ...Option) Option {
	return NewOption(func(c Seed) {
		o.AddTo(c)
		for _, o = range more {
			o.AddTo(c)
		}
	})
}

//If applies the options if the condition is true.
func If(condition bool, options ...Option) Option {
	return NewOption(func(c Seed) {
		if condition {
			for _, o := range options {
				o.AddTo(c)
			}
		}
	})
}

//Mutate a seed with the given data.
//panics on illegal arguments.
func Mutate(f interface{}) Option {
	T := reflect.TypeOf(f)
	if T.Kind() != reflect.Func || T.In(0).Kind() != reflect.Ptr || T.NumIn() > 2 ||
		(T.NumIn() == 2 && T.In(1) != reflect.TypeOf(seed.Seed{})) {
		panic("illegal argument to seed.Mutate")
	}

	V := reflect.ValueOf(f)

	data := reflect.New(T.In(0).Elem())

	return NewOption(func(c Seed) {
		c.Load(data.Interface())

		if T.NumIn() == 2 {
			V.Call([]reflect.Value{data, reflect.ValueOf(c)})
		} else {
			V.Call([]reflect.Value{data})
		}

		c.Save(data.Elem().Interface())
	})
}
