package script

import "fmt"

//Dynamic can be any script.Type
type Dynamic struct {
	Q Ctx
	Native
}

func (d Dynamic) String() String {
	return d.Q.Value(fmt.Sprintf(`(typeof %v == "string" && %v) || ""`,
		d.LanguageType().Raw(), d.LanguageType().Raw())).String()
}

//Array returns the dynamic value as an array.
func (d Dynamic) Array() Array {
	return d.Q.Value(fmt.Sprintf(`(typeof %v == "array" && %v) || []`,
		d.LanguageType().Raw(), d.LanguageType().Raw())).Array()
}
