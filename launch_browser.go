// +build !webview

package seed

import (
	"errors"
)

func launch_native(url string) error {
	return errors.New("no native runtime")
}
