// +build !wasm

package wasm

import (
	"reflect"
	"runtime"

	"qlova.org/seed/use/js"
)

const InstantiateStreaming = `
if ("Go" in window) {
	if (!WebAssembly.instantiateStreaming) { // polyfill
		WebAssembly.instantiateStreaming = async (resp, importObject) => {
			const source = await (await resp).arrayBuffer();
			return await WebAssembly.instantiate(source, importObject);
		};
	}

	const go = new Go();
	let mod, inst;

	let success, failure;
	let finshed = new Promise(function(yes, no) {
		success = yes;
		failure = no;
	});

	window.GoPromise = success;

	await WebAssembly.instantiateStreaming(fetch("assets/wasm/index.wasm"), go.importObject).then((result) => {
		mod = result.module;
		inst = result.instance;
		go.run(inst);
	}).catch((err) => {
		console.error(err);
		failure();
	});

	await finshed;
}`

var exports = make(map[string]struct{})

//Exported returns true if the given function was exported with this package.
func Exported(f interface{}) bool {
	_, ok := exports[runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()]
	return ok
}

//Export exports the given function so that it can be ran with Run.
func Export(f interface{}) {
	exports[runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()] = struct{}{}
}

//Run returns a client.Script that runs the given function with the given arguments.
func Run(f interface{}, args ...js.AnyValue) js.Script {
	return js.Script(func(q js.Ctx) {
		q(js.Require("assets/js/wasm_exec.js", ""))
		q(js.Func("window['" + runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name() + "']").Run(args...))
	})
}

//Call returns a client.Script that calls the given function with the given arguments.
func Call(f interface{}, args ...js.AnyValue) js.AnyValue {
	return js.NewFunction(js.Script(func(q js.Ctx) {
		q(js.Require("assets/js/wasm_exec.js", ""))
		q(js.Return(js.Func("window['" + runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name() + "']").Call(args...)))
	})).Call()
}

//Download downloads the response of the given function.
func Download(f interface{}, args ...js.AnyValue) js.Script {
	return js.Script(func(q js.Ctx) {
		q(js.Require("assets/js/wasm_exec.js", ""))
		q(js.Func("window['" + runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name() + ".download']").Run(args...))
	})
}
