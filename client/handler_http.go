package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"qlova.org/seed/client/clientsafe"
	"qlova.org/seed/use/js"
)

//Requesty is any type that can load itself from a Request.
type Requesty interface {
	FromRequest(Request) error
}

//Handler returns a handler for handling remote procedure calls.
func Handler(w http.ResponseWriter, r *http.Request, id string) {
	f, ok := goExports[id]
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte("404 not found"))
		return
	}

	var in = make([]reflect.Value, 0, f.Type().NumIn())
	var cr = NewRequest(w, r)

	var StartFrom = 0
	//The function can take an optional client as it's first argument.
	if f.Type().NumIn() > 0 && f.Type().In(0) == reflect.TypeOf(Request{}) {
		StartFrom = 1

		//Pass client.Request as the first argument.
		in = append(in, reflect.ValueOf(cr))

	}

	//The function can take an optional Requesty as it's first argument.
	if f.Type().NumIn() > 0 && reflect.PtrTo(f.Type().In(0)).Implements(reflect.TypeOf([0]Requesty{}).Elem()) {
		StartFrom = 1

		var zero = reflect.New(f.Type().In(0))

		if err := zero.Interface().(Requesty).FromRequest(cr); err != nil {
			log.Println(err)

			safe, ok := err.(clientsafe.Error)
			if ok {
				fmt.Fprintf(w, "throw %v;", strconv.Quote(safe.ClientError()))
			}
			return
		}

		//Make the user, the first argument.
		in = append(in, zero.Elem())

	}

	//Parse each argument as JSON.
	for i := StartFrom; i < f.Type().NumIn(); i++ {

		var key = string(rune('a' + (i - StartFrom)))
		var val = cr.request.FormValue(key)

		var rvalue reflect.Value

		switch f.Type().In(i) {

		//Argument is a file/stream.
		case reflect.TypeOf(Stream{}), reflect.TypeOf([0]io.Reader{}).Elem():
			var stream Stream

			file, header, err := cr.request.FormFile(key)
			if err == nil {
				stream.head = header
				stream.file = file
			}

			rvalue = reflect.ValueOf(stream)

		case reflect.TypeOf(time.Time{}):
			t, err := time.Parse(`"2006-01-02"`, val)
			if err != nil {
				t, err = time.Parse(`"2006-01"`, val)
				if err != nil {
					w.WriteHeader(400)
					log.Println(err)
					fmt.Fprint(w, "throw 'bad request';")
					return
				}
			}
			rvalue = reflect.ValueOf(t)

		case reflect.TypeOf(true):
			if val == "true" {
				rvalue = reflect.ValueOf(true)
			} else {
				rvalue = reflect.ValueOf(false)
			}

		case reflect.TypeOf(url.URL{}):
			s, err := strconv.Unquote(val)
			if err != nil {
				w.WriteHeader(400)
				log.Println(err)
				fmt.Fprint(w, "throw 'bad request';")
				return
			}
			location, err := url.Parse(s)
			if err != nil {
				w.WriteHeader(400)
				log.Println(err)
				fmt.Fprint(w, "throw 'bad request';")
				return
			}
			rvalue = reflect.ValueOf(*location)

		default:
			var shell = reflect.New(f.Type().In(i)).Interface()
			if err := json.NewDecoder(strings.NewReader(val)).Decode(shell); err != nil {
				w.WriteHeader(400)
				log.Println(err)
				fmt.Fprint(w, "throw 'bad request';")
				return
			}

			rvalue = reflect.ValueOf(shell).Elem()
		}

		var ElemType, ArgType = rvalue.Type(), f.Type().In(i)
		if ElemType != ArgType {
			if !(ArgType.Kind() == reflect.Interface && ElemType.Implements(ArgType)) {
				log.Println("type mismatch")
				w.WriteHeader(400)
				fmt.Fprint(w, "throw 'bad request';")
				return
			}
		}

		in = append(in, rvalue)
	}

	//Call the function.
	var results = f.Call(in)

	if len(results) == 0 {
		fmt.Fprintf(cr.writer, "//ok")
		return
	}

	//Check if an error was returned.
	if err, ok := results[len(results)-1].Interface().(error); ok && err != nil {
		log.Println(err)

		safe, ok := err.(clientsafe.Error)
		if ok {
			fmt.Fprintf(w, "throw %v;", strconv.Quote(safe.ClientError()))
			return
		}

		fmt.Fprintf(w, "throw %v;", strconv.Quote("there was an error"))
		return
	}

	var result = results[0].Interface()

	var buffer bytes.Buffer

	var encoder = json.NewEncoder(&buffer)

	type JSONEncodable interface {
		JSONEncoder() func(interface{}) ([]byte, error)
	}

	switch v := result.(type) {
	case JSONEncodable:
		var encoder = v.JSONEncoder()

		b, err := encoder(result)
		if err != nil {
			fmt.Println("rpc function could not send return value: ", err)
			return
		}
		buffer.Write(b)

	case json.Marshaler:
		encoded, err := v.MarshalJSON()
		if err != nil {
			fmt.Println("rpc function could not send return value: ", err)
			return
		}
		buffer.Write(encoded)

	case Script:
		ctx := js.NewCtx(cr.writer)
		ctx(v.GetScript())
		ctx.Flush()

	default:
		err := encoder.Encode(results[0].Interface())
		if err != nil {
			fmt.Println("rpc function could not send return value: ", err)
		}
	}

	//This is slow for arrays.
	fmt.Fprintf(w, `return %v;`, buffer.String())
	return
}
