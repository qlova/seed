package style

//Group is a complete styleset with default + portrait and landscape
type Group struct {
	Style

	Portrait, Landscape Style
}

//NewGroup returns a pointer to a new initialised style.Group.
func NewGroup() *Group {
	var g Group
	g.Init()
	return &g
}

//Init initializes the group.
func (group *Group) Init() {
	group.Style = New()
	group.Portrait = New()
	group.Landscape = New()
	group.selectors = make(map[string]Style)
}
