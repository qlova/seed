package script

import qlova "github.com/qlova/script"

//URL is a reference to a resource/location on the internet.
type URL struct {
	location qlova.String
	q        Script
}

//Website is an alias to URL.
type Website = URL

//URL returns a new URL based on the given String.
func (script Script) URL(location qlova.String) URL {
	return Website{
		q:        script,
		location: location,
	}
}

//Website returns a new URL based on the given String.
func (script Script) Website(location qlova.String) URL {
	return script.URL(location)
}

//Open opens the URL in a new tab.
func (url URL) Open() {
	url.q.Javascript("window.open(" + string(url.location.LanguageType().Raw()) + ");")
}

//Goto goes directly to the URL.
func (url URL) Goto() {
	url.q.Javascript("window.location.href = (" + string(url.location.LanguageType().Raw()) + ");")
}
