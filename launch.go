package seed

import (
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
}

func launch(hostport string) {
	var err error

	url := "http://localhost" + hostport

	if launchNative(url) == nil {
		return
	}

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
