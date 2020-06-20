package html

import (
	"bytes"

	"qlova.org/seed"
)

type Document struct {
	seed.Seed
	Head, Body seed.Seed
}

func New() Document {
	html := seed.New(SetTag("html"))
	head := seed.New(SetTag("head"))
	body := seed.New(SetTag("body"))

	html = html.With(
		head, body,
	)

	return Document{html, head, body}
}

func (doc Document) Render() []byte {
	var b bytes.Buffer

	b.WriteString("<!DOCTYPE html>")
	b.Write(Render(doc))

	return b.Bytes()
}
