//Package api provides a structure for designing an API for your app.
package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/qlova/seed"
	"github.com/qlova/seed/user"
)

//Endpoint is an API endpoint with a route and handler.
type Endpoint struct {
	Route   string
	Handler interface{}
}

//Design is an API design for your app.
type Design struct {
	Endpoints []Endpoint
}

//New returns a new API design.
func New() Design {
	return Design{}
}

//AddTo app.
func (d Design) AddTo(app *seed.App) {
	for _, endpoint := range d.Endpoints {
		var f = reflect.ValueOf(endpoint.Handler)
		var route = endpoint.Route
		app.Handlers[route] = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			var path = r.URL.Path
			if len(path) > len(route) && path[0:len(route)] == route {
				r.URL.Path = path[len(route):]
			}

			var u = user.CtxFromHandler(w, r)

			var in []reflect.Value

			var StartFrom = 0
			//The function can take an optional client as it's first argument.
			if f.Type().NumIn() > 0 && f.Type().In(0) == reflect.TypeOf(user.Ctx{}) {
				StartFrom = 1

				//Make the user, the first argument.
				in = append(in, reflect.ValueOf(u))
			}

			for i := StartFrom; i < f.Type().NumIn(); i++ {
				var arg = u.Arg(strconv.Itoa(i - StartFrom))

				switch f.Type().In(i).Kind() {
				case reflect.String:

					in = append(in, reflect.ValueOf(arg.String()))

				case reflect.Int:
					var number, _ = strconv.Atoi(arg.String())
					in = append(in, reflect.ValueOf(number))

				default:
					println("unimplemented callHandler for " + f.Type().String())
					return
				}
			}

			var results = f.Call(in)

			if len(results) == 0 {
				return
			}

			if len(results) == 1 {
				var buffer bytes.Buffer
				err := json.NewEncoder(&buffer).Encode(results[0].Interface())
				if err != nil {
					fmt.Println("rpc function could not send return value: ", err)
				}
				u.Execute(fmt.Sprintf(`return %v;`, buffer.String()))
				return
			}
		})
	}
}

//Endpoint creates a new endpoint for your API.
//This will panic if the handler is invalid.
func (d *Design) Endpoint(route string, handler interface{}) error {

	if handler == nil {
		panic("handler is nil")
	}

	d.Endpoints = append(d.Endpoints, Endpoint{
		Route:   route,
		Handler: handler,
	})

	return nil
}
