package js

type ElseIfChain struct {
	q Ctx
}

func (e ElseIfChain) Else(do Script) {
	q := e.q
	q(" else {")
	do(q)
	q("}")
}

func (q Ctx) If(condition AnyBool, do Script) ElseIfChain {
	q("if(")
	q(condition.GetBool())
	q(") {")
	do(q)
	q("}")

	return ElseIfChain{q}
}
