// +build !generate,wasm

package wasm

import "syscall/js"

func Generate() {
	js.Global().Get("GoPromise").Invoke()
	<-make(chan bool)
}
