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
