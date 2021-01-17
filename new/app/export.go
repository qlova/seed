package app

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"

	"qlova.org/seed/assets/inbed"
)

//Export exports the app to the folder "export" in the current working directory.
func (a App) Export() error {

	//TODO implement basic auth for staging endpoints.
	//use crypto/subtle

	var app app
	a.Load(&app)

	inbed.File("assets")

	if err := inbed.Done(); err != nil {
		log.Println(err)
	}

	a.build()

	var rendered = app.document.Render()

	var document, err = mini(rendered)
	if err != nil {
		document = rendered
	}

	//Todo package other assets.
	//var scripts = js.Scripts(app.document.Seed)
	//var stylesheets = css.Stylesheets(app.document.Seed)
	//var imports = js.Imports()

	//Checksum is used for versioning, ensure deterministic renderers are used to prevent distributed versions from mismatching.
	//use deterministic ordered-maps instead of default maps or sort the keys before iteration.
	var checksum = md5.Sum(document)

	var version = hex.EncodeToString(checksum[:])

	app.worker.Version = version

	var worker = app.worker.Render()

	os.Mkdir("export", os.ModePerm)

	{
		icon, _ := fsByte(false, "/Qlovaseed.png")
		if err := ioutil.WriteFile("export/Qlovaseed.png", icon, os.ModePerm); err != nil {
			return err
		}
	}
	{
		var manifest = app.manifest.Render()
		if err := ioutil.WriteFile("export/app.webmanifest", manifest, os.ModePerm); err != nil {
			return err
		}
	}
	{
		if err := ioutil.WriteFile("export/robots.txt", []byte("\n"), os.ModePerm); err != nil {
			return err
		}
	}
	{
		if err := ioutil.WriteFile("export/index.js", worker, os.ModePerm); err != nil {
			return err
		}
	}
	{
		if err := ioutil.WriteFile("export/index.html", document, os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}
