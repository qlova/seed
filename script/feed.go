package script

type Feed struct {
	Seed
}

//Refresh the feeds content from the server.
func (feed Feed) Refresh() {
	feed.Q.Javascript(feed.Element()+".onready();")
}

//Set the index of this feed, ready for the next refresh.
func (feed Feed) SetIndex(index String) {
	feed.Q.Javascript(feed.Element()+".index = "+raw(index)+";")
	feed.Q.Javascript(`window.localStorage.setItem("`+feed.ID+`_index", `+raw(index)+`);`)
}