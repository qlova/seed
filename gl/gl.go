//+build ignore

//Package gl is currently disabled until further notice.
package gl

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
)

//Context is a gl context for rendering to the screen.
type Context struct {
	*context

	//The number of initialised objects.
	buffers, programs, shaders, attributes int32

	//Bitfield types.
	ColorBufferBit, DepthBufferBit, StencilBufferBit BitField

	VertexShader, FragmentShader ShaderType

	ArrayBuffer, ElementArrayBuffer BufferTarget

	StaticDraw UsagePattern

	Float DataType

	Triangles DrawMode
}

type context struct {
	mutex   sync.RWMutex
	clients []*websocket.Conn
}

func sprint(i interface{}) string {
	return fmt.Sprint(i)
}

//NewContext returns a gl context from the specified seed.
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

	const (
		ColorBufferBit = iota
		DepthBufferBit
		StencilBufferBit
	)
	const (
		VertexShader = iota
		FragmentShader
	)
	const (
		ArrayBuffer = iota
		ElementArrayBuffer
	)
	const (
		StaticDraw = iota
	)
	const (
		Float = iota
	)
	const (
		Triangles = iota
	)

	element.OnReady(func(q script.Ctx) {
		q.Javascript(`let gl = ` + element.Ctx(q).Element() + `.getContext("webgl");`)
		q.Javascript(`var socket = new WebSocket(((window.location.protocol === "https:") ? "wss://" : "ws://") + window.location.host + "/gl/` + element.ID() + `"); socket.binaryType = 'arraybuffer';`)
		q.Javascript(`socket.onopen = function() {`)
		q.Javascript(`	console.log('connected!');`)
		q.Javascript(`};`)
		q.Javascript(`
		
		function OverrideRingBuffer(size){
			this.size = size;
			this.head = 0;
			this.buffer = new Array();
		};
		
		OverrideRingBuffer.prototype.push = function(value){      
			if(this.head >= this.size) this.head -= this.size;    
			this.buffer[this.head] = value;
			this.head++;
		};
		
		OverrideRingBuffer.prototype.getAverage = function(){
			if(this.buffer.length === 0) return 0;
		
			var sum = 0;    
		
			for(var i = 0; i < this.buffer.length; i++){
				sum += this.buffer[i];
			}    
		
			return (sum / this.buffer.length).toFixed(1);
		};

		function FpsCounter(){
			this.count = 0;
			this.fps = 0;
			this.prevSecond;  
			this.minuteBuffer = new OverrideRingBuffer(60);
		}
		
		FpsCounter.prototype.update = function(){
			if (!this.prevSecond) {     
				this.prevSecond = new Date().getTime();
					this.count = 1;
			}
			else {
				var currentTime = new Date().getTime();
				var difference = currentTime - this.prevSecond;
				if (difference > 1000) {      
					this.prevSecond = currentTime;
					this.fps = this.count; 
					this.minuteBuffer.push(this.count);
					this.count = 0;
				}
				else{
					this.count++;
				}
			}   
			console.log(fpsCounter.getCountPerSecond());
		};
		
		FpsCounter.prototype.getCountPerMinute = function(){
			return this.minuteBuffer.getAverage();
		};
		
		FpsCounter.prototype.getCountPerSecond = function(){
			return this.fps;
		}; var fpsCounter = new FpsCounter();;`)
		q.Javascript(`
			let buffers = [null];
			let programs = [null];
			let shaders = [null];
			let attributes = [null];
			let decoder = new TextDecoder("utf-8");
		`)
		q.Javascript(`socket.onmessage = function(msg) {`)
		q.Javascript(`
			
			let cmd = (new Uint16Array(msg.data, 0, 1))[0];
			let args;
			
			switch (cmd) {
				case ` + sprint(activeTexture) + `: {
					args = (new Int32Array(msg.data, 4, 1))
					gl.activeTexture(args[0])
					break;
				}
				case ` + sprint(attachShader) + `: {
					args = (new Int32Array(msg.data, 4, 2))
					gl.attachShader(programs[args[0]], shaders[args[1]])
					break;
				}
				case ` + sprint(bindBuffer) + `: {
					let target = (new Uint8Array(msg.data, 4, 1))[0];
					switch (target) {
					case ` + sprint(ArrayBuffer) + `:
						target = gl.ARRAY_BUFFER;
						break;
					case ` + sprint(ElementArrayBuffer) + `:
						target = gl.ELEMENT_ARRAY_BUFFER;
						break;
					default:
						console.error("invalid BufferTarget: ", target);
						socket.close();
						return;
					}
					let buffer = buffers[(new Int32Array(msg.data, 8, 1))[0]];
					gl.bindBuffer(target, buffer)
					break;
				}
				case ` + sprint(bufferData) + `: {
					args = (new Uint8Array(msg.data, 4, 2))
					target = args[0];

					switch (target) {
					case ` + sprint(ArrayBuffer) + `:
						target = gl.ARRAY_BUFFER;
						break;
					case ` + sprint(ElementArrayBuffer) + `:
						target = gl.ELEMENT_ARRAY_BUFFER;
						break;
					default:
						console.error("invalid BufferTarget: ", target);
						socket.close();
						return;
					}

					let usage = args[1];
					switch (usage) {
					case ` + sprint(StaticDraw) + `:
						usage = gl.STATIC_DRAW;
						break;
					default:
						console.error("invalid UsagePattern: ", target);
						socket.close();
						return;
					}

					let length = (new Int32Array(msg.data, 8, 1))[0];
					let array = new Int32Array(msg.data, 12, length);

					gl.bufferData(target, array, usage);
					break;
				}
				case ` + sprint(clearColor) + `: {
					args = (new Float32Array(msg.data, 4, 4))
					gl.clearColor(args[0], args[1], args[2], args[3]);
					break;
				}
				case ` + sprint(clear) + `: {
					args = (new Uint8Array(msg.data, 4, 1))
					switch (args[0]) {
						case ` + sprint(ColorBufferBit) + `:
							gl.clear(gl.COLOR_BUFFER_BIT);
							fpsCounter.update();
							break;
						case ` + sprint(DepthBufferBit) + `:
							gl.clear(gl.DEPTH_BUFFER_BIT);
							break;
						case ` + sprint(StencilBufferBit) + `:
							gl.clear(gl.STENCIl_BUFFER_BIT);
							break;
						default:
							console.error("invalid BufferBit: ", args[0]);
							socket.close();
							return;
					}	
					break;
				}
				case ` + sprint(compileShader) + `: {
					args = (new Int32Array(msg.data, 4, 1))
					gl.compileShader(shaders[args[0]]);
					break;
				}
				case ` + sprint(createBuffer) + `: {
					buffers.push(gl.createBuffer());
					break;
				}
				case ` + sprint(createShader) + `: {
					args = (new Uint8Array(msg.data, 4, 1)) 
					switch (args[0]) {
						case ` + sprint(VertexShader) + `:
							shaders.push(gl.createShader(gl.VERTEX_SHADER));
							break;
						case ` + sprint(FragmentShader) + `:
							shaders.push(gl.createShader(gl.FRAGMENT_SHADER));
							break;
						default:
							console.error("invalid ShaderType: ", args[0]);
							socket.close();
							return;
					}
					
					break;
				}
				case ` + sprint(createProgram) + `: {
					programs.push(gl.createProgram());
					break;
				}
				case ` + sprint(drawArrays) + `: {
					let mode = (new Uint8Array(msg.data, 4, 1))[0];

					switch (mode) {
						case ` + sprint(Triangles) + `:
							mode = gl.TRIANGLES;
							break;
						default:
							console.error("invalid DrawType: ", mode);
							socket.close();
							return;
					}

					args = (new Int32Array(msg.data, 8, 2));
					let first = args[0];
					let count = args[1];
					gl.drawArrays(mode, first, count);
					break;
				}
				case ` + sprint(enableVertexAttribArray) + `: {
					args = (new Int32Array(msg.data, 4, 1))
					gl.enableVertexAttribArray(attributes[args[0]]);
					break;
				}
				case ` + sprint(getAttribLocation) + `: {
					args = (new Int32Array(msg.data, 4, 2))
					let program = programs[args[0]];
					let length = args[1];
					let name = new Uint8Array(msg.data, 12, length);
					attributes.push(gl.getAttribLocation(program, decoder.decode(name)));
					break;
				}
				case ` + sprint(linkProgram) + `: {
					args = (new Int32Array(msg.data, 4, 1))
					gl.linkProgram(programs[args[0]]);
					break;
				}
				case ` + sprint(shaderSource) + `: {
					args = (new Int32Array(msg.data, 4, 2))
					let shader = shaders[args[0]];
					let length = args[1];
					let source = new Uint8Array(msg.data, 12, length);
					gl.shaderSource(shader, decoder.decode(source));
					break;
				}
				case ` + sprint(useProgram) + `: {
					args = (new Int32Array(msg.data, 4, 1))
					gl.useProgram(programs[args[0]]);
					break;
				}
				case ` + sprint(vertexAttribPointer) + `: {
					args = (new Int32Array(msg.data, 4, 1))
					let attribute = attributes[args[0]];
					args = (new Uint8Array(msg.data, 8, 3))
					let size = args[0];
					let type = args[1];

					switch (type) {
						case ` + sprint(Float) + `:
							type = gl.FLOAT;
							break;
						default:
							console.error("invalid DataType: ", args[0]);
							socket.close();
							return;
					}

					let normalized = args[2];
					args = (new Int32Array(msg.data, 12, 2))
					let stride = args[0];
					let offset = args[1];

					gl.vertexAttribPointer(attribute, size, type, normalized, stride, offset);
					break;
				}
				default:
					console.error("invalid cmd: ", cmd);
					socket.close();

			}
		`)
		q.Javascript(`};`)
	})

	return Context{
		context: &ctx,

		ColorBufferBit:   ColorBufferBit,
		DepthBufferBit:   DepthBufferBit,
		StencilBufferBit: StencilBufferBit,

		VertexShader:   VertexShader,
		FragmentShader: FragmentShader,

		ArrayBuffer:        ArrayBuffer,
		ElementArrayBuffer: ElementArrayBuffer,

		StaticDraw: StaticDraw,

		Float: Float,

		Triangles: Triangles,
	}
}

func (ctx *Context) send(cmd uint32, args ...interface{}) {
	ctx.mutex.RLock()
	defer ctx.mutex.RUnlock()

	var buffer bytes.Buffer

	binary.Write(&buffer, binary.LittleEndian, cmd)

	for _, arg := range args {
		if s, ok := arg.(string); ok {
			binary.Write(&buffer, binary.LittleEndian, int32(len(s)))
			buffer.WriteString(s)
			continue
		}
		if reflect.TypeOf(arg).Kind() == reflect.Slice {
			binary.Write(&buffer, binary.LittleEndian, int32(reflect.ValueOf(arg).Len()))
		}
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

//ActiveTexture specifies the active texture unit.
func (ctx *Context) ActiveTexture(texture TextureUnit) {
	ctx.send(activeTexture, texture)
}

//BufferData creates a buffer in memory and initializes it with array data.
//If no array is provided, the contents of the buffer is initialized to 0.
func (ctx *Context) BufferData(target BufferTarget, array interface{}, usage UsagePattern) {
	ctx.send(bufferData, target, usage, byte(0), byte(0), array)
}

//ClearColor clears the color of the buffer.
func (ctx *Context) ClearColor(red, green, blue, alpha float32) {
	ctx.send(clearColor, red, green, blue, alpha)
}

//Clear clears the mask.
func (ctx *Context) Clear(mask BitField) {
	ctx.send(clear, mask)
}

//CreateBuffer creates and initializes a Buffer.
func (ctx *Context) CreateBuffer() Buffer {
	ctx.buffers++
	ctx.send(createBuffer)
	return Buffer{ctx, ctx.buffers, 0}
}

//CreateShader returns an empty vertex or fragment shader object based on the type specified.
func (ctx *Context) CreateShader(T ShaderType) Shader {
	ctx.shaders++
	ctx.send(createShader, T)
	return Shader{ctx, T, ctx.shaders, ""}
}

//CreateProgram creates an empty Program  to which shaders can be bound.
func (ctx *Context) CreateProgram() Program {
	ctx.programs++
	ctx.send(createProgram)
	return Program{ctx, ctx.programs, Shader{}, Shader{}, nil}
}

//DrawArrays renders geometric primitives from bound and enabled vertex data.
func (ctx *Context) DrawArrays(mode DrawMode, first, count int32) {
	ctx.send(drawArrays, mode, byte(0), byte(0), byte(0), first, count)
}
