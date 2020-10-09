//Package clientzip provides clientside zipping functionality.
package clientzip

import (
	"fmt"
	"strconv"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/js"
	"qlova.org/seed/s/dropzone"
)

//Policy determines what is zipped.
type Policy int

func (p Policy) String() string {
	switch p {
	case Always:
		return "always"
	case FoldersOnly:
		return "folders"
	}
	return ""
}

//Policies.
const (
	//File is always zipped.
	Always Policy = iota

	//Only Folders are zipped, if a File is set to a single file, it will be left as is.
	FoldersOnly
)

//File is a (potentially zipped) file in client memory.
type File struct {
	clientside.File

	ZipPolicy Policy

	//WhenZippingSet sets the file to set the given boolean when it is being zipped.
	WhenZippingSet *clientside.Bool
}

//Choose asks the client to choose a File.
func (f *File) Choose() client.Script {
	return js.Script(func(q js.Ctx) {
		q(`"'#(import "/jszip.js")`)
		var input = js.Global().Get("document").Call("createElement", js.NewString("input")).Var(q)
		q(input.Set("type", js.NewString("file")))
		q(input.Run("setAttribute", js.NewString("multiple"), js.NewString("")))
		q(input.Run("setAttribute", js.NewString("directory"), js.NewString("")))
		q(input.Run("setAttribute", js.NewString("webkitdirectory"), js.NewString("")))
		q(input.Run("setAttribute", js.NewString("allowdirs"), js.NewString("")))

		q("let active = seed.active;")

		q(input.Set("oninput", js.NewFunction(func(q js.Ctx) {
			q(js.Global().Get("seed").Set("active", js.NewValue(`active`)))

			if f.WhenZippingSet != nil {
				q(f.WhenZippingSet.Set(true))
			}

			fmt.Fprintf(q, `
			let file;

			var target = event.target;
			if (target.files && target.files.length) {
				file = await clientzip.WebkitRelativePaths(target, %[1]v);
			} else if ("getFilesAndDirectories" in target) {
				file = await FilesAndDirectories(target, %[1]v);
			} else if (event.file) {
				if (%[1]v == "folders") file = target.files[0];
			}
			`, strconv.Quote(f.ZipPolicy.String()))

			if f.WhenZippingSet != nil {
				q(f.WhenZippingSet.Set(false))
			}

			q(f.SetToRaw(js.Null()))
			q(f.SetToRaw(js.NewValue("file")))

		}, "event")))

		q(input.Run("click"))
	})
}

//Dropzone sets the seed to act as a dropzone for this file.
func (f *File) Dropzone() seed.Option {
	return seed.Options{
		js.Require("/jszip.js", jszip),

		client.On("dragenter", js.Script(func(q js.Ctx) {
			q(`arguments[0].preventDefault();`)
		})),
		client.On("dragover", js.Script(func(q js.Ctx) {
			q(`arguments[0].preventDefault();`)
		})),

		dropzone.OnDrop(js.Script(func(q js.Ctx) {

			if f.WhenZippingSet != nil {
				q(f.WhenZippingSet.Set(true))
			}

			fmt.Fprintf(q, `{
				let event = arguments[0];
				event.preventDefault();
				event.stopPropagation();

				let dt = event.dataTransfer;
				let file;
				let name;

				if (dt.items && dt.items.length && "webkitGetAsEntry" in dt.items[0] && dt.items[0].webkitGetAsEntry()) {
					file = await clientzip.AsEntries(dt.items, %[1]v);
				} else if ("getFilesAndDirectories" in dt) {
					file = await clientzip.FilesAndDirectories(dt, %[1]v);
				} else if (dt.files) {
					file = await clientzip.WebkitRelativePaths(dt, %[1]v);
				} else if (event.file) {
					if (%[1]v == "folders") file = event.file;
				}
			`, strconv.Quote(f.ZipPolicy.String()))

			if f.WhenZippingSet != nil {
				q(f.WhenZippingSet.Set(false))
			}

			q(f.SetToRaw(js.Null()))
			q(f.SetToRaw(js.NewValue("file")))

			q("}")
		})),
	}
}
