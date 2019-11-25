package internal

import (
	//Global ids.
	"encoding/base64"
	"math/big"

	"github.com/qlova/seed/style"
)

//Context is a global app context.
type Context struct {
	*context
}

//NewContext returns a new app context.
func NewContext() Context {
	return Context{&context{
		Dependencies: make(map[string]struct{}),
		Animations:   make(map[*Animation]string),
		FontCache:    make(map[string]style.Font),
		Pages:        make(map[string]Page),
	}}
}

type Page interface{}

type context struct {
	Dependencies map[string]struct{}
	Animations   map[*Animation]string
	FontCache    map[string]style.Font
	Pages        map[string]Page
}

var animationID int64 = 1

//Animation adds an animation to the global Context if it doesn't already exist.
func (context Context) Animation(animation *Animation) string {
	if id, ok := context.Animations[animation]; ok {
		return id
	}

	var id = base64.RawURLEncoding.EncodeToString(big.NewInt(animationID).Bytes())

	animationID++

	context.Animations[animation] = id

	return id
}

//AddPage adds a page to the context.
func (context Context) AddPage(id string, page interface{}) {
	context.Pages[id] = page
}

//ClearPages clears the current pages.
func (context Context) ClearPages() {
	context.Pages = make(map[string]Page)
}
