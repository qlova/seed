package html

//Attribute is an HTML element attribute.
type Attribute string

//Collection of attributes.
const (
	Source Attribute = "src"
)

//Tag is an HTML tag.
type Tag string

//Collection of HTML tags.
const (
	HTML    Tag = "html"
	Divider     = "div"
)

//Element is an HTML element.
type Element struct {
	Tag
	Attributes map[Attribute]string
}

//Set the attribute of an HTML element.
func (element *Element) Set(attribute Attribute, value string) {
	if element.Attributes == nil {
		element.Attributes = make(map[Attribute]string)
	}
	element.Attributes[attribute] = value
}
