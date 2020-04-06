//Package api provides a structure for designing an API for your app.
package api

import (
	"net/http"
	"reflect"

	"github.com/qlova/seed"
	"github.com/qlova/seed/user"
)

type Option func(*Design)

//Endpoint is an API endpoint with a route and handler.
func Endpoint(route string, handler interface{}) Option {
	return func(d *Design) {
		d.endpoints = append(d.endpoints, endpoint{
			route:   route,
			handler: handler,
		})
	}
}

type endpoint struct {
	route   string
	handler interface{}
}

//Design is an API design for your app.
type Design struct {
	endpoints []endpoint
}

//New returns a new API design.
func New(options ...Option) Design {
	d := Design{}
	for _, o := range options {
		o(&d)
	}
	return d
}

type data struct {
	seed.Data

	handlers map[string]http.Handler
}

func (d Design) And(more ...seed.Option) seed.Option {
	return seed.And(d, more...)
}

//AddTo app.
func (d Design) AddTo(c seed.Seed) {
	var data data
	c.Read(&data)

	if data.handlers == nil {
		data.handlers = make(map[string]http.Handler)
		c.Write(data)
	}

	for _, endpoint := range d.endpoints {
		var route = endpoint.route
		var handler = endpoint.handler
		data.handlers[route] = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var path = r.URL.Path
			if len(path) > len(route) && path[0:len(route)] == route {
				r.URL.Path = path[len(route):]
			}

			var u = user.CtxFromHandler(w, r)

			//Create API struct
			T := reflect.TypeOf(handler)
			if T.NumIn() == 2 {
				var args = reflect.New(T.In(1))

				for i := 0; i < args.Elem().NumField(); i++ {
					var field = args.Elem().Type().Field(i)
					switch field.Type {
					case reflect.TypeOf(""):
						args.Elem().Field(i).Set(reflect.ValueOf(u.Arg(field.Name).String()))
					}

				}

				var in = []reflect.Value{reflect.ValueOf(u), args.Elem()}

				reflect.ValueOf(handler).Call(in)
			} else if T.NumIn() == 1 {
				var in = []reflect.Value{reflect.ValueOf(u)}

				reflect.ValueOf(handler).Call(in)
			}
		})
	}

	/*
		var f = reflect.ValueOf(endpoint.handler)
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
	}*/
}
