// +build wasm

package dating

import (
	"fmt"
	"net/http"
	"syscall/js"
)

var Host = js.Global().Get("location").Get("host").String()
var Protocol = js.Global().Get("location").Get("protocol").String()

func init() {
	go func() {
		var res, err = http.Get(Protocol + "//" + Host + "/assets/NZ-AUK-2021.json")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		readPopular(res.Body)
	}()
}
