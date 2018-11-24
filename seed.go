package seed

import "github.com/qlova/seed/style"
import "github.com/qlova/seed/manifest"
import "github.com/qlova/seed/interfaces"

import (
	"bytes"
	"net/http"
	"math/big"
	"encoding/base64"
)

//#seedsafe
type Slice []*seed
type Seed struct {
	*seed
}

type seed struct {
	style.Style
	
	id string
	tag, attr, class string
	children []interfaces.App
	
	styled bool
	
	fonts bytes.Buffer
	
	content []byte
	page bool
	
	onclick func(Script)
	onchange func(Script)
	onready func(Script)
	
	parent interfaces.App
	
	//This is a list of scripts that are needed by this seed.
	//eg. []string{"jquery.js"}
	scripts []string
	
	manifest manifest.Manifest
	handlers []func(w http.ResponseWriter, r *http.Request)
	
	dynamicText func(Client)
}

//All seeds have a unique id.
var id int64 = 1;

//Create and return a new seed.
func New() Seed {
	seed := new(seed)
	
	//Seed identification is compressed to base64.
	seed.id = base64.RawURLEncoding.EncodeToString(big.NewInt(id).Bytes())
	id++

	seed.Style = style.New()	
	seed.tag = "div"
	
	//All seeds have the potential to be the root seed, so they all need a minimal viable manifest.
	seed.manifest = manifest.New()
	
	seed.SetMargin("0")
	seed.SetWidth("100%")
	seed.SetHeight("100%")

	return Seed{seed:seed}
}

func (seed Seed) getScripts() []string {
	var scripts = seed.scripts

	for _, child := range seed.children {
		scripts = append(scripts, child.(Seed).getScripts()...)
	}
	
	return scripts
}

func (seed Seed) Scripts() map[string]struct{} {
	
	var scripts = seed.getScripts()
	var uniques = make(map[string]struct{})

	for _, script := range scripts {
		uniques[script] = struct{}{}
	}

	return uniques
}

func (seed Seed) buildOnReady(buffer *bytes.Buffer) {
	
	if seed.onready != nil {
		buffer.WriteByte('{')
		buffer.Write(toJavascript(seed.onready))
		buffer.WriteByte('}')
	}
	
	for _, child := range seed.children {
		child.(Seed).buildOnReady(buffer)
	}
}


func (seed Seed) BuildOnReady() []byte {
	var buffer bytes.Buffer
	buffer.WriteString(`document.addEventListener('DOMContentLoaded', function() {`)
	
	seed.buildOnReady(&buffer)
	
	buffer.WriteString(`}, false);`)
	return buffer.Bytes()
}
