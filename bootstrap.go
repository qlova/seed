package seed

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
)

func (runtime Runtime) launchWasm() error {
	var Build = exec.Command("go", "build", `-ldflags=-s -w`, "-o", "seed.wasm")
	Build.Env = append(os.Environ(), []string{
		"GOOS=js",
		"GOARCH=wasm",
	}...)
	Build.Stdout = os.Stdout
	Build.Stderr = os.Stderr
	Build.Dir = Dir
	err := Build.Run()
	if err != nil {
		fmt.Println(err)
		return err
	}

	go launch(":1234")

	var root, _ = exec.Command("go", "env", "GOROOT").CombinedOutput()

	var scripts = runtime.app.Scripts(runtime.app.platform)

	return http.ListenAndServe(runtime.Listen, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/Qlovaseed.png" {
			w.Header().Set("Content-Type", "image/png")
			icon, _ := fsByte(false, "/Qlovaseed.png")
			w.Write(icon)
			return
		}

		//Is this an embedded resource? Imported libraries will add these.
		if embedded(w, r) {
			return
		}

		if r.URL.String() == "/wasm_exec.js" {
			http.ServeFile(w, r, string(root[:len(root)-1])+"/misc/wasm/wasm_exec.js")
		}

		if r.URL.String() == "/seed.wasm" {
			http.ServeFile(w, r, Dir+"/seed.wasm")
		}

		if r.URL.String() == "/" {
			fmt.Fprintf(w, `<!doctype html><html><head>`)

			for script := range scripts {
				if path.Ext(script) == ".js" {
					fmt.Fprintf(w, `<script src="`+script+`" defer></script>`)
				}
			}

			fmt.Fprintf(w, `</head><body><script src="wasm_exec.js"></script><script>if (!WebAssembly.instantiateStreaming){instantiateStreaming=async (resp, importObject)=>{const source=await (await resp).arrayBuffer();return await WebAssembly.instantiate(source, importObject);};}const go=new Go();let mod, inst;WebAssembly.instantiateStreaming(fetch("seed.wasm"), go.importObject).then((result)=>{mod=result.module;inst=result.instance;run();}).catch((err)=>{console.error(err);});async function run(){await go.run(inst);inst=await WebAssembly.instantiate(mod, go.importObject);}</script></body></html>`)
		}
	}))
}
