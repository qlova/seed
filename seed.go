package seed

import "github.com/qlova/seed/style/css"
import "github.com/qlova/seed/style"
import "github.com/qlova/seed/user"

import (
	"net/http"
	"math/big"
	"encoding/base64"
	"strings"
	"html"

	"path/filepath"
	"os"
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

func NewFont(path string) Font {
	return Font{style.NewFont(path), path}
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
	
	id string
	tag, attr, class string
	children []Interface
	
	styled bool
	
	font style.Font
	animation Animation
	
	content []byte
	page bool
	
	onclick func(Script)
	onchange func(Script)
	onready func(Script)
	
	parent Interface
	
	//This is a list of scripts that are needed by this seed.
	//eg. []string{"jquery.js"}
	scripts []string

	handlers []func(w http.ResponseWriter, r *http.Request)
	
	dynamicText func(User)

	Landscape, Portrait style.Style

	desktop, mobile, tablet, watch, tv, native Seed
	
	app *App
	
	assets []Asset
}

func (seed Seed) ID() string {
	return seed.id
}

func (seed Seed) SetClass(class string) {
	seed.class = class
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
	
	var OldStyleImplemenation = another.Style.Style.Stylable.(css.Implementation)
	var NewStyleImplementation = make(css.Implementation, len(OldStyleImplemenation))
	
	for key := range OldStyleImplemenation {
		NewStyleImplementation[key] = OldStyleImplemenation[key]
	}
	another.Style.Style.Stylable = NewStyleImplementation
	
	another.id =  base64.RawURLEncoding.EncodeToString(big.NewInt(id).Bytes())
	id++
	return Seed{ seed: &another }
}

//All seeds have a unique id.
var id int64 = 1;

var allSeeds = make(map[string]*seed)

//Create and return a new seed.
func New() Seed {
	s := new(seed)
	
	//Seed identification is compressed to base64.
	s.id = base64.RawURLEncoding.EncodeToString(big.NewInt(id).Bytes())

	if s.id[0] >= '0' && s.id[0] <= '9' {
		s.id = "_"+s.id
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
	
	return Seed{seed:s}
}

func AddTo(parent Interface) Seed {
	var seed = New()
	parent.Root().Add(seed)
	return seed
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

//Review these methods:

//TODO Should be Internal.
func (seed Seed) SetPlaceholder(placeholder string) {
	seed.attr += " placeholder='"+placeholder+"' "
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
