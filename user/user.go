//Package user allows communication with users from Go code.
package user

import (
	"bytes"
	"net/http"
	"regexp"
	"strings"

	"github.com/qlova/seed/js"
)

//Handler is a user handler.
type Handler func(Ctx)

//Ctx is a user-context, meaning a current connection to a user of you're application.
type Ctx struct {
	w http.ResponseWriter
	r *http.Request

	buffer *bytes.Buffer
}

//CtxFromHandler returns a user ctx from the request and responsewriter inside an http Handler.
func CtxFromHandler(w http.ResponseWriter, r *http.Request) Ctx {
	return Ctx{w: w, r: r, buffer: new(bytes.Buffer)}
}

//Valid returns true if the context is valid.
func (u Ctx) Valid() bool {
	return u.r != nil
}

//ResponseWriter returns the ResponseWriter passed to the Ctx when it was created.
func (u Ctx) ResponseWriter() http.ResponseWriter {
	return u.w
}

//Request returns the Request passed to the Ctx when it was created.
func (u Ctx) Request() *http.Request {
	return u.r
}

//Execute sends and evaluates the provided javascript.
func (u Ctx) Execute(script js.Script) {
	if u.w != nil {
		js.NewCtx(u.w)(script)
	}
}

var intranet, _ = regexp.Compile(`(^192\.168\.([0-9]|[0-9][0-9]|[0-2][0-5][0-5])\.([0-9]|[0-9][0-9]|[0-2][0-5][0-5]):.*$)`)

//Local returns true if the user is connecting from localhost.
func (u Ctx) Local() (local bool) {
	r := u.Request()

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
