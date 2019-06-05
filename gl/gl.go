package gl

import "sync"
import "github.com/qlova/seed"
import "github.com/gorilla/websocket"
import "encoding/binary"
import "net/http"
import "bytes"
import "runtime"
import "fmt"

type BitField uint8

type Context struct {
	*context

	ColorBufferBit BitField
}

type context struct {
	mutex   sync.RWMutex
	clients []*websocket.Conn
}

func sprint(i interface{}) string {
	return fmt.Sprint(i)
}

func NewContext(element seed.Seed) Context {

	var ctx context

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	element.AddHandler(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/gl/"+element.ID() {
			conn, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				//log.Println(err)
				return
			}

			ctx.mutex.Lock()
			ctx.clients = append(ctx.clients, conn)
			ctx.mutex.Unlock()
		}
	})

	const ColorBufferBit = 0

	element.OnReady(func(q seed.Script) {
		q.Javascript(`let gl = ` + element.Script(q).Element() + `.getContext("webgl");`)
		q.Javascript(`var socket = new WebSocket(((window.location.protocol === "https:") ? "wss://" : "ws://") + window.location.host + "/gl/` + element.ID() + `"); socket.binaryType = 'arraybuffer';`)
		q.Javascript(`socket.onopen = function() {`)
		q.Javascript(`	console.log('connected!');`)
		q.Javascript(`};`)
		q.Javascript(`socket.onmessage = function(msg) {`)
		q.Javascript(`
		
			let cmd = (new Uint16Array(msg.data, 0, 1))[0];
			let args;
			
			switch (cmd) {
				case ` + sprint(clearColor) + `:
					args = (new Float32Array(msg.data, 4, 4))
					gl.clearColor(args[0], args[1], args[2], args[3]);
					break;
					
				case ` + sprint(clear) + `:
					args = (new Uint8Array(msg.data, 4, 1))
					switch (args[0]) {
						case ` + sprint(ColorBufferBit) + `:
							gl.clear(gl.COLOR_BUFFER_BIT);
							break;
					}
					
					break;
				
			}
		`)
		q.Javascript(`};`)
	})

	return Context{
		context: &ctx,

		ColorBufferBit: ColorBufferBit,
	}
}

func (ctx *Context) Send(cmd uint32, args ...interface{}) {
	ctx.mutex.RLock()
	defer ctx.mutex.RUnlock()

	var buffer bytes.Buffer

	binary.Write(&buffer, binary.LittleEndian, cmd)

	for _, arg := range args {
		binary.Write(&buffer, binary.LittleEndian, arg)
	}

	var message = buffer.Bytes()

	for len(ctx.clients) == 0 {
		ctx.mutex.RUnlock()
		runtime.Gosched()
		ctx.mutex.RLock()
	}

	for _, client := range ctx.clients {
		client.WriteMessage(websocket.BinaryMessage, message)
	}
}

func (ctx *Context) ClearColor(red, green, blue, alpha float32) {
	ctx.Send(clearColor, red, green, blue, alpha)
}

func (ctx *Context) Clear(mask BitField) {
	ctx.Send(clear, mask)
}
