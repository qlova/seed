package user

import (
	"bytes"
	"net/http"

	"github.com/gorilla/websocket"
)

//AreConnected returns true if the user is connected.
func (u Ctx) AreConnected() bool {
	if u.conn == nil {
		return true
	}
	return u.conn.WriteMessage(websocket.TextMessage, []byte("{}")) == nil
}

//CtxFromSocket creates a new user from an incomming websocket request.
func CtxFromSocket(r *http.Request, w http.ResponseWriter) (Ctx, error) {
	conn, err := new(websocket.Upgrader).Upgrade(w, r, nil)
	if err != nil {
		return Ctx{}, err
	}
	return Ctx{r: r, w: w, conn: conn, buffer: new(bytes.Buffer)}, nil
}
