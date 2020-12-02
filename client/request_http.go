package client

import (
	"io"
	"mime/multipart"
	"net/http"
	"regexp"
	"strings"
	"time"
)

//Cookie is an client-request associated value that is encrypted by default.
type Cookie struct {
	Name   string
	MaxAge time.Duration
}

//NewCookie creates a new cookie with the given name.
func NewCookie(name string) Cookie {
	return Cookie{
		Name:   name,
		MaxAge: time.Hour * 24 * 60,
	}
}

//Redirect is a special error-type that signifies that a redirect is required to complete the request.
type Redirect string

func (r Redirect) Error() string {
	return "see other " + string(r)
}

//Request holds the metadata about an incomming request from the client.
type Request struct {
	Path string

	Response io.Writer

	writer  http.ResponseWriter
	request *http.Request
}

//NewRequest returns a new request from the given values.
func NewRequest(w http.ResponseWriter, r *http.Request) Request {
	return Request{
		Path:     r.URL.Path,
		Response: w,

		writer:  w,
		request: r,
	}
}

//Arg returns the named query value with the given name.
func (cr Request) Arg(name string) string {
	return cr.request.FormValue(name)
}

//Serve serves a http.Handler as a response to this request.
func (cr Request) Serve(handler http.Handler) {
	handler.ServeHTTP(cr.writer, cr.request)
}

//Set sets the value of a cookie associated with requests by this client.
func (cr Request) Set(c Cookie, value string) {
	var cookie = &http.Cookie{
		Name: c.Name,

		Expires:  time.Now().Add(c.MaxAge),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		Secure:   !cr.Local(),

		Value: Encrypt([]byte(value)),
	}
	cr.request.AddCookie(cookie)
	http.SetCookie(cr.writer, cookie)
}

//Get gets the value of a cookie associated with requests by this client.
func (cr Request) Get(c Cookie) string {
	a, err := cr.request.Cookie(c.Name)
	if err != nil {
		return ""
	}
	return string(Decrypt(a.Value))
}

var intranet, _ = regexp.Compile(`(^192\.168\.([0-9]|[0-9][0-9]|[0-2][0-5][0-5])\.([0-9]|[0-9][0-9]|[0-2][0-5][0-5]):.*$)`)

//Local returns true if the user is connecting from localhost.
func (cr Request) Local() (local bool) {
	r := cr.request

	if r.Header.Get("X-Real-IP") != "" || r.Header.Get("X-Forwarded-For") != "" {
		return false
	}

	local = strings.Contains(r.RemoteAddr, "[::1]") || strings.Contains(r.RemoteAddr, "127.0.0.1")
	if local {
		return
	}
	if intranet.Match([]byte(r.RemoteAddr)) {
		local = true
	}

	split := strings.Split(r.Host, ":")
	if len(split) == 0 {
		local = false
	} else {
		if split[0] != "localhost" {
			local = false
		}
	}

	return
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
