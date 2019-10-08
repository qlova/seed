package seed

import (
	"os"
	"path/filepath"
)

//TargetEnum is an experimental target type.
type TargetEnum int

//Purely experimental targets for Qlovaseed.
const (
	ReactNative TargetEnum = iota
	Flutter
	Website
)

//Export is a method for investigating experimental export targets.
//This should not be used naively.
func (app App) Export(t TargetEnum) error {
	if t == Website {
		var dir = filepath.Dir(os.Args[0])

		os.Mkdir(dir+"/website", 0755)

		var index, err = os.Create(dir + "/website/index.html")
		if err != nil {
			panic(err.Error())
		}
		defer index.Close()

		for name, data := range embeddings {
			var file, err = os.Create(dir + "/website/" + name)
			if err != nil {
				panic(err.Error())
			}
			defer file.Close()

			file.Write(data.Data)
		}

		//Default icon.
		{
			var file, err = os.Create(dir + "/website/Qlovaseed.png")
			if err != nil {
				panic(err.Error())
			}
			defer file.Close()

			icon, _ := fsByte(false, "/Qlovaseed.png")

			file.Write(icon)
		}

		//App Manifest.
		{
			var file, err = os.Create(dir + "/website/app.webmanifest")
			if err != nil {
				panic(err.Error())
			}
			defer file.Close()

			file.Write(app.Manifest.Render())
		}

		//Service worker.
		{
			var file, err = os.Create(dir + "/website/index.js")
			if err != nil {
				panic(err.Error())
			}
			defer file.Close()

			file.Write(app.Worker.Render())
		}

		index.Write(app.render(true, Mobile))
		return nil
	}
	panic("Invalid Export Target")
}
