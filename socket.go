//+build !wasm

package seed

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var singleLocalConnection = false

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var localSockets = make(map[string]*websocket.Conn)
var reloading = false

func socket(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	localSockets[r.RemoteAddr] = c

	reloading = false

	for {
		_, message, err := c.ReadMessage()

		if bytes.Equal(message, []byte("I'll be back")) {
			c.WriteMessage(websocket.TextMessage, []byte("window.LocalhostWebsocket.onclose = function() {};"))
			delete(localSockets, r.RemoteAddr)
			localClients--
			return
		}

		if err != nil {
			singleLocalConnection = localClients == 1

			if singleLocalConnection && !reloading {
				cleanup()
				os.Exit(0)
			} else {
				delete(localSockets, r.RemoteAddr)
				localClients--
				return
			}
		}
	}
}
