package internal

import (
	//Global ids.
	"encoding/base64"
	"math/big"
)

type Context struct {
	*context
}

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

var animation_id int64 = 1

func (context Context) Animation(animation *Animation) string {
	if id, ok := context.Animations[animation]; ok {
		return id
	}

	var id = base64.RawURLEncoding.EncodeToString(big.NewInt(animation_id).Bytes())

	animation_id++

	context.Animations[animation] = id

	return id
}
