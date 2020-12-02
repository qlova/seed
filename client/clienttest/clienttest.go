package clienttest

import (
	"io"
	"net/http/httptest"
	"strings"

	"qlova.org/seed/client"
)

//NewRequest creates a new dummy request suitable for passing to a function that takes a client.Request as its first argument.
func NewRequest(args ...string) client.Request {
	method := "GET"
	if len(args) > 0 {
		method = args[0]
	}
	url := "/"
	if len(args) > 1 {
		url = args[1]
	}
	body := io.Reader(nil)
	if len(args) > 2 {
		body = strings.NewReader(args[2])
	}
	return client.NewRequest(
		httptest.NewRecorder(),
		httptest.NewRequest(method, url, body),
	)
}
