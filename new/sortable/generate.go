// +build generate

package main

import (
	"os"

	"qlova.org/seed/assets/inbed"
	"qlova.org/seed/use/js"
)

func main() {
	inbed.Root, _ = os.Getwd()
	inbed.SingleFile = "inbed.go"
	inbed.PackageName = "sortable"

	js.Bundle("assets/js/sortable.js", "https://cdn.jsdelivr.net/npm/sortablejs@latest/Sortable.min.js")

	inbed.Done()
}
