package meta

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/html/attr"
)

//New returns an HTML meta element.
func Key(name string, content string) seed.Seed {
	return New(attr.Set("name", name), attr.Set("content", content))
}

func New(options ...seed.Option) seed.Seed {
	return seed.New(html.SetTag("meta"), seed.Options(options))
}

//Charset returns an HTML meta element with charset set to the given string.
func Charset(charset string) seed.Seed {
	return New(attr.Set("charset", charset))
}

//Description returns an HTML meta element with description set to the given string.
func Description(description string) seed.Seed {
	return New(attr.Set("description", description))
}
