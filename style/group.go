package style

//Group is a complete styleset with default + portrait and landscape
type Group struct {
	Style

	Portrait, Landscape Style
}

//Init initializes the group.
func (group *Group) Init() {
	group.Style = New()
	group.Portrait = New()
	group.Landscape = New()
}
