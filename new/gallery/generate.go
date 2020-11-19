// +build generate
//go:generate go run -tags generate generate.go

package main

import (
	"os"

	"qlova.org/seed/assets/inbed"
	"qlova.org/seed/use/js"
)

func main() {
	inbed.Root, _ = os.Getwd()
	inbed.SingleFile = "inbed.go"
	inbed.PackageName = "gallery"

	js.Bundle("assets/js/photoswipe.js", "https://raw.githubusercontent.com/dimsemenov/PhotoSwipe/master/dist/photoswipe.min.js")
	js.Bundle("assets/js/photoswipe-ui.js", "https://raw.githubusercontent.com/dimsemenov/PhotoSwipe/master/dist/photoswipe-ui-default.min.js")
	js.Bundle("assets/css/photoswipe.css", "https://raw.githubusercontent.com/dimsemenov/PhotoSwipe/master/dist/photoswipe.css")
	js.Bundle("assets/photoswipe/default-skin.css", "https://raw.githubusercontent.com/dimsemenov/PhotoSwipe/master/dist/default-skin/default-skin.css")
	js.Bundle("assets/photoswipe/default-skin.svg", "https://raw.githubusercontent.com/dimsemenov/PhotoSwipe/master/dist/default-skin/default-skin.svg")
	js.Bundle("assets/photoswipe/default-skin.png", "https://raw.githubusercontent.com/dimsemenov/PhotoSwipe/master/dist/default-skin/default-skin.png")
	js.Bundle("assets/photoswipe/preloader.gif", "https://raw.githubusercontent.com/dimsemenov/PhotoSwipe/master/dist/default-skin/preloader.gif")

	inbed.Done()
}
