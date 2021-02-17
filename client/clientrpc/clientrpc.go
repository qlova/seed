//Package clientrpc provides helpers for client.Run implementations.
package clientrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"time"

	"qlova.org/seed/client/clientsafe"
	"qlova.org/seed/use/js"
)

//Cookie is an client-request associated value that should be encrypted by default.
type Cookie struct {
	Name   string
	MaxAge time.Duration
}

//Request is a request that is hitting the clientrpc point.
type Request interface {
	Path() string
	Arg(name string) string

	Header(key string) string
	SetHeader(key, value string)

	Get(Cookie) string
	Set(Cookie, string)

	Writer() io.Writer
}

//Scanner is any type that can scan itself from an input.
type Scanner interface {
	Scan(input interface{}) error
}

//Validator can validate itself.
type Validator interface {
	Validate() error
}

//RequestScanner is any type that can scan itself from a Request.
type RequestScanner interface {
	ScanRequest(Request) error
}

//Stream is a readable stream sent from the client.
type Stream interface {
	io.ReadCloser

	Name() string
	Size() int64
}

//Context for a clientrpc.Call
type Context struct {
	Request Request
}

//Return returns a javascript function body from a Call result.
func (ctx Context) Return(result interface{}, err error) {
	w := ctx.Request.Writer()

	if err != nil {
		safe, ok := err.(clientsafe.Error)
		if ok {
			fmt.Fprintf(w, "throw %v;", strconv.Quote(safe.ClientError()))
			return
		}

		fmt.Fprintf(w, "throw %v;", strconv.Quote("there was an error"))
		return
	}

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

	case js.AnyScript:
		ctx := js.NewCtx(w)
		ctx(v.GetScript())
		ctx.Flush()

	default:
		err := encoder.Encode(result)
		if err != nil {
			fmt.Println("rpc function could not send return value: ", err)
		}
	}

	//This is slow for arrays.
	fmt.Fprintf(w, `return %v;`, buffer.String())
}

/*
Call takes the Go function 'fn', originating from 'req', and deserialises any
string arguments as required and returns the result.
If any deserialisation errors occur or the Go function returns a non-nil error
then that error is returned by Call.
*/
func (ctx Context) Call(fn interface{}, args ...interface{}) (interface{}, error) {
	f := reflect.ValueOf(fn)

	var in = make([]reflect.Value, 0, f.Type().NumIn())

	var StartFrom = 0

	//The function may take an optional Request as it's first argument.
	if f.Type().NumIn() > 0 && f.Type().In(0) == reflect.TypeOf([0]Request{}).Elem() {
		StartFrom = 1

		//Pass the Request as the first argument.
		in = append(in, reflect.ValueOf(ctx.Request))
	}

	//The function can take an optional RequestScanner as it's first argument.
	if f.Type().NumIn() > 0 && reflect.PtrTo(f.Type().In(0)).Implements(reflect.TypeOf([0]RequestScanner{}).Elem()) {
		StartFrom = 1

		var zero = reflect.New(f.Type().In(0))

		if err := zero.Interface().(RequestScanner).ScanRequest(ctx.Request); err != nil {
			return nil, err
		}

		//Make the user, the first argument.
		in = append(in, zero.Elem())

	}

	var skip = 0

	//Parse each argument as JSON.
	for i := StartFrom; i < f.Type().NumIn(); i++ {

		var val = args[i-StartFrom-skip]

		var rvalue reflect.Value

		switch f.Type().In(i) {

		case reflect.TypeOf([0]func() time.Time{}).Elem():
			rvalue = reflect.ValueOf(func() time.Time {
				return time.Now().In(time.UTC)
			})
			skip++

		default:
			var err error
			rvalue, err = Scan(f.Type().In(i), val)
			if err != nil {
				return nil, err
			}
		}

		if f.Type().In(i).Implements(reflect.TypeOf([0]Validator{}).Elem()) {
			if err := rvalue.Interface().(Validator).Validate(); err != nil {
				return nil, err
			}
		}

		in = append(in, rvalue)
	}

	//Call the function.
	var results = f.Call(in)

	if len(results) == 0 {
		return nil, nil
	}

	//Check if an error was returned.
	if err, ok := results[len(results)-1].Interface().(error); ok && err != nil {
		return nil, err
	}

	return results[0].Interface(), nil
}
