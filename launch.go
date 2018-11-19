package seed

import "os/exec"

var browsers = []string {
	"google-chrome",
	"chromium",
	"google-chrome-stable",	
}

func launch(hostport string) {
	var err error
	
	for _, browser := range browsers {
		err = exec.Command(browser, "--app=http://localhost"+hostport).Run()
		if err == nil {
			return
		}
	}
}
