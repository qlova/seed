package seed

import (
	"net/http"
	"log"
	"os"
)

import "github.com/gorilla/websocket"

var SingleLocalConnection = false

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true //r.Header.Get("Origin") == "https://realmoforder.com"
	},	
}

func socket(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	
	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			if SingleLocalConnection {
				os.Exit(0)
			} else {
				return
			}
		}
	}
}
