package repeater

import (
	"reflect"
	"sort"

	"qlova.org/seed"
)

type Data struct {
	index, value interface{}
}

func (d Data) String() string {
	return d.value.(string)
}

func (d Data) Int() int {
	return d.value.(int)
}

func (d Data) Index() Data {
	return Data{nil, d.index}
}

func (d Data) Interface() interface{} {
	return d.value
}

type Seed struct {
	seed.Seed
	Data Data
}

type seedData struct {
	data Data
}

//New returns a repeater capable of repeating itself based on the given Go data.
func New(data interface{}, options ...seed.Option) seed.Seed {
	var repeater = seed.New()
	var value = reflect.ValueOf(data)
	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			var d seedData
			repeater.Load(&d)

			d.data = Data{i, value.Index(i).Interface()}
			repeater.Save(d)

			for _, o := range options {
				repeater = repeater.With(o)
			}
		}
	case reflect.Map:

		//Deterministic render.
		keys := make([]string, 0, value.Len())

		for _, key := range value.MapKeys() {
			var keystring, ok = key.Interface().(string)
			if !ok {
				panic("nondeterministic data type passed to repeater")
			}
			keys = append(keys, keystring)
		}
		sort.Strings(keys)

		for _, key := range keys {
			var d seedData
			repeater.Load(&d)

			d.data = Data{key, value.MapIndex(reflect.ValueOf(key)).Interface()}
			repeater.Save(d)

			for _, o := range options {
				repeater.With(o)
			}
		}

	default:
		panic("repeater.New: unsupported data type: " + reflect.TypeOf(data).String())
	}
	return repeater
}

//Do runs f.
func Do(f func(Seed)) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var d seedData
		c.Load(&d)
		f(Seed{c, d.data})
	})
}
