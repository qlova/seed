package repeater

import (
	"reflect"

	"github.com/qlova/seed"
)

type Interface struct {
	any interface{}
}

func (i Interface) String() string {
	s, _ := i.any.(string)
	return s
}

type data struct {
	Value Interface
}

var seeds = make(map[seed.Seed]data)

//New returns a repeater capable of repeating itself based on the given Go data.
func New(data interface{}, options ...seed.Option) seed.Seed {
	var repeater = seed.New()
	var value = reflect.ValueOf(data)
	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			data := seeds[repeater]
			data.Value = Interface{value.Index(i).Interface()}
			seeds[repeater] = data
			for _, o := range options {
				repeater.Add(o)
			}
		}
		return repeater
	default:
		panic("repeater.New: unsupported data type")
	}
}

//Data returns the current repeater data associated with the current seed.
func Data(c seed.Seed) Interface {
	data, _ := seeds[c]
	return data.Value
}
