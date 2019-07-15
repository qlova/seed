package script

import qlova "github.com/qlova/script"

type time struct {
	Script
}

type Time struct {
	Q Script
	qlova.Native
}

func (q time) Now() Time {
	return Time{q.Script, q.Value("(new Date())").Native()}
}

func (time Time) String() String {
	return time.Q.Value(time.LanguageType().Raw() + ".toString()").String()
}
