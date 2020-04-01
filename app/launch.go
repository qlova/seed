package app

import (
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
func (app App) Launch() error {

	handler := app.Handler()

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return err
	}

	go launch("http://" + listener.Addr().String())

	return http.Serve(listener, handler)
}
