//+build !wasm

package seed

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"time"

	"github.com/radovskyb/watcher"
)

//import ua "github.com/avct/uasurfer"

//Live signals if the application is live or not.
var Live bool

var exporting bool

func init() {

	for _, arg := range os.Args {
		if arg == "-live" {
			Live = true
		}
	}
	for _, arg := range os.Args {
		if arg == "-export=static" {
			exporting = true
		}
	}

	if Production {
		Live = true
	}
}

var defers []func()

func cleanup() {
	for _, f := range defers {
		f()
	}
}
func deferFunction(f func()) {
	defers = append(defers, f)
}

func init() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()
}

//Launch launches the app listening on the given port.
func (runtime Runtime) Launch(port ...string) {
	if runtime.app.Seed.seed != nil {

		if len(port) > 0 {
			runtime.Listen = port[0]
		}

		//Allow port config from Env
		if port := os.Getenv("PORT"); port != "" {
			runtime.Listen = port
		}

		if runtime.Listen == "" {
			runtime.Listen = ":1234"
		}

		if !Live && !Production {

			if runtime.bootstrapWasm {
				runtime.launchWasm()
				return
			}

			//Launch the app if possible.
			go launch(":10000")

			var Process = exec.Command(os.Args[0], "-live")
			Process.Stdout = os.Stdout
			Process.Stderr = os.Stderr
			Process.Start()

			patrol := watcher.New()
			defer patrol.Close()

			patrol.SetMaxEvents(1)

			patrol.AddFilterHook(func(info os.FileInfo, fullPath string) error {
				if path.Ext(fullPath) == ".go" {
					return nil
				}
				return watcher.ErrSkip
			})

			var Compiler *exec.Cmd

			deferFunction(func() {
				if Process.Process != nil {
					Process.Process.Kill()
				}
				if Compiler != nil && Compiler.Process != nil {
					Compiler.Process.Kill()
				}
			})

			var Compiling bool

			err := patrol.AddRecursive(".")
			if err != nil {
				log.Fatal(err)
			}

			go patrol.Start(time.Millisecond * 100)

			go func() {
				for {
					select {
					case event := <-patrol.Event:
						fmt.Println(event) // Print the event's info.

						if Compiling {
							continue
						}

						Compiler = exec.Command("go", "build", "-i", "-o", os.Args[0])
						Compiling = true
						go func() {
							err := Compiler.Run()
							if err == nil {
								if Process.Process != nil {
									Process.Process.Kill()
								}
								Process = exec.Command(os.Args[0], "-live")
								Process.Stdout = os.Stdout
								Process.Stderr = os.Stderr
								Process.Start()

								reloading = true
								for _, socket := range localSockets {
									socket.WriteMessage(1, []byte("update();"))
								}
							} else {
								println(err.Error())
							}
							Compiling = false

						}()

					case err := <-patrol.Error:
						log.Fatalln(err)
					case <-patrol.Closed:
						return
					}
				}
			}()

			proxy(runtime.Listen, ":10000")

		} else {
			http.Handle("/", runtime.app.Handler())
			http.ListenAndServe(runtime.Listen, nil)
		}
		return
	}
	panic("No seeds!")
}
