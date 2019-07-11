package seed

import "github.com/qlova/seed/style/css"
import "github.com/qlova/seed/style"
import "github.com/qlova/seed/user"

import (
	"encoding/base64"
	"html"
	"math/big"
	"net/http"
	"strings"

	"os"
	"path/filepath"
)

var Dir = filepath.Dir(os.Args[0])

type User = user.User

//Return a reference to a new type of user data, this is small data that is used to identify the user.
func UserData() user.Data {
	return user.DataType()
}

//The Vm value is relative to the viewport size.
const Vm = style.Vm

//The Em value is relative to the current font size.
const Em = style.Em

//The Px value (pixels)
const Px = style.Px

//The auto value where appropriate will automatically select a suitable value.
const Auto = style.Auto

const Top = style.Top
const Bottom = style.Bottom
const Left = style.Left
const Right = style.Right
const Center = style.Center

var Arial = style.Font{
	FontFace: css.FontFace{
		FontFamily: "Arial",
	},
}

type Font struct {
	style.Font
	path string
}

var FontCache = make(map[string]Font)

func NewFont(path string) Font {
	if font, ok := FontCache[path]; ok {
		return font
	}

	var font = Font{style.NewFont(path), path}
	FontCache[path] = font
	return font
}

func (seed Seed) SetFont(font Font) {
	seed.font = font.Font
	NewAsset(font.path).AddTo(seed)
	seed.Style.SetFont(font.Font)
}

type Seed struct {
	*seed
}

type seed struct {
	style.Style

	id               string
	tag, attr, class string
	children         []Interface

	styled bool
	ready  bool

	font      style.Font
	animation Animation

	content []byte
	page    bool

	onclick  func(Script)
	onchange func(Script)
	onready  func(Script)

	template bool

	parent Interface

	//This is a list of scripts that are needed by this seed.
	//eg. []string{"jquery.js"}
	scripts []string

	handlers []func(w http.ResponseWriter, r *http.Request)

	dynamicText func(User)

	Landscape, Portrait style.Style

	screenSmallerThan, screenGreaterThan map[Unit]style.Style

	desktop, mobile, tablet, watch, tv, native Seed

	app *App

	assets []Asset

	states map[State]func(Script)
}

type Unit = complex128

func (seed Seed) ScreenSmallerThan(unit Unit) style.Style {
	if seed.screenSmallerThan == nil {
		seed.screenSmallerThan = make(map[Unit]style.Style)
	}

	if s, ok := seed.screenSmallerThan[unit]; ok {
		return s
	}
	var s = style.New()
	seed.screenSmallerThan[unit] = s
	return s
}

func (seed Seed) Null() bool {
	return seed.seed == nil
}

func (seed Seed) MarshalText() ([]byte, error) {
	return []byte("#" + seed.ID()), nil
}

func (seed Seed) ID() string {
	if seed.seed == nil {
		return ""
	}
	return seed.id
}

func (seed Seed) Children() []Interface {
	return seed.children
}

func (seed Seed) SetClass(class string) {
	seed.class = class
}

func (seed Seed) Tag() string {
	return seed.tag
}

func (seed Seed) SetTag(tag string) {
	seed.tag = tag
}

func (seed Seed) SetAttributes(attr string) {
	seed.attr = attr
}

func (seed Seed) Attributes() string {
	return seed.attr
}

func (seed Seed) AddTo(parent Interface) Seed {
	parent.Root().Add(seed)
	return seed
}

func (seed Seed) clone() Seed {
	var clone = New()
	clone.id = seed.id
	clone.tag = seed.tag
	clone.attr = seed.attr
	clone.class = seed.class
	clone.content = seed.content
	clone.parent = seed.parent

	return clone
}

//Return the seed that should replace this seed when on the Desktop.
func (seed Seed) Desktop() Seed {
	if seed.desktop.seed == nil {
		seed.desktop = seed.clone()
	}
	return seed.desktop
}

//Return the seed that should replace this seed when on ReactNative.
func (seed Seed) ReactNative() Seed {
	if seed.native.seed == nil {
		seed.native = seed.clone()
	}
	return seed.native
}

//Return the seed itself, when embedded in a struct, this is good way to retrieve the original seed.
func (seed Seed) Root() Seed {
	return seed
}

//Return the parent seed.
func (seed Seed) Parent() Seed {
	return seed.parent.Root()
}

func (seed Seed) Copy() Seed {
	var another = *seed.seed

	another.Style = another.Style.Copy()
	another.Portrait = another.Portrait.Copy()
	another.Landscape = another.Landscape.Copy()

	//TODO copy media queries?

	another.id = base64.RawURLEncoding.EncodeToString(big.NewInt(id).Bytes())
	id++
	return Seed{seed: &another}
}

//All seeds have a unique id.
var id int64 = 1

var allSeeds = make(map[string]*seed)

//Create and return a new seed.
func New() Seed {
	s := new(seed)

	//Seed identification is compressed to base64.
	s.id = base64.RawURLEncoding.EncodeToString(big.NewInt(id).Bytes())

	if s.id[0] >= '0' && s.id[0] <= '9' {
		s.id = "_" + s.id
	}

	s.id = strings.Replace(s.id, "-", "__", -1)

	id++

	s.Style = style.New()
	s.Landscape = style.New()
	s.Portrait = style.New()
	s.tag = "div"

	allSeeds[s.id] = s

	//Intial style.
	//seed.SetSize(100, 100)

	return Seed{seed: s}
}

func AddTo(parent Interface) Seed {
	var seed = New()
	parent.Root().Add(seed)
	return seed
}

const OnPress = `op`

//Run a script when this seed is clicked.
func (seed Seed) OnClick(f func(Script)) {
	seed.onclick = f
	seed.OnReady(func(q Script) {
		q.Javascript(OnPress + "('" + seed.id + "', function(event) {")
		f(q)
		q.Javascript("});")
	})
}

//Run a script when this seed is clicked, allows click to propagate to other scripts.
func (seed Seed) OnClickThrough(f func(Script)) {
	seed.onclick = f
	seed.OnReady(func(q Script) {
		q.Javascript(OnPress + "('" + seed.id + "', function(event) {")
		f(q)
		q.Javascript("}, true);")
	})
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

func (seed Seed) Page() bool {
	return seed.page
}

func (seed Seed) Require(script string) {
	seed.scripts = append(seed.scripts, script)
	NewAsset(script).AddTo(seed)
}

//Add a child seed to this seed.
func (seed Seed) Add(child Interface) {
	seed.children = append(seed.children, child)
	child.Root().parent = seed
	if seed.template {
		child.Root().template = true
	}
}

//Add a handler to the seed, when this seed is launched as root, the handlers will be executed for each incomming request.
func (seed Seed) AddHandler(handler func(w http.ResponseWriter, r *http.Request)) {
	seed.handlers = append(seed.handlers, handler)
}

//Add text, html or whatever!
func (seed Seed) SetContent(data string) {
	seed.content = []byte(data)
}

//Set the text content of the seed.
func (seed Seed) Text() string {
	return string(seed.content)
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

//Shorthand for seed.OnClick(func(q seed.Script){ page.Script(q).Goto() })
func (seed Seed) OnSwipeLeft(f func(Script)) {
	seed.Require("hammer.js")
	seed.OnReady(func(q Script) {
		q.Javascript("{")
		q.Javascript("let hammertime = new Hammer(" + seed.Script(q).Element() + ");")
		q.Javascript(`hammertime.on("swipeleft", function() {`)
		f(q)
		q.Javascript("});")
		q.Javascript("}")
	})
}

//Shorthand for seed.OnClick(func(q seed.Script){ page.Script(q).Goto() })
func (seed Seed) OnSwipeRight(f func(Script)) {
	seed.Require("hammer.js")
	seed.OnReady(func(q Script) {
		q.Javascript("{")
		q.Javascript("let hammertime = new Hammer(" + seed.Script(q).Element() + ");")
		q.Javascript(`hammertime.on("swiperight", function() {`)
		f(q)
		q.Javascript("});")
		q.Javascript("}")
	})
}

//Run a script when this seed is clicked.
func (seed Seed) OnFocus(f func(Script)) {
	seed.OnReady(func(q Script) {
		q.Javascript("{")
		q.Javascript(`let onfocus = function(ev) {`)
		f(q)
		q.Javascript(`};`)
		q.Javascript(seed.Script(q).Element() + `.onfocus = onfocus;`)
		q.Javascript("}")
	})
}

//Run a script when this seed is clicked.
func (seed Seed) OnInput(f func(Script)) {
	seed.OnReady(func(q Script) {
		q.Javascript("{")
		q.Javascript(`let oninput = function(ev) {`)
		f(q)
		q.Javascript(`};`)
		q.Javascript(seed.Script(q).Element() + `.oninput = oninput;`)
		q.Javascript("}")
	})
}

//Run a script when this seed is clicked.
func (seed Seed) OnEnter(f func(Script)) {
	seed.OnReady(func(q Script) {
		q.Javascript("{")
		q.Javascript(`let onenter = function(ev) {if (ev.keyCode == 13 || ev.which == 13){`)
		f(q)
		q.Javascript(`}};`)
		q.Javascript(seed.Script(q).Element() + `.onkeypress = onenter;`)
		q.Javascript("}")
	})
}

//Run a script when this seed is clicked.
func (seed Seed) OnFocusLost(f func(Script)) {
	seed.OnReady(func(q Script) {
		q.Javascript("{")
		q.Javascript(`let onfocuslost = function(ev) {`)
		f(q)
		q.Javascript(`};`)
		q.Javascript(seed.Script(q).Element() + `.onfocusout = onfocuslost;`)
		q.Javascript("}")
	})
}

//Run a script when this seed is clicked.
func (seed Seed) OnLongPress(f func(Script)) {
	seed.OnReady(func(q Script) {
		q.Javascript("{")
		q.Javascript(`let onlongpress = function(ev) {ev.preventDefault()`)
		f(q)
		q.Javascript(`};`)
		q.Javascript(seed.Script(q).Element() + `.addEventListener('long-press', onlongpress);`)
		q.Javascript("}")
	})
}
