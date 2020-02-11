package script

import qlova "github.com/qlova/script"

//URL is a reference to a resource/location on the internet.
type URL struct {
	location qlova.String
	q        Ctx
}

//Website is an alias to URL.
type Website = URL

//URL returns a new URL based on the given String.
func (script Ctx) URL(location qlova.String) URL {
	return Website{
		q:        script,
		location: location,
	}
}

//Website returns a new URL based on the given String.
func (script Ctx) Website(location qlova.String) URL {
	return script.URL(location)
}

const openURI = `function openURI(uri) {
	let splits = uri.split(':');
	if (splits.length > 0 && splits[0] != "https" && splits[0] != "http") {
		window.location.href = uri;
		return;
	} 
	window.open(uri, '_blank', 'noopener');
}`

//Open opens the URL in a new tab if possible.
func (url URL) Open() {
	url.q.Require(openURI)
	url.q.Javascript("openURI(" + string(url.location.LanguageType().Raw()) + ");")
}

//Goto goes directly to the URL if possible..
func (url URL) Goto() {
	url.q.Javascript(`if (window.LocalhostWebsocket) LocalhostWebsocket.send("I'll be back"); window.location.href = (` + string(url.location.LanguageType().Raw()) + ");")
}
