package style

//Group is a complete styleset with default + portrait and landscape
type Group struct {
	Style

	Portrait, Landscape Style
}

//Init initialises the styleset.
func (group *Group) Init() {
	group.Style = New()
	group.Portrait = New()
	group.Landscape = New()
}
