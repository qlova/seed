// +build !webview

package seed

import (
	"errors"
)

func launchNative(url string) error {
	return errors.New("no native runtime")
}
