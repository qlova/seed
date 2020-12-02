package js

func Try(s Script) Script {
	return func(q Ctx) {
		q("try {")
		q(s)
		q("}")
	}
}

func Throw(v AnyValue) Script {
	return func(q Ctx) {
		q("throw ")
		q(v.GetValue().String())
		q(";")
	}
}

func (s Script) Catch(do Script, e ...string) Script {

	var err = "err"
	if len(e) > 0 {
		err = e[0]
	}

	return func(q Ctx) {
		s(q)
		q("catch(" + err + ") {")
		do(q)
		q("}")
	}
}
