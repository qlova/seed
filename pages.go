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
	seed.Style.Style.SetWidth(css.Number(100).Vw())
	seed.Style.Style.SetHeight(css.Number(100).Vh())

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

func (page Page) SyncVisibilityWith(seed Interface) {
	var root = seed.Root()
	page.OnPageEnter(func(q Script) {
		root.Script(q).SetVisible()
	})
	page.OnPageExit(func(q Script) {
		root.Script(q).SetHidden()
	})
}

func (page Page) Script(q Script) script.Page {
	return script.Page{page.Seed.Script(q)}
}

//Run a script when this page is entered/ongoto.
func (seed Page) OnPageEnter(f func(Script)) {
	seed.OnReady(func(q Script) {
		q.Javascript("{")
		q.Javascript("let old_enterpage = " + seed.Script(q).Element() + ".enterpage;")
		q.Javascript(seed.Script(q).Element() + ".enterpage = function() {")
		q.Javascript("if (old_enterpage) old_enterpage();")
		f(q)
		q.Javascript("};")
		q.Javascript("}")
	})
}

//Run a script when this leaving this page (onleave).
func (seed Page) OnPageExit(f func(Script)) {
	seed.OnReady(func(q Script) {
		q.Javascript("{")
		q.Javascript("let old_exitpage = " + seed.Script(q).Element() + ".exitpage;")
		q.Javascript(seed.Script(q).Element() + ".exitpage = function() {")
		q.Javascript("if (old_exitpage) old_exitpage();")
		f(q)
		q.Javascript("};")
		q.Javascript("}")
	})
}
