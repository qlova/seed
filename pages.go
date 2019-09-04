package seed

import "github.com/qlova/seed/script"
import "github.com/qlova/seed/style/css"

type Page struct {
	Seed
	tags map[string]bool
}

func NewPage() Page {
	seed := New()
	seed.SetCol()

	seed.page = true
	seed.class = "page"

	seed.SetHidden()
	seed.SetWillChange(css.Property.Display)
	seed.SetWillChange(css.Property.Transform)

	seed.SetPosition(css.Absolute)
	seed.SetTop(css.Zero)
	seed.SetLeft(css.Zero)
	//seed.Style.Style.SetWidth(css.Number(100).Vw())
	//seed.Style.SetHeight(100)*/
	seed.Style.SetSize(100, 100)

	return Page{seed, nil}
}

type pages map[string]Page

func (p pages) Get(key string) Page {
	return p[key]
}

var Pages = make(pages)

func (seed Seed) NewPage() Page {
	return AddPageTo(seed)
}

func AddPageTo(parent Interface) Page {
	var page = NewPage()
	parent.Root().Add(page)
	return page
}

func (page Page) SetBack(back Page) {
	page.SetAttributes(page.Attributes() + ` data-back="` + back.ID() + `"`)
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

//OnBack is triggered before the back action is triggered, return q.Bool(true) to prevent default behaviour.
func (page Page) OnBack(f func(q Script)) {
	page.OnReady(func(q Script) {
		q.Javascript("{")
		q.Javascript("let old_onback = " + page.Script(q).Element() + ".onback;")
		q.Javascript(page.Script(q).Element() + ".onback = function() {")
		q.Javascript("if (old_onback) old_onback();")
		f(q)
		q.Javascript("};")
		q.Javascript("}")
	})
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
