// +build webview

package seed

import "github.com/zserge/webview"

func launch_native(url string) error {
	return webview.Open("App", url, 800, 600, true)
}
