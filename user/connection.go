package user

import (
	"bytes"
	"net/http"

	"github.com/gobwas/ws"
)

//FromConnection creates a new user from an incomming websocket request.
func FromConnection(r *http.Request, w http.ResponseWriter) (User, error) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		return User{}, err
	}
	return User{user: user{
		Request: r,

		Update: Update{
			Response:     new(string),
			Document:     make(map[string]string),
			LocalStorage: make(map[string]string),
			script:       new(bytes.Buffer),
		},
	}, conn: conn, buff: new(bytes.Buffer)}, nil
}
