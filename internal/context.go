package internal

import (
	//Global ids.
	"encoding/base64"
	"math/big"
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
	}}
}

type context struct {
	Dependencies map[string]struct{}
	Animations   map[*Animation]string
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
