package repeater

import (
	"reflect"

	"github.com/qlova/seed"
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
	seed.Data

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
			repeater.Read(&d)

			d.data = Data{i, value.Index(i).Interface()}
			repeater.Write(d)

			for _, o := range options {
				repeater.Add(o)
			}
		}
	case reflect.Map:
		for _, i := range value.MapKeys() {
			var d seedData
			repeater.Read(&d)

			d.data = Data{i.Interface(), value.MapIndex(i).Interface()}
			repeater.Write(d)

			for _, o := range options {
				repeater.Add(o)
			}
		}

	default:
		panic("repeater.New: unsupported data type: " + reflect.TypeOf(data).String())
	}
	return repeater
}

//Do runs f.
func Do(f func(Seed)) seed.Option {
	return seed.Do(func(c seed.Seed) {
		var d seedData
		c.Read(&d)
		f(Seed{c, d.data})
	})
}
