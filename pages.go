package seed

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

import "github.com/qlova/seed/script"
import "github.com/qlova/seed/style/css"

type Page struct {
	Seed

	content map[string]string
}

func NewPage() Page {
	seed := New()
	seed.SetCol()

	seed.page = true
	seed.class = "page"

	seed.SetHidden()
	seed.SetWillChange(css.Property.Display)

	seed.SetPosition(css.Fixed)
	seed.SetTop(css.Zero)
	seed.SetLeft(css.Zero)
	seed.SetWidth(css.Number(100).Vw())
	seed.SetHeight(css.Number(100).Vh())

	return Page{seed, nil}
}

type pages map[string]Page

func (p pages) Get(key string) Page {
	return p[key]
}

var Pages = make(pages)

//We need to load all of the static pages.
func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
		return
	}

	files, err := ioutil.ReadDir(dir + "/content")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".page" {
			var name = strings.Replace(file.Name(), ".page", "", 1)
			Pages[name] = NewPage()
			Pages[name] = Page{
				Seed:    Pages[name].Seed,
				content: openIML(dir + "/content/" + file.Name()),
			}
		}
	}
}

func AddPageTo(parent Interface) Page {
	var page = NewPage()
	parent.Root().Add(page)
	return page
}

func (page Page) Get(key string) string {
	return page.content[key]
}

func (page Page) Script(q Script) script.Page {
	return script.Page{page.Seed.Script(q)}
}

func (seed Seed) SetPage(page Page) {
	seed.OnReady(func(q Script) {
		q.Javascript(`if (window.localStorage.getItem("update")) {`)
		q.Javascript(`window.localStorage.removeItem("update");`)
		q.Javascript(`window.localStorage.removeItem("*CurrentPage");`)
		q.Javascript(`}`)
		q.Javascript(`if (!window.localStorage.getItem("*CurrentPage")) goto("` + page.id + `");`)
	})
}
