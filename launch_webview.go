// +build webview

package seed

import "github.com/zserge/webview"

func launchNative(url string) error {
	return webview.Open("App", url, 800, 600, true)
}
