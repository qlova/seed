//Package api provides a structure for designing an API for your app.
package api

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientsafe"
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
	handlers map[string]http.Handler
}

func (d Design) And(more ...seed.Option) seed.Option {
	return seed.And(d, more...)
}

//AddTo app.
func (d Design) AddTo(c seed.Seed) {
	var data data
	c.Load(&data)

	if data.handlers == nil {
		data.handlers = make(map[string]http.Handler)
		c.Save(data)
	}

	for _, endpoint := range d.endpoints {
		var route = endpoint.route
		var handler = endpoint.handler
		data.handlers[route] = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var path = r.URL.Path
			if len(path) > len(route) && path[0:len(route)] == route {
				r.URL.Path = path[len(route):]
			}

			var cr = client.NewRequest(w, r)

			var in = []reflect.Value{reflect.ValueOf(cr)}

			//Call the function.
			var results = reflect.ValueOf(handler).Call(in)

			if len(results) == 0 {
				return
			}

			//Check if an error was returned.
			if err, ok := results[len(results)-1].Interface().(error); ok && err != nil {
				log.Println(err)

				switch e := err.(type) {
				case clientsafe.Error:
					fmt.Fprintf(w, "%v", e.ClientError())
				case client.Redirect:
					http.Redirect(w, r, string(e), http.StatusSeeOther)
				}

				return
			}

			return
		})
	}
}
