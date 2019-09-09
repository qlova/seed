package script

//Feed is a seed.Feed, containing feed data.
type Feed struct {
	Seed
}

//Refresh the feeds content from the server.
func (feed Feed) Refresh() {
	feed.Q.Javascript(feed.Element() + ".onready();")
}

//SetIndex sets the index of this feed, ready for the next refresh.
func (feed Feed) SetIndex(index String) {
	feed.Q.Javascript(feed.Element() + ".index = " + raw(index) + ";")
	feed.Q.Javascript(`window.localStorage.setItem("` + feed.ID + `_index", ` + raw(index) + `);`)
}

//Data returns the data associated with this feed for the current template.
func (feed Feed) Data(key ...string) String {
	if len(key) > 0 {
		return feed.wrap(`data["` + key[0] + `"]`)
	}
	return feed.wrap(`data`)
}
