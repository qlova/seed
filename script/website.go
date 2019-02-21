package script

import qlova "github.com/qlova/script"

type Website struct {
	location qlova.String
	q Script
}

func (script Script) Website(location qlova.String) Website {
	return Website{
		q: script,
		location: location,
	}
}

func (website Website) Open() {
	website.q.Javascript("window.open("+string(website.location.LanguageType().Raw())+");")
}