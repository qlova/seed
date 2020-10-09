//Package dropzone provides functionality for file dropzones.
package dropzone

import (
	"fmt"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/html"
	"qlova.org/seed/js"
)

//OnDrop executes the given script when something is dropped on this dropzone.
func OnDrop(do ...client.Script) seed.Option {
	return client.On("drop", do...)
}

//WhenDraggingSet updates the given boolean to true when an item is being dragged over this, and false otherwise.
func WhenDraggingSet(dragging *clientside.Bool) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var el = html.Element(c)

		js.Require("/dragster.js", dragster).AddTo(c)

		OnDrop(dragging.Set(false)).AddTo(c)

		client.OnLoad(js.Script(func(q js.Ctx) {
			fmt.Fprintf(q, "new Dragster(%v);", el)

			fmt.Fprintf(q, `
				document.addEventListener( "dragster:enter", async function (e) {
					if (e.target == %v) {`, el)

			q(dragging.Set(true))

			fmt.Fprintf(q, `}
				}, false );
				
				document.addEventListener( "dragster:leave", async function (e) {
					if (e.target == %v) {`, el)

			q(dragging.Set(false))

			fmt.Fprintf(q, `}
				}, false );
			`)
		})).AddTo(c)
	})
}
