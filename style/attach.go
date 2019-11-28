package style

import "github.com/qlova/seed/style/css"

//Attacher is used to specify an attachpoint.
type Attacher struct {
	style Style
}

//Top attaches this to the top of the context.
func (a Attacher) Top() Attacher {
	a.style.CSS().SetTop(css.Zero)
	return a
}

//Bottom attaches this to the bottom of the context.
func (a Attacher) Bottom() Attacher {
	a.style.CSS().SetBottom(css.Zero)
	return a
}

//Left attaches this to the left of the context.
func (a Attacher) Left() Attacher {
	a.style.CSS().SetLeft(css.Zero)
	return a
}

//Right attaches this to the right of the context.
func (a Attacher) Right() Attacher {
	a.style.CSS().SetRight(css.Zero)
	return a
}

//AttachToScreen attaches this element to the screen, returns an attacher to specify where.
func (style Style) AttachToScreen() Attacher {
	style.CSS().SetPosition(css.Fixed)
	return Attacher{style}
}

//AttachToParent attaches this element to its parent, returns an attacher to specify where.
func (style Style) AttachToParent() Attacher {
	style.CSS().SetPosition(css.Absolute)
	return Attacher{style}
}

//StickyToScreen sticks this element to the screen, returns an attacher to specify where.
func (style Style) StickyToScreen() Attacher {
	style.CSS().SetPosition(css.Sticky)
	return Attacher{style}
}

//Detach detaches the element from the parent and/or the screen.
//This causes the element to behave like a default element.
func (style Style) Detach() {
	style.CSS().SetPosition(css.Relative)
}
