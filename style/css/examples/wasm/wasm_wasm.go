// +build js,wasm

//Makes the background of the page yellow.
package main

import (
	"fmt"
	"syscall/js"

	"github.com/qlova/seed/style/css"
)

func main() {
	var body = js.Global().Get("document").Get("body")
	var style = css.StyleOf(body)

	style.SetMargin(css.Zero)
	style.SetWidth(css.Number(100).Vw())
	style.SetHeight(css.Number(100).Vh())
	style.SetBackgroundColor(css.LightGoldenRodYellow)

	fmt.Println(style.BackgroundColor())
}
