package script

import "fmt"

//Dynamic can be any script.Type
type Dynamic struct {
	Q Ctx
	Native
}

func (d Dynamic) String() String {
	return d.Q.Value(fmt.Sprintf(`(%v||"")`,
		d.LanguageType().Raw())).String()
}

//Object returns the dynamic value as an object.
func (d Dynamic) Object() Object {
	return d.Q.Value(fmt.Sprintf(`(%v||{})`,
		d.LanguageType().Raw())).Object()
}

//Array returns the dynamic value as an array.
func (d Dynamic) Array() Array {
	return d.Q.Value(fmt.Sprintf(`(%v||[])`,
		d.LanguageType().Raw())).Array()
}
