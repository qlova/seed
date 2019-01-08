package seed

import (
	"path"
	"path/filepath"
	"os"
	"log"
	"strings"
	"net/http"
) 
import "github.com/NYTimes/gziphandler"

/* 
	A Launcher is used to launch your seed into an application.
	
*/
type Launcher struct {
	Seed //The default seed.

	//Hostname and port where you want the application to be listening on.
	Listen string

	//Here you can pass different seeds to be used for different devices.
	Mobile Seed
	Tablet Seed
}


//Returns a http handler that serves this application.
func (launcher Launcher) Handler() http.Handler {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
		log.Fatal(err)
    }

	minified, err := mini(launcher.Seed.render(true))
	if err != nil {
		//Panic?
	}

	var html = launcher.Seed.render(false)

	var worker = ServiceWorker.Render()
	var manifest = launcher.Seed.manifest.Render()
	var dynamic = launcher.Seed.BuildDynamicHandler()

	var LocalClients = 0

	return gziphandler.GzipHandler(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		var local = strings.Contains(request.RemoteAddr, "[::1]")

		//Editmode socket.
		if request.URL.Path == "/socket" && local {
			LocalClients++
			SingleLocalConnection = LocalClients == 1
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
			callHandler(response, request, request.URL.Path[6:] )
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

		//Anything else? Serve application.
		if local {
			response.Write(html)
		} else {
			response.Write(minified)
		}
	}))
}

func (launcher Launcher) Launch() {
	if launcher.Seed.seed != nil {
		if launcher.Listen == "" {
			launcher.Listen = ":1234"
		}
		http.Handle("/", launcher.Handler())

		//Allow port config from Env
		if port := os.Getenv("PORT"); port != "" {
			launcher.Listen = port
		}
			
		//Launch the app if possible.
		go launch(launcher.Listen)
	
		http.ListenAndServe(launcher.Listen, nil)
		return
	}
	panic("No seeds!")
}