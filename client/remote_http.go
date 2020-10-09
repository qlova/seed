package client

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"

	"qlova.org/seed/js"
)

//Remote acts as a 'remote-controller' to a 'remote' client.
type Remote struct {
	writer  http.ResponseWriter
	request *http.Request
}

//Set sets the value of a cookie on the client.
func (client Remote) Set(c http.Cookie, v string) {
	c.Value = v
	http.SetCookie(client.writer, &c)
}

//Get gets the value of a cookie on the client.
func (client Remote) Get(c http.Cookie) string {
	a, err := client.request.Cookie(c.Name)
	if err != nil {
		return ""
	}
	return a.Value
}

//Execute remotely writes the given scripts to the client, scheduled for execution when the Remote is closed.
func (client Remote) Execute(do ...Script) error {
	if client.writer != nil {
		ctx := js.NewCtx(client.writer)
		ctx(NewScript(do...))
		ctx.Flush()
		return nil
	}

	return errors.New("nil writer")
}

//Stream is an incomming data steam from the client.
type Stream struct {
	head *multipart.FileHeader
	file multipart.File
}

//Name returns the name of the steam, if a file is being sent, this is the name of the file.
func (s Stream) Name() string {
	if s.head == nil {
		return ""
	}

	return s.head.Filename
}

//Read implements io.Reader
func (s Stream) Read(b []byte) (int, error) {
	if s.head == nil {
		return 0, io.EOF
	}
	return s.file.Read(b)
}

//Close implements io.Closer
func (s Stream) Close() error {
	if s.head == nil {
		return nil
	}
	return s.file.Close()
}

//Size returns the expected size of the stream.
func (s Stream) Size() int64 {
	if s.head == nil {
		return 0
	}

	return s.head.Size
}
