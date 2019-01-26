package seed

import (
	"path"
	"path/filepath"
	"os"
	"log"
	"strings"
	"net/http"
) 

import ua "github.com/avct/uasurfer"
import "github.com/NYTimes/gziphandler"

import "github.com/qlova/seed/script"


type launcher struct {
	App

	//Hostname and port where you want the application to be listening on.
	Listen string
}


//Returns a http handler that serves this application.
func (launcher launcher) Handler() http.Handler {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
		log.Fatal(err)
    }

	minified, err := mini(launcher.render(true, Default))
	if err != nil {
		//Panic?
	}

	var html = launcher.render(false, Default)

	var worker = ServiceWorker.Render()
	var manifest = launcher.Manifest.Render()
	var dynamic = launcher.Seed.BuildDynamicHandler()

	var desktop = launcher.render(true, Desktop)

	var LocalClients = 0

	return gziphandler.GzipHandler(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		var local = strings.Contains(request.RemoteAddr, "[::1]")

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
			script.Handler(response, request, request.URL.Path[6:] )
			return
		}
		
		//Remote procedure calls.
		if len(request.URL.Path) > 6 && request.URL.Path[:7] == "/feeds/" {
			feedHandler(response, request, request.URL.Path[7:])
			return
		}

		//Run custom handlers.
		if request.URL.Path != "/" {
			for _, handler := range launcher.Seed.handlers {
				handler(response, request)
			}
			
			if path.Ext(request.URL.Path) == "" {
				return
			}
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
		if os.Getenv("IGNORE_PLATFORM") == "" {
			device := ua.Parse(request.UserAgent())
	
			if device.DeviceType == ua.DeviceComputer {
				response.Write(desktop)
				return
			}
		}

		//Anything else? Serve application.
		if local {
			response.Write(html)
		} else {
			response.Write(minified)
		}
	}))
}

func (launcher launcher) Launch(port ...string) {
	if launcher.Seed.seed != nil {
		http.Handle("/", launcher.Handler())

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
			
		//Launch the app if possible.
		go launch(launcher.Listen)
	
		http.ListenAndServe(launcher.Listen, nil)
		return
	}
	panic("No seeds!")
}
