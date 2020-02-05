package script

import (
	"encoding/base64"
	"fmt"
	"math/big"
	"net/http"

	qlova "github.com/qlova/script"
	"github.com/qlova/seed/user"
)

//Connection is a script interface to a channel.
type Connection struct {
	Q Ctx
	qlova.Native
}

var connectionID int64 = 1

var connectionHandlers = make(map[string]func(u user.Ctx))

//Open is the JS required for q.Open(..).
const Open = `
	seed.connections = {};
	seed.open = function(path) {
		if (path in seed.connections && seed.connections[path].readyState != WebSocket.CLOSED) {
			return;
		}

		let url = new URL(path, window.location.href);
		url.protocol = url.protocol.replace('http', 'ws');

		let socket = new WebSocket(url.href);

		seed.connections[path] = socket;

		socket.addEventListener('message', function (event) {
			slave(event.data);
		});

		return socket;
	}
`

//Open opens a script <-> Go connection for continuos cross-communication.
func (q Ctx) Open(f func(u user.Ctx)) Connection {

	//Get a unique string reference for f.
	var name = base64.RawURLEncoding.EncodeToString(big.NewInt(connectionID).Bytes())

	connectionID++

	var WebSocketEndpoint = `/conn/` + name

	connectionHandlers[name] = f

	q.Require(Open)
	q.Require(Request)

	var variable = Unique()
	q.Javascript(`let %v = seed.open("%v");`, variable, WebSocketEndpoint)

	return Connection{q, q.Value(variable).Native()}
}

//ConnectionHandler handles connection connections.
//TODO rename.
func ConnectionHandler(w http.ResponseWriter, r *http.Request, call string) {
	f, ok := connectionHandlers[call]
	if !ok {
		return
	}

	var u, err = user.CtxFromSocket(r, w)
	if err != nil {
		fmt.Println(err)
		return
	}

	f(u)
}
