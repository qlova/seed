package app

import (
	"fmt"
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

	a.port = port

	handler := a.Handler()

	listener, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	if port == ":0" {
		splits := strings.Split(listener.Addr().String(), ":")
		port = ":" + splits[len(splits)-1]
		go launch("http://" + listener.Addr().String())
	}

	var data app
	a.Load(&data)

	fmt.Printf("\nlaunching %v version %v on http://localhost%v\n", data.name, data.worker.Version, port)

	return http.Serve(listener, handler)
}
