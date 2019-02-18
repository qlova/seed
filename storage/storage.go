package storage

import "reflect"
import "bytes"
import "encoding/json"

type View struct {
	Node
	Path []string
}

func (v View) JSON() JSON {
	return JSON{v}
}

func (v View) Put(data []byte) []byte {
	return v.Node.Put(v, data)
}

func (v View) Create() bool {
	return v.Node.Create(v)
}

func (v View) ForEach(f func(key, data []byte)) {
	v.Node.ForEach(v, f)
}

type JSON struct {
	View
}

func (view JSON) Put(pointer interface{}) []byte {
	var buffer bytes.Buffer

	err := json.NewEncoder(&buffer).Encode(pointer)
	if err != nil {
		//TODO report error?
		return nil
	}

	return view.View.Put(buffer.Bytes())
}

func (view JSON) ForEach(structure interface{}, f func(id []byte, value interface{})) {

	var shell = reflect.New(reflect.TypeOf(structure)).Interface()

	view.View.ForEach(func(key []byte, data []byte) {

		

		json.NewDecoder(bytes.NewReader(data)).Decode(shell)
		
		f(key, reflect.ValueOf(shell).Elem().Interface())
		
	})

}