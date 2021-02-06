package client

import (
	"bytes"
	"fmt"
	"sort"

	"qlova.org/seed"
	"qlova.org/seed/use/js"
	"qlova.org/seed/use/wasm"
)

type Renderer func(root seed.Seed) []byte

var renderers, rootRenderers []Renderer

func RegisterRenderer(r Renderer) {
	renderers = append(renderers, r)
}

func RegisterRootRenderer(r Renderer) {
	rootRenderers = append([]Renderer{r}, rootRenderers...)
}

func init() {
	RegisterRootRenderer(func(seed.Seed) []byte {
		return []byte(wasm.InstantiateStreaming)
	})
}

func render(child seed.Seed) []byte {
	var b bytes.Buffer
	var d Data
	child.Load(&d)

	//Deterministic render.
	keys := make([]string, 0, len(d.On))
	for i := range d.On {
		keys = append(keys, string(i))
	}
	sort.Strings(keys)

	for _, event := range keys {
		handler := d.On[event]
		js.NewCtx(&b)(func(q js.Ctx) {
			fmt.Fprintf(q, `seed.on(%v, "%v", async function() {`, Element(child), event)
			handler.GetScript()(q)
			q(`});`)
		})
	}

	for _, child := range child.Children() {
		b.Write(render(child))
	}

	if _, ok := d.On["ready"]; ok {
		js.NewCtx(&b)(func(q js.Ctx) {
			fmt.Fprintf(q, `await %[1]v.onready();`, Element(child))
		})
	}

	return b.Bytes()
}

//Render renders the Javascript attached to this seed and its children.
func Render(root seed.Seed) []byte {
	var b bytes.Buffer

	b.WriteString(`seed = {}; seeds = {}; c = seed; s = seeds;
seed.production = (location.hostname != "localhost" && location.hostname != "127.0.0.1" && location.hostname != "[::]" && location.hostname != "[::1]");

seed.op = function(element, func, propagate) {
	let handler = async function(event) {
		await func(event);
	};
	
	let moved = false;
	let point = [0, 0];
	
	element.ontouchstart = function(e) {
		var changedTouch = event.changedTouches[0];
			point[0]  = changedTouch.clientX;
			point[1]  = changedTouch.clientY;
	};
	
	element.ontouchmove = function(event) {
		var changedTouch = event.changedTouches[0];
		var elem = document.elementFromPoint(changedTouch.clientX, changedTouch.clientY);
					
		if (elem != event.target) moved = true;
					
		let a = changedTouch.clientX - point[0];
		let b = changedTouch.clientY - point[1];
		if ((a*a + b*b) > 50*50) moved = true;
	};
	
	element.ontouchend = async function(ev) {
		if (ev.stopPropagation && !propagate) ev.stopPropagation(); 
		ev.preventDefault(); 
		if (moved) {
			moved = false; 
			return; 
		}
		ev = ev.changedTouches[0];
		await handler(ev);
	};

	element.onclick = handler;
}

seed.title = document.title;

seed.active = null;

//seed.report is the error handling function. Pass the current element for 'OnError' based error handling.
seed.report = async function(err, element) {
	if (err == "") return; //ignore empty errors.

	if (!element) element = seed.active;

	if (element) {
		while (true) {
			if (element.onerror) {
				try {
					await element.onerror(err);
					break;
				} catch(e) {
					err = e;
				}
			}
			element = element.parentElement || element.parent;
			if (!element) break;
		}
	}
	
	console.error(err);
};

seed.globals = {};

//seed.on registers an event handler for the given element and the given event name.
//if 'id' is provided, then the handler can be removed by calling seed.off(element, event, id)
seed.on = function(element, event, handler, id) {
	if (id != null) {
		if (!element["don"+event]) element["don"+event] = {};

		element["don"+event][id] = handler;
		return;
	}

	let f = async function(ev) {
		seed.active = element;
		try {
			await handler(ev);

			let dynamic_handlers = element["don"+event];
			if (dynamic_handlers) {
				for (const h in dynamic_handlers) {
					await dynamic_handlers[h]();
				}
			}
			
		} catch(e) {
			//prevent infinite error-handling loops.
			if (event == "error") {
				throw e;
			} else {
				seed.report(e, element);
			}
		}
	};

	if (event == "press") {
		seed.op(element, f);
	} else if (event == "enter") {
		element["onkeypress"] = async function(ev) {
			if (ev.keyCode == 13 || ev.which == 13) { 
				await f(ev);
			}
		};
	} else {
		element["on"+event] = f;
	}
};

seed.off = function(element, event, id) {
	if (!element["don"+event]) return;

	delete element["don"+event][id];
}

seed.get_template = function(template, id) {

	let element;
	if (id[0] == ".") {
		element = template.content.querySelector(id);
	} else {
		element = template.content.getElementById(id);
	}

	if (element) return element;

	let templates = template.content.querySelectorAll('template');
	for (let template of templates) {
		element = seed.get_template(template, id);
		if (element) {
			if (seed.get.cache) seed.get.cache[id] = element;
			if (!element.parentElement) {
				element.parent = template;
			}
			return element;
		}
	}

	return null;
}

seed.debug = false;

seed.globals = {};

seed.arg = function(id, arg) {
	let c = q.get(id)
	if (c && c.args) {
		return c.args[arg];
	}
	return null;
};

seed.get = (id) => {
	if (id instanceof HTMLElement) return id;

	if (seed.get.cache && id in seed.get.cache) {
		return seed.get.cache[id];
	}

	let element;
	if (id[0] == ".") {
		element = document.getElementsByClassName(id.slice(1));
		if (element) element = element[0];
	} else {
		element = document.getElementById(id);
	}
	
	if (element) {
		if (seed.get.cache) seed.get.cache[id] = element;
		return element;
	}
	
	//Check the templates
	let templates = document.querySelectorAll('template');
	for (let template of templates) {
		element = seed.get_template(template, id);
		if (element) {
			if (seed.get.cache) seed.get.cache[id] = element;
			if (!element.parentElement) {
				element.parent = template;
			}
			return element;
		}
	}

	return null;
}
seed.get.cache = {};

seed.download = async function(name, path) {
	let link = document.createElement('a');
	link.download = name;
	link.href = path;
	
	document.body.appendChild(link);
	link.click();
	document.body.removeChild(link);
}

window.AsyncFunction = Object.getPrototypeOf(async function(){}).constructor;

seed.request = async function(method, formdata, url, manual, active, onprogress) {

	const slave = async function(response) {
		if (response == "") return null;

		return await (new AsyncFunction(response))();
	}

	if (window.rpc && rpc[url]) {
		slave(rpc[url](formdata));
		return;
	}

	if (window.ServiceWorker_Registration) ServiceWorker_Registration.update();

	if (manual) {
		var xhr = new XMLHttpRequest();
		xhr.open(method, url);
		return xhr;
	}

	let promise = new Promise(function (resolve, reject) {
		var xhr = new XMLHttpRequest();

		if (onprogress) {
			xhr.upload.onprogress = function(event) {
				if (event.lengthComputable) {
					onprogress(event.loaded/event.total);
				} else {
					onprogress(-1);
				}
			};
		}

		xhr.onload = function () {
			if (this.status >= 200 && this.status < 300) {
				resolve(xhr.response);
			} else {
				if (onprogress) onprogress(0);

				if (this.status != 404) slave(xhr.response).then(function() {
					reject(reject(seed.httpErrString(this.status)));
				}).catch(function(e) {
					reject(e);
				});
				reject(seed.httpErrString(this.status));
			}
		};
		xhr.onerror = function () {
			if (onprogress) onprogress(0);
			
			reject(seed.httpErrString(this.status));
		};

		xhr.open(method, url, true);
		xhr.send(formdata);
	});

	let response = await promise;
	return await slave(response);
}

seed.httpErrString = function(status) {
	switch (status) {
	case 413:
		return "Payload Too Large";
	default:
		return "Connection Failed";
	}
}

seed.dynamic = {};

	`)

	for i := len(rootRenderers) - 1; i >= 0; i-- {
		b.Write(rootRenderers[i](root))
	}

	for i := len(renderers) - 1; i >= 0; i-- {
		b.Write(renderers[i](root))
	}

	b.Write(render(root))

	return b.Bytes()
}

//Adopt returns and removes the script from the given seed.
func Adopt(c seed.Seed) Script {
	var s = js.Script(func(q js.Ctx) {})

	s = s.Append(adopt(c).GetScript())

	return s
}

func adopt(child seed.Seed) Script {
	var s = js.Script(func(q js.Ctx) {})
	var d Data
	child.Load(&d)

	//Deterministic render.
	keys := make([]string, 0, len(d.On))
	for i := range d.On {
		keys = append(keys, string(i))
	}
	sort.Strings(keys)

	for _, event := range keys {
		handler := d.On[event]
		var e = event
		var h = handler
		s = s.Append(func(q js.Ctx) {
			fmt.Fprintf(q, `seed.on(%v, "%v", async function() {`, Element(child), e)
			h.GetScript()(q)
			fmt.Fprint(q, `});`)
		})
		if event != "ready" {
			delete(d.On, event)
		}
	}

	for _, child := range child.Children() {
		s = s.Append(adopt(child).GetScript())
	}

	if _, ok := d.On["ready"]; ok {
		s = s.Append(func(q js.Ctx) {
			fmt.Fprintf(q, `await %[1]v.onready();`, Element(child))
		})
		delete(d.On, "ready")
	}

	return s
}
