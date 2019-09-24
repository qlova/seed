package seed

import (
	"github.com/qlova/seed/style"
	"github.com/qlova/seed/style/css"
	"github.com/qlova/seed/user"
	"gopkg.in/russross/blackfriday.v2"

	"encoding/base64"
	"html"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	HTML "github.com/qlova/seed/html"
)

//Dir is the working directory of the seed.
var Dir = filepath.Dir(os.Args[0])

//User is an alias to the user.User type.
type User = user.User

//Shadow is an alias to the style.Shadow type.
type Shadow = style.Shadow

//UserData returns a reference to a new type of user data, this is small data that is used to identify the user.
func UserData() user.Data {
	return user.DataType()
}

//Vm is a unit relative to the viewport size.
const Vm Unit = style.Vm

//Em is a unit  value is relative to the current font size.
const Em Unit = style.Em

//Px is a unit (pixels)
const Px Unit = style.Px

//Auto is a unit where appropriate will automatically select a suitable value.
const Auto Unit = style.Auto

//Direction constants.
const (
	Top    = style.Top
	Bottom = style.Bottom
	Left   = style.Left
	Right  = style.Right
	Center = style.Center
)

//Arial is a default arial font.
var Arial = style.Font{
	FontFace: css.FontFace{
		FontFamily: "Arial",
	},
}

//Font is a font type.
type Font struct {
	style.Font
	path string
}

//FontCache caches fonts that have been registered.
var FontCache = make(map[string]Font)

//NewFont registers the font and creates it.
func NewFont(path string) Font {
	if font, ok := FontCache[path]; ok {
		return font
	}

	var font = Font{style.NewFont(path), path}
	FontCache[path] = font
	return font
}

//SetFont sets the font of the specified seed.
func (seed Seed) SetFont(font Font) {
	seed.font = font.Font
	NewAsset(font.path).AddTo(seed)
	seed.Style.SetFont(font.Font)
}

//Seed is a component of an app.
//It can contain children seeds and have styles, scripts attached to it.
type Seed struct {
	*seed
}

type seed struct {
	HTML.Element
	style.Style
	dynamic

	id               string
	tag, attr, class string
	children         []Interface

	attributes map[string]string

	styled bool
	ready  bool

	font      style.Font
	animation Animation

	content []byte
	page    bool
	splash  bool

	onclick  func(Script)
	onchange func(Script)
	onready  func(Script)

	on map[string]func(Script)

	template bool

	parent Interface

	//This is a list of scripts that are needed by this seed.
	//eg. []string{"jquery.js"}
	scripts []string

	handlers []func(w http.ResponseWriter, r *http.Request)

	Landscape, Portrait style.Style

	screenSmallerThan, screenGreaterThan map[Unit]style.Style

	desktop, mobile, tablet, watch, tv, native Seed

	app *App

	assets []Asset

	states map[State]func(Script)

	tags map[string]bool
}

//On runs a script callback when the specified event is fired.
func (seed Seed) On(event string, callback func(Script)) {
	if seed.on == nil {
		seed.on = make(map[string]func(Script))
	}

	if original, ok := seed.on[event]; ok {
		seed.on[event] = func(q Script) {
			original(q)
			callback(q)
		}
		return
	}
	seed.on[event] = callback
}

//Is returns true if a and b are the same seed.
func (seed Seed) Is(b Interface) bool {
	return seed.seed == b.Root().seed
}

//Unit is a measurement for styles.
type Unit = complex128

//ScreenSmallerThan is a conditional query that returns a style that effects that condition.
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

//Null returns true if the seed is null.
func (seed Seed) Null() bool {
	return seed.seed == nil
}

//MarshalText marshals this seed as a HTML id.
func (seed Seed) MarshalText() ([]byte, error) {
	return []byte("#" + seed.ID()), nil
}

//ID returns the id of this seed.
func (seed Seed) ID() string {
	if seed.seed == nil {
		return ""
	}
	return seed.id
}

//Children returns a slice of children of this seed.
func (seed Seed) Children() []Interface {
	return seed.children
}

//SetClass sets the HTML class of this seed [DEPRECIATED].
func (seed Seed) SetClass(class string) {
	seed.class = class
}

//Tag returns the HTML tag of this seed [DEPRECIATED].
func (seed Seed) Tag() string {
	return seed.tag
}

//SetTag sets the HTML tag for this seed [DEPRECIATED].
func (seed Seed) SetTag(tag string) {
	seed.tag = tag
}

//SetAttributes sets the HTML attributes for this seed [DEPRECIATED].
func (seed Seed) SetAttributes(attr string) {
	seed.attr = attr
}

//Attributes returns the HTML attributes for this seed [DEPRECIATED].
func (seed Seed) Attributes() string {
	return seed.attr
}

//AddTo adds this seed to the specified parent.
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

//Desktop returns the seed that should replace this seed when on the Desktop.
func (seed Seed) Desktop() Seed {
	if seed.desktop.seed == nil {
		seed.desktop = seed.clone()
	}
	return seed.desktop
}

//ReactNative returns the seed that should replace this seed when on ReactNative [EXPERIMENTAL].
func (seed Seed) ReactNative() Seed {
	if seed.native.seed == nil {
		seed.native = seed.clone()
	}
	return seed.native
}

//Root returns the seed itself, when embedded in a struct, this is good way to retrieve the original seed.
func (seed Seed) Root() Seed {
	return seed
}

//Parent returns the parent seed.
func (seed Seed) Parent() Seed {
	return seed.parent.Root()
}

//Copy creates a copy of this seed.
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

//New create and return a new seed.
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

//AddTo creates and returns a seed attached to the specified parent.
func AddTo(parent Interface) Seed {
	var seed = New()
	parent.Root().Add(seed)
	return seed
}

//OnPress is the JS required for OnPress handling.
const OnPress = `
	function op(element, func, propagate) {
		let handler = function(event) {
			func(event);
		};
		
		let moved = false;
		let point = [0, 0];
		
		element.ontouchstart = function(e) {
			var changedTouch = event.changedTouches[0];
				point[0]  = changedTouch.clientX;
				point[1]  = changedTouch.clientY;
		};
		
		element.ontouchmove = function(event) {
			var changedTouch = event.changedTouches[0];
			var elem = document.elementFromPoint(changedTouch.clientX, changedTouch.clientY);
						
			if (elem != event.target) moved = true;
						
			let a = changedTouch.clientX - point[0];
			let b = changedTouch.clientY - point[1];
			if ((a*a + b*b) > 50*50) moved = true;
		};
		
		element.ontouchend = function(ev) {
			if (ev.stopPropagation && !propagate) ev.stopPropagation(); 
			ev.preventDefault(); 
			if (moved) {
				moved = false; 
				return; 
			}
			ev = ev.changedTouches[0];
			handler(ev);
		};

		element.onclick = handler;
	}
`

//OnClick runs a script when this seed is clicked.
func (seed Seed) OnClick(f func(Script)) {
	seed.onclick = f
	seed.OnReady(func(q Script) {
		q.Require(OnPress)
		q.Javascript("op(" + seed.Script(q).Element() + ", function(event) {")
		f(q)
		q.Javascript("});")
	})
}

//OnClickThrough runs a script when this seed is clicked, allows click to propagate to other scripts.
func (seed Seed) OnClickThrough(f func(Script)) {
	seed.onclick = f
	seed.OnReady(func(q Script) {
		q.Require(OnPress)
		q.Javascript("op('" + seed.id + "', function(event) {")
		f(q)
		q.Javascript("}, true);")
	})
}

//OnClickGoto is shorthand for seed.OnClick(func(q seed.Script){ page.Script(q).Goto() })
func (seed Seed) OnClickGoto(page Page) {
	seed.OnClick(func(q Script) {
		page.Script(q).Goto()
	})
}

//OnReady runs a script when this seed is ready/loaded/onload/init.
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

//OnChange runs a script when this seed's value is changed by the user.
func (seed Seed) OnChange(f func(Script)) {
	seed.OnReady(func(q Script) {
		q.Javascript("{")
		q.Javascript(`let onchange = function(ev) {`)
		f(q)
		q.Javascript(`};`)
		q.Javascript(seed.Script(q).Element() + `.onchange = onchange;`)
		q.Javascript("}")
	})
}

//Page returns true if this seed is a page.
func (seed Seed) Page() bool {
	return seed.page
}

//Require requires an external script needed by this seed.
func (seed Seed) Require(script string) {
	seed.scripts = append(seed.scripts, script)
	NewAsset(script).AddTo(seed)
}

//Add a child seed to this seed.
func (seed Seed) Add(child Interface) {
	seed.children = append(seed.children, child)
	child.Root().parent = seed
	child.Root().app = seed.app
	if seed.template {
		child.Root().template = true
	}
}

//AddHandler adds a handler to the seed, when this seed is launched as root, the handlers will be executed for each incomming request.
func (seed Seed) AddHandler(handler func(w http.ResponseWriter, r *http.Request)) {
	seed.handlers = append(seed.handlers, handler)
}

//SetContent adds text, html or whatever!
func (seed Seed) SetContent(data string) {
	seed.content = []byte(data)
}

//Text returns the text content of the seed.
func (seed Seed) Text() string {
	return string(seed.content)
}

//SetText sets the text content of this seed.
func (seed Seed) SetText(data string) {
	data = html.EscapeString(data)
	data = strings.Replace(data, "\n", "<br>", -1)
	data = strings.Replace(data, "  ", "&nbsp;&nbsp;", -1)
	data = strings.Replace(data, "\t", "&emsp;", -1)
	seed.content = []byte(data)
}

//SetMarkdown sets the content of the seed in markdown format.
func (seed Seed) SetMarkdown(data string) {
	seed.content = blackfriday.Run([]byte(data))
}

//OnSwipeLeft runs a script when swiped left.
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

//OnSwipeRight runs a script when swiped right.
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

//OnFocus run a script when this seed is focused.
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

//OnInput runs a script when this seed has input.
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

//OnEnter runs a script when this seed has enter input.
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

//OnFocusLost runs a script when this seed has focus lost.
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

//LongPress is the JS required for handling longpreses.
const LongPress = `
!function(t,e){"use strict";function n(){this.dispatchEvent(new CustomEvent("long-press",{bubbles:!0,cancelable:!0})),clearTimeout(o)}var o=null,u="ontouchstart"in t||navigator.MaxTouchPoints>0||navigator.msMaxTouchPoints>0,s=u?"touchstart":"mousedown",i=u?"touchcancel":"mouseout",a=u?"touchend":"mouseup",c=u?"touchmove":"mousemove";"initCustomEvent"in e.createEvent("CustomEvent")&&(t.CustomEvent=function(t,n){n=n||{bubbles:!1,cancelable:!1,detail:void 0};var o=e.createEvent("CustomEvent");return o.initCustomEvent(t,n.bubbles,n.cancelable,n.detail),o},t.CustomEvent.prototype=t.Event.prototype),e.addEventListener(s,function(t){var e=t.target,u=parseInt(e.getAttribute("data-long-press-delay")||"500",10);o=setTimeout(n.bind(e),u)}),e.addEventListener(a,function(t){clearTimeout(o)}),e.addEventListener(i,function(t){clearTimeout(o)}),e.addEventListener(c,function(t){clearTimeout(o)})}(this,document);
`

//OnLongPress runs a script when this seed is long-clicked.
func (seed Seed) OnLongPress(f func(Script)) {
	seed.OnReady(func(q Script) {
		q.Require(LongPress)
		q.Javascript("{")
		q.Javascript(`let onlongpress = function(ev) {ev.preventDefault()`)
		f(q)
		q.Javascript(`};`)
		q.Javascript(seed.Script(q).Element() + `.addEventListener('long-press', onlongpress);`)
		q.Javascript("}")
	})
}
