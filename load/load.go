package load

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"archive/tar"
	"compress/gzip"

	"github.com/qlova/seed"

	"github.com/qlova/seed/load/pencil"
)

//File scans a Pencil file and returns it as an app.
func File(name string) *seed.App {
	var file, err = os.Open(name)
	if err != nil {
		//TODO error handling
		fmt.Println(err)
		os.Exit(1)
	}

	gzf, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tarReader := tar.NewReader(gzf)

	var project = make(map[string][]byte)

	for {
		f, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			//TODO error handling
			fmt.Println(err)
			os.Exit(1)
		}

		project[f.Name], err = ioutil.ReadAll(tarReader)

		if err != nil {
			//TODO error handling
			fmt.Println(err)
			os.Exit(1)
		}
	}

	var App = seed.NewApp()

	var document pencil.Document
	var decoder = xml.NewDecoder(bytes.NewReader(project["content.xml"]))
	err = decoder.Decode(&document)
	if err != nil {
		fmt.Println(err.Error())
	}

	//For each page in the document.
	for _, page := range document.Pages.Page {
		var SeedPage = App.NewPage()

		var Container = seed.AddTo(SeedPage)
		Container.SetSize(100, 100)
		Container.SetScrollable()

		App.SetPage(SeedPage)

		var p pencil.Page
		var decoder = xml.NewDecoder(bytes.NewReader(project[page.Href]))
		err = decoder.Decode(&p)
		if err != nil {
			fmt.Println(err.Error())
		}

		//fmt.Println(string(project[page.Href]))

		var SVG = seed.AddTo(Container)
		SVG.SetTag("svg")
		SVG.Set("viewBox", fmt.Sprint("0 0 ", p.Width(), " ", p.Height()))
		SVG.Set("preserveAspectRatio", "meet")
		SVG.Set("preserveAspectRatio", "xMidYMin meet")
		SVG.SetContent(string(p.Content.Data))
	}

	return App
}
