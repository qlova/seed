package client

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"qlova.org/seed/client/clientrpc"
)

//Requester is any type that can load itself from a Request.
type Requester interface {
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

	var cr = NewRequest(w, r)

	ctx := clientrpc.Context{Request: cr}

	var args []interface{}

	for i := 0; i < f.Type().NumIn(); i++ {
		var key = string(rune('a' + i))

		file, header, err := cr.request.FormFile(key)
		if err == nil {
			args = append(args, Stream{header, file})
			continue
		}

		s := cr.request.FormValue(key)
		if strings.HasPrefix(s, "\"") {
			s, err = strconv.Unquote(s)
			if err != nil {
				ctx.Return(nil, err)
				return
			}
		}

		args = append(args, s)
	}

	i, err := ctx.Call(f.Interface(), args...)
	if err != nil {
		log.Println(err)
	}

	ctx.Return(i, err)
	return
}
