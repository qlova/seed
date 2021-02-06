// +build !generate

package wasm

func Generate() {
	<-make(chan bool)
}
