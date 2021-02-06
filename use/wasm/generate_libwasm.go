// +build generate

package wasm

import (
	"fmt"
	"go/build"
	"os"
	"os/exec"

	"qlova.org/seed/assets/inbed"
	"qlova.org/seed/use/js"
)

var TinyGo bool = false

func Generate() {
	inbed.Root, _ = os.Getwd()
	inbed.SingleFile = "inbed.go"
	inbed.PackageName = "main"

	var cmd *exec.Cmd

	if TinyGo {
		cmd = exec.Command("tinygo", "build", "-target", "wasm", "-o", "index.wasm")
	} else {
		cmd = exec.Command("go", "build", "-o", "index.wasm")
		cmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")
	}

	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	if TinyGo {
		root := "/usr/local/lib/tinygo"

		js.Bundle("assets/js/wasm_exec.js", root+"/targets/wasm_exec.js")

	} else {

		root := build.Default.GOROOT

		js.Bundle("assets/js/wasm_exec.js", root+"/misc/wasm/wasm_exec.js")
	}
	js.Bundle("assets/wasm/index.wasm", "index.wasm")

	inbed.Done()
}
