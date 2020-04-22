package wasm

import (
	"fmt"
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/qlova/seed"
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
)

type Package string

func (Package) AddTo(c seed.Seed) {
	wasmexec, err := ioutil.ReadFile(build.Default.GOROOT + "/misc/wasm/wasm_exec.js")
	if err != nil {
		log.Fatalln(err)
	}

	c.Add(js.Require("/wasm_exec.js", string(wasmexec)))
}

func (p Package) And(more ...seed.Option) seed.Option {
	return seed.And(p, more...)
}

//Exec executes a go package as a script.
func (p Package) Exec() script.Script {
	var path = string(p)

	var ctx = build.Default
	ctx.GOARCH = "wasm"
	ctx.GOOS = "js"
	ctx.Dir = seed.Dir

	var filename = strings.Replace(path, "/", "_", -1) + ".wasm"

	var f = func(q script.Ctx) {

		q(fmt.Sprintf(`
		if (!WebAssembly.instantiateStreaming) { // polyfill
			WebAssembly.instantiateStreaming = async (resp, importObject) => {
				const source = await (await resp).arrayBuffer();
				return await WebAssembly.instantiate(source, importObject);
			};
		}

		const go = new Go();
		let mod, inst;
		let result = await WebAssembly.instantiateStreaming(fetch(%v), go.importObject);

		mod = result.module;
		inst = result.instance;
		await go.run(inst);		
		`, strconv.Quote("/assets/"+filename)))
	}

	abs, err := filepath.Abs(ctx.Dir)
	if err != nil {
		return f
	}

	pkg, err := ctx.Import(path, "", build.FindOnly)
	if err != nil {
		return f
	}

	cmd := exec.Command("go", "build", "-o", abs+"/assets/"+filename)
	cmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")
	cmd.Dir = pkg.Dir
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}

	return f
}
