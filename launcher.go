package seed

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"path/filepath"
	"strings"
	"regexp"
)

//import ua "github.com/avct/uasurfer"
import "github.com/NYTimes/gziphandler"

import "github.com/qlova/seed/script"
import "github.com/qlova/seed/user"

import "github.com/fsnotify/fsnotify"

type launcher struct {
	App

	//Hostname and port where you want the application to be listening on.
	Listen string
}

//Returns a http handler that serves this application.
func (launcher launcher) Handler() http.Handler {
	launcher.App.build()

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	minified, err := mini(launcher.render(true, Default))
	if err != nil {
		//Panic?
	}

	var html = launcher.render(false, Default)

	var worker = launcher.App.Worker.Render()
	var manifest = launcher.Manifest.Render()

	var dynamic = launcher.App.DynamicHandler()
	var custom = launcher.App.CustomHandler()

	//var desktop = launcher.render(true, Desktop)

	var LocalClients = 0

	intranet, err := regexp.Compile(`(^192\.168\.([0-9]|[0-9][0-9]|[0-2][0-5][0-5])\.([0-9]|[0-9][0-9]|[0-2][0-5][0-5]):.*$)`)
	if err != nil {
		panic("invalid regexp!")
	}
	
	return gziphandler.GzipHandler(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if origin := request.Header.Get("Origin"); origin == "https://"+launcher.App.host && origin != "" {
			response.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			response.Header().Set("Access-Control-Allow-Origin", "file://")
		}
		response.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		response.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if request.Method == "OPTIONS" {
			return
		}

		var local = strings.Contains(request.RemoteAddr, "[::1]") || strings.Contains(request.RemoteAddr, "127.0.0.1")
		if intranet.Match([]byte(request.RemoteAddr)) {
			local = true
		}

		//Editmode socket.
		if request.URL.Path == "/socket" && local {
			LocalClients++
			singleLocalConnection = LocalClients == 1
			socket(response, request)
			return
		}

		//Dynamic data.
		if request.URL.Path == "/dynamic" && dynamic != nil {
			response.Write([]byte("{"))
			dynamic(response, request)
			response.Write([]byte("}"))
			return
		}

		//Is this an embedded resource? Imported libraries will add these.
		if embedded(response, request) {
			return
		}

		//Remote procedure calls.
		if len(request.URL.Path) > 5 && request.URL.Path[:6] == "/call/" {
			script.Handler(response, request, request.URL.Path[6:])
			return
		}

		//Remote procedure calls.
		if len(request.URL.Path) > 6 && request.URL.Path[:7] == "/feeds/" {
			feedHandler(response, request, request.URL.Path[7:])
			return
		}

		//Serve assets.
		if len(request.URL.Path) > len("/attachments/") && request.URL.Path[:len("/attachments/")] == "/attachments/" {
			http.ServeFile(response, request, dir+"/"+user.AttachmentDirectory+"/"+request.URL.Path[len("/attachments/"):])
			return
		}

		//Run custom handlers.
		if request.URL.Path != "/" {
			if custom != nil {
				custom(response, request)
			}

			if path.Ext(request.URL.Path) == "" {
				return
			}
		}

		if launcher.App.rest != "" && (request.URL.Host == launcher.App.rest || request.Host == launcher.App.rest) {
			response.Write([]byte(string("This place is for computers")))
			return

		}

		if request.URL.Path == "/.well-known/assetlinks.json" && launcher.App.pkg != "" {
			response.Header().Set("Content-Type", "application/json")
			response.Write([]byte(`[{
  "relation": ["delegate_permission/common.handle_all_urls"],
  "target" : { "namespace": "android_app", "package_name": "` + launcher.App.pkg + `",
               "sha256_cert_fingerprints": [`))

			for i, hash := range launcher.App.hashes {
				response.Write([]byte("\"" + hash + "\""))
				if i < len(launcher.App.hashes)-1 {
					response.Write([]byte(`,`))
				}
			}

			response.Write([]byte(`] }
}]`))
			return
		}

		//Serve service worker.
		if request.URL.Path == "/index.js" {
			if local {
				//Don't use a web worker if we are running locally.
				response.Header().Set("content-type", "text/javascript")
				response.Write([]byte(`self.addEventListener('install', () => {self.skipWaiting();});`))
			} else {
				response.Header().Set("content-type", "text/javascript")
				response.Write(worker)
			}
			return
		}

		//Serve web manifest.
		if request.URL.Path == "/app.webmanifest" {
			response.Header().Set("content-type", "application/json")
			response.Write(manifest)
			return
		}

		//Serve assets.
		if path.Ext(request.URL.Path) != "" {
			http.ServeFile(response, request, dir+"/assets"+request.URL.Path)
			return
		}

		//Identify platform.
		/*if os.Getenv("IGNORE_PLATFORM") == "" {
			device := ua.Parse(request.UserAgent())

			if device.DeviceType == ua.DeviceComputer {
				response.Write(desktop)
				return
			}
		}*/

		//Anything else? Serve application.
		if local {
			response.Write(html)
		} else {
			response.Write(minified)
		}
	}))
}

//This is a flag, that signals if the application is live or not.
var Live bool

func init() {
	for _, arg := range os.Args {
		if arg == "-live" {
			Live = true
		}
	}
}

var defers []func()

func Cleanup() {
	for _, f := range defers {
		f()
	}
}
func Defer(f func()) {
	defers = append(defers, f)
}

func init() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		<-c
		Cleanup()
		os.Exit(1)
	}()
}

func (launcher launcher) Launch(port ...string) {
	if launcher.Seed.seed != nil {

		if len(port) > 0 {
			launcher.Listen = port[0]
		}

		//Allow port config from Env
		if port := os.Getenv("PORT"); port != "" {
			launcher.Listen = port
		}

		if launcher.Listen == "" {
			launcher.Listen = ":1234"
		}

		if !Live {

			//Launch the app if possible.
			go launch(":10000")

			var Process = exec.Command(os.Args[0], "-live")
			Process.Stdout = os.Stdout
			Process.Start()

			watcher, err := fsnotify.NewWatcher()
			if err != nil {
				log.Fatal(err)
			}
			defer watcher.Close()

			var Compiler *exec.Cmd

			Defer(func() {
				if Process.Process != nil {
					Process.Process.Kill()
				}
				if Compiler != nil && Compiler.Process != nil {
					Compiler.Process.Kill()
				}
			})

			var Compiling bool

			go func() {
				for {
					select {
					case event, ok := <-watcher.Events:
						if !ok {
							return
						}
						//log.Println("event:", event)
						if event.Op&fsnotify.Write == fsnotify.Write {

							if path.Ext(event.Name) == ".go" {

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
										Process.Start()

										RELOADING = true
										for _, socket := range LocalSockets {
											socket.WriteMessage(1, []byte("window.location.reload();"))
										}
									} else {
										println(err.Error())
									}
									Compiling = false

								}()

							}

						}
					case err, ok := <-watcher.Errors:
						if !ok {
							return
						}
						log.Println("error:", err)
					}
				}
			}()

			err = watcher.Add(path.Dir(os.Args[0]))
			if err != nil {
				log.Fatal(err)
			}

			proxy(launcher.Listen, ":10000")

		} else {
			http.Handle("/", launcher.Handler())
			http.ListenAndServe(launcher.Listen, nil)
		}
		return
	}
	panic("No seeds!")
}
