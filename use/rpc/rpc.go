//Package rpc provides useful interfaces and types for RPC requests and their implementations.
package rpc

import (
	"context"
	"io"
	"net/http"
)

//Stream is a readable stream with an implementation-defined backing mechanism.
type Stream interface {
	io.ReadCloser

	//Name should return the client's name for this stream.
	//When the stream is a file, this should be the filename.
	Name() string

	//Size returns the expected size of the stream.
	Size() int64
}

//Scanner is any type that can scan itself from an input value.
type Scanner interface {
	Scan(input interface{}) error
}

//Validator types can validate their contents.
type Validator interface {

	//Validate returns an error if the contents of the
	//type are invalid in any way.
	Validate() error
}

//Request is a request from the client for some kind of procedure, action or function to
//execute on the server. The Request type is optional for rpc handlers to include and
//is only necessary when the client needs to be identified from the context of their request.
type Request struct {
	context.Context

	//Request has fields defined in the http package, however, an rpc.Request need not be
	//made over http, in which case a different transport will fill these values where
	//appropriate.
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}

//RequestScanner is any type that can scan itself from a Request type.
//These types will automatically be filled by the Request and are not
//explicitly Passed by the client.
type RequestScanner interface {
	ScanRequest(Request) error
}
