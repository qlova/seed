package script

//Feed is a seed.Feed, containing feed data.
type Feed struct {
	Seed
}

//Refresh the feeds content from the server.
func (feed Feed) Refresh() {
	feed.Q.Javascript(feed.Element() + ".onready();")
}

//Data returns the data associated with this feed for the current template.
func (feed Feed) Data(key ...string) String {
	if len(key) > 0 {
		return feed.wrap(`data["` + key[0] + `"]`)
	}
	return feed.wrap(`data`)
}
