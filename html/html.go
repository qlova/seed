package html

//Attribute is an HTML element attribute.
type Attribute string

//Language is a valid language attribute.
type Language string

//LinkType indicate the relationship between two documents.
//See https://developer.mozilla.org/en-US/docs/Web/HTML/Link_types
type LinkType string

//Collection of attributes.
const (
	Source          Attribute = "src"
	Value           Attribute = "value"
	AllowFullscreen Attribute = "allowfullscreen"
	Frameborder     Attribute = "frameborder"

	Download                   Attribute = "download"
	HypertextReference         Attribute = "href"
	HypertextReferenceLanguage Attribute = "hreflang"
	Relationship               Attribute = "rel"
	Target                     Attribute = "target"
	Type                       Attribute = "type"
	Ping                       Attribute = "ping"
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

	HTML []byte
}

//Set the attribute of an HTML element.
func (element *Element) Set(attribute Attribute, value ...string) {
	if element.Attributes == nil {
		element.Attributes = make(map[Attribute]string)
	}
	var v string
	if len(value) > 0 {
		v = value[0]
	}
	element.Attributes[attribute] = v
}

//SetHTML from a string. Shorthand for element.HTML = []byte(html)
func (element *Element) SetHTML(html string) {
	element.HTML = []byte(html)
}
