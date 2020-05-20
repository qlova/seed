package js

//Element is the most general base class from which all element objects (i.e. objects that represent elements) in a Document inherit. It only has methods and properties common to all kinds of elements. More specific classes inherit from Element. For example, the HTMLElement interface is the base interface for HTML elements, while the SVGElement interface is the basis for all SVG elements. Most functionality is specified further down the class hierarchy.
type Element struct {
	Value
}

//RequestFullscreen requests that the user agent switch from full-screen mode back to windowed mode.
func (element Element) RequestFullscreen() Script {
	return element.Run("requestFullscreen")
}
