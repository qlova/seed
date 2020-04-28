package js

func Throw(v AnyValue) Script {
	return func(q Ctx) {
		q("throw ")
		q(v.GetValue().String())
		q(";")
	}
}
