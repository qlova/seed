//Package user allows communication with users from Go code.
package user

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

//Ctx is a user-context, meaning a current connection to a user of you're application.
type Ctx struct {
	w http.ResponseWriter
	r *http.Request

	conn *websocket.Conn
	buff *bytes.Buffer

	buffer *bytes.Buffer
}

//CtxFromHandler returns a user ctx from the request and responsewriter inside an http Handler.
func CtxFromHandler(w http.ResponseWriter, r *http.Request) Ctx {
	return Ctx{w: w, r: r, buffer: new(bytes.Buffer)}
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
func (u Ctx) Execute(script string) {
	if u.w != nil {
		fmt.Fprint(u.w, script)
	}
	if u.conn != nil {
		u.conn.WriteMessage(websocket.TextMessage, []byte(script))
	}
}
