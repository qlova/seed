package app

import (
	"fmt"
	"hash/fnv"
	"math"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var browsers = []string{
	"google-chrome",
	"chromium",
	"google-chrome-stable",
	"/usr/bin/google-chrome-stable",
	"/usr/bin/google-chrome",
	"/usr/bin/chromium",
	"/usr/bin/chromium-browser",
	"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
	"/Applications/Google Chrome Canary.app/Contents/MacOS/Google Chrome Canary",
	"/Applications/Chromium.app/Contents/MacOS/Chromium",
	"C:/Users/" + os.Getenv("USERNAME") + "/AppData/Local/Google/Chrome/Application/chrome.exe",
	"C:/Program Files (x86)/Google/Chrome/Application/chrome.exe",
	"C:/Program Files (x86)/Microsoft/Edge/Application/msedge.exe",
}

func launch(url string) {

	//fix strange windows bug.
	if runtime.GOOS == "windows" {
		url = strings.Replace(url, "[::]", "[::1]", 1)
	}

	var err error

	/*if launchNative(url) == nil {
		return
	}*/

	for _, browser := range browsers {
		err = exec.Command(browser, "--app="+url).Run()
		if err == nil {
			return
		}
	}

	switch runtime.GOOS {
	case "android":
		exec.Command("am", "start", "--user", "0", "-a", "android.intent.action.VIEW", "-d", url).Run()
	case "linux":
		exec.Command("xdg-open", url).Run()
	case "darwin":
		exec.Command("open", url).Run()
	case "windows":
		r := strings.NewReplacer("&", "^&")
		exec.Command("cmd", "/c", "start", r.Replace(url)).Run()
	}
}

//Launch launches the app.
func (a App) Launch() error {

	var port = ":0"

	//Allow port config from Env
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	if os.Getenv("GOPORT") != "" {
		port = os.Getenv("GOPORT")
	}

	var data app
	a.Load(&data)

	var browser bool

	var iport uint16

	//Determine a stable port number from the app's name.
	if port == ":0" {
		var hash = fnv.New64()
		hash.Write([]byte(data.name))

		iport = uint16(hash.Sum64()%(math.MaxUint16-1000)) + 1000

		port = fmt.Sprint(":", iport)
		browser = true
	}

	a.port = port

	handler := a.Handler()

retry:
	listener, err := net.Listen("tcp", port)
	if err != nil {
		if strings.Contains(err.Error(), "address already in use") {
			iport++
			port = fmt.Sprint(":", iport)
			goto retry
		}
		return err
	}

	if browser {
		go launch("http://" + listener.Addr().String())
	}

	fmt.Printf("\nlaunching %v version %v on http://localhost%v\n", data.name, data.worker.Version, port)

	return http.Serve(listener, handler)
}
