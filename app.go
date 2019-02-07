package seed

import "github.com/qlova/seed/style"

import (
	"net/http"
	"html"
	"strings"
)
//DEPRECIATED
func (seed Seed) ID() string {
	return seed.id
}

//TODO Should be Internal.
func (seed Seed) SetClass(class string) {
	seed.class = class
}

//TODO Should be Internal.
func (seed Seed) SetTag(tag string) {
	seed.tag = tag
}
//TODO Should be Internal.
func (seed Seed) SetAttributes(attr string) {
	seed.attr = attr
}
//TODO Should be Internal.
func (seed Seed) Attributes() string {
	return seed.attr
}

//TODO Should be Internal.
func (seed Seed) SetPlaceholder(placeholder string) {
	seed.attr += " placeholder='"+placeholder+"' "
}

//Add a font to the seed.
//TODO merge with style?
/*func (seed Seed) AddFont(name, file, weight string) {
	
	switch weight {
		case "black":
			weight = "900"
		case "semi-bold":
			weight = "600"
		case "regular":
			weight = "400"
		case "light":
			weight = "300"
		case "extra-light":
			weight = "200"
	}
	
	RegisterAsset(file)
	
	seed.fonts.Write([]byte(`@font-face {
	font-family: '`+name+`';
	src: url('`+file+`');
	font-weight: `+weight+`;
}
`))
}*/

//Does this need to be here?
func (seed Seed) GetStyle() *style.Style {
	return &seed.Style
}

func (seed Seed) Page() bool {
	return seed.page
}

func (seed Seed) Require(script string) {
	seed.scripts = append(seed.scripts, script)
}

//Add a child seed to this seed.
func (seed Seed) Add(child Interface) {
	seed.children = append(seed.children, child)
	child.Root().SetParent(seed)
}

//Add a handler to the seed, when this seed is launched as root, the handlers will be executed for each incomming request.
func (seed Seed) AddHandler(handler func(w http.ResponseWriter, r *http.Request)) {
	seed.handlers = append(seed.handlers, handler)
}


func (seed Seed) GetParent() Interface {
	return seed.parent
}


func (seed Seed) SetParent(parent Interface) {
	seed.parent = parent
}

func (seed Seed) GetChildren() []Interface {
	return seed.children
}

//Add text, html or whatever!
func (seed Seed) SetContent(data string) {
	seed.content = []byte(data)
}

//Set the text content of the seed.
func (seed Seed) SetText(data string) {
	data = html.EscapeString(data)
	data = strings.Replace(data, "\n", "<br>", -1)
	data = strings.Replace(data, "  ", "&nbsp;", -1)
	data = strings.Replace(data, "\t", "&emsp;", -1)
	seed.content = []byte(data)
}

//Set the text content of the seed which will be dynamic at runtime.
func (seed Seed) SetDynamicText(f func(User)) {
	seed.dynamicText = f
}

//Run a script when this seed is clicked.
func (seed Seed) OnClick(f func(Script)) {
	if seed.onclick == nil {
		seed.onclick = f
	} else {
		var old = seed.onclick
		seed.onclick = func(q Script) {
			old(q)
			f(q)
		}
	}
}

//Shorthand for seed.OnClick(func(q seed.Script){ page.Script(q).Goto() })
func (seed Seed) OnClickGoto(page Page) {
	seed.OnClick(func(q Script) {
		page.Script(q).Goto()
	})
}

//Run a script when this seed is ready/loaded/onload/init.
func (seed Seed) OnReady(f func(Script)) {
	if seed.onready == nil {
		seed.onready = f
	} else {
		var old = seed.onready
		seed.onready = func(q Script) {
			old(q)
			f(q)
		}
	}
}

//Run a script when this page is entered/ongoto.
func (seed Seed) OnPageEnter(f func(Script)) {
	seed.OnReady(func(q Script) {
		q.Javascript("{")
			q.Javascript("let old_enterpage = "+seed.Script(q).Element()+".enterpage;")
			q.Javascript(seed.Script(q).Element()+".enterpage = function() {")
			q.Javascript("if (old_enterpage) old_enterpage();")
			f(q)
			q.Javascript("};")
		q.Javascript("}")
	})
}

//Run a script when this leaving this page (onleave).
func (seed Seed) OnPageExit(f func(Script)) {
	seed.OnReady(func(q Script) {
		q.Javascript("{")
			q.Javascript("let old_exitpage = "+seed.Script(q).Element()+".exitpage;")
			q.Javascript(seed.Script(q).Element()+".exitpage = function() {")
			q.Javascript("if (old_exitpage) old_exitpage();")
			f(q)
			q.Javascript("};")
		q.Javascript("}")
	})
}

//Run a script when this seed's value is changed by the user.
func (seed Seed) OnChange(f func(Script)) {
	if seed.onchange == nil {
		seed.onchange = f
	} else {
		var old = seed.onchange
		seed.onchange = func(q Script) {
			old(q)
			f(q)
		}
	}
}
