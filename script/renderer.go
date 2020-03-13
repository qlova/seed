package script

import (
	"bytes"
	"fmt"

	"github.com/qlova/script"
	"github.com/qlova/script/language"

	"github.com/qlova/seed"
)

type Renderer func(root seed.Any) []byte

var renderers []Renderer

func RegisterRenderer(r Renderer) {
	renderers = append(renderers, r)
}

func render(child seed.Any) []byte {
	var b bytes.Buffer
	var data = seeds[child.Root()]

	for event, handler := range data.on {
		b.Write(language.Javascript(func(q script.Ctx) {
			if event == "press" {
				fmt.Fprintf(q, `seed.op(%v, async function() {`, child.Root().Ctx(q).Element())
			} else {
				fmt.Fprintf(q, `%v.on%v = async function() {`, child.Root().Ctx(q).Element(), event)
			}
			handler(Ctx{q})
			if event == "press" {
				fmt.Fprint(q, `});`)
			} else {
				fmt.Fprint(q, `};`)
			}
		}))
	}

	for _, child := range child.Root().Children() {
		b.Write(render(child))
	}

	if _, ok := data.on["ready"]; ok {
		b.Write(language.Javascript(func(q script.Ctx) {
			fmt.Fprintf(q, `%[1]v.onready();`, child.Root().Ctx(q).Element())
		}))
	}

	return b.Bytes()
}

//Render renders the Javascript attached to this seed and its children.
func Render(root seed.Any) []byte {
	var b bytes.Buffer

	b.WriteString(`seed = {};
seed.production = (location.hostname != "localhost" && location.hostname != "127.0.0.1");

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

seed.get = function(id) {
	if (id in seed.get.cache) {
		return seed.get.cache[id];
	}

	let element = document.getElementById(id);
	
	if (element) {
		seed.get.cache[id] = element;
		return element;
	}
	
	//Check the templates
	let templates = document.querySelectorAll('template');
	for (let template of templates) {
		element = template.content.getElementById(id);
		if (element) {
			seed.get.cache[id] = element;
			if (!element.parentElement) {
				element.parent = template;
			}
			return element;
		}
	}

	return null;
}
seed.get.cache = {};

seed.request = function(method, formdata, url, manual) {

	const slave = function(response) {
		const AsyncFunction = Object.getPrototypeOf(async function(){}).constructor;
		return (new AsyncFunction(response))();
	}

	if (window.rpc && rpc[url]) {
		slave(rpc[url](formdata));
		return;
	}

	if (window.ServiceWorker_Registration) ServiceWorker_Registration.update();

	//if (url.charAt(0) == "/") url = host+url;

	if (manual) {
			var xhr = new XMLHttpRequest();
			xhr.open(method, url);
		return xhr;
	}

	return new Promise(function (resolve, reject) {
		var xhr = new XMLHttpRequest();
		xhr.open(method, url, true);
		xhr.onload = function () {
			if (this.status >= 200 && this.status < 300) {
				resolve(slave(xhr.response));
			} else {
				if (this.status != 404) slave(xhr.response);
				reject({
					status: this.status,
					statusText: xhr.statusText,
					response: xhr.response
				});
			}
		};
		xhr.onerror = function () {
			reject({
				status: this.status,
				statusText: xhr.statusText,
				response: xhr.response
			});
		};
		xhr.send(formdata);
	});
}

seed.dynamic = {};
	`)

	for _, renderer := range renderers {
		b.Write(renderer(root))
	}

	b.Write(render(root))

	return b.Bytes()
}

//Adopt returns and removes the script from the given seed.
func Adopt(c seed.Any) Script {
	var s = Script(func(q Ctx) {})

	s = s.Then(adopt(c))

	return s
}

func adopt(child seed.Any) Script {
	var s = Script(func(q Ctx) {})
	var data = seeds[child.Root()]

	for event, handler := range data.on {
		s = s.Then(func(q Ctx) {
			if event == "press" {
				fmt.Fprintf(q, `seed.op(%v, async function() {`, child.Root().Ctx(q).Element())
			} else {
				fmt.Fprintf(q, `%v.on%v = async function() {`, child.Root().Ctx(q).Element(), event)
			}
			handler(q)
			if event == "press" {
				fmt.Fprint(q, `});`)
			} else {
				fmt.Fprint(q, `};`)
			}
		})
		if event != "ready" {
			delete(data.on, event)
		}
	}

	for _, child := range child.Root().Children() {
		s = s.Then(adopt(child))
	}

	if _, ok := data.on["ready"]; ok {
		s = s.Then(func(q Ctx) {
			fmt.Fprintf(q, `%[1]v.onready();`, child.Root().Ctx(q).Element())
		})
		delete(data.on, "ready")
	}

	return s
}
