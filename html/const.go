package html

//TargetType is a valid hypertext-reference target type.
type TargetType string

//TargetTypes
const (
	Self   TargetType = "_self"
	Blank  TargetType = "_blank"
	Parent TargetType = "_parent"
	Top    TargetType = "_top"
)
