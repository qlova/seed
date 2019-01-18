package seed

import "github.com/qlova/seed/style/css"
import "github.com/qlova/seed/style"
import "github.com/qlova/seed/manifest"

import (
	"net/http"
	"math/big"
	"encoding/base64"
	"strings"
)

const Vm = style.Vm
const Em = style.Em
const Px = style.Px
const Top = style.Top
const Bottom = style.Bottom
const Left = style.Left
const Right = style.Right
const Auto = style.Auto
const Center = style.Center

var Arial = style.Font{
	FontFace: css.FontFace{
		FontFamily: "Arial",
	},
}

func Font(path string) style.Font {
	RegisterAsset(path)
	
	return style.NewFont(path)
}

func (seed Seed) SetFont(font style.Font) {
	seed.font = font
	seed.Style.SetFont(font)
}


//#seedsafe
type Slice []*seed
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
	
	manifest manifest.Manifest
	handlers []func(w http.ResponseWriter, r *http.Request)
	
	dynamicText func(Client)

	Landscape, Portrait style.Style

	desktop, mobile, tablet, watch, tv Seed
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

func (seed Seed) GetSeed() Seed {
	return seed
}

func (seed Seed) Parent() Seed {
	return seed.parent.GetSeed()
}

func (seed Seed) Child(number int) Seed {
	return seed.children[number-1].(Seed)
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
	seed := new(seed)
	
	//Seed identification is compressed to base64.
	seed.id = base64.RawURLEncoding.EncodeToString(big.NewInt(id).Bytes())

	if seed.id[0] >= '0' && seed.id[0] <= '9' {
		seed.id = "_"+seed.id
	}
	
	seed.id = strings.Replace(seed.id, "-", "__", -1)
	
	id++

	seed.Style = style.New()
	seed.Landscape = style.New()
	seed.Portrait = style.New()
	seed.tag = "div"
	
	//All seeds have the potential to be the root seed, so they all need a minimal viable manifest.
	seed.manifest = manifest.New()

	allSeeds[seed.id] = seed

	//Intial style.
	//seed.SetSize(100, 100)
	
	return Seed{seed:seed}
}
