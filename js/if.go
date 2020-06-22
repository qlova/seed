package js

type ElseIfChain struct {
	q      Ctx
	ignore bool
}

func (e ElseIfChain) Else(do Script) {
	if e.ignore {
		return
	}
	q := e.q
	q(" else {")
	do(q)
	q("}")
}

func (q Ctx) If(condition AnyBool, do Script) ElseIfChain {
	if condition == nil {
		return ElseIfChain{q, true}
	}

	q("if(")
	q(condition.GetBool())
	q(") {")
	q(do)
	q("}")

	return ElseIfChain{q, false}
}

func If(condition AnyBool, do Script) Script {
	if condition == nil {
		return nil
	}
	return func(q Ctx) {
		q.If(condition, do)
	}
}

func (s Script) Else(do Script) Script {
	if s == nil {
		return nil
	}
	return func(q Ctx) {
		s(q)
		q("else {")
		do(q)
		q("}")
	}
}
