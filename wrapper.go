package seed

import "reflect"
import "github.com/qlova/script"
//import "github.com/qlova/script/language"
//import "github.com/qlova/script/language/javascript"

func wrap(q script.Script) Script {
	return Script{seedScript: &seedScript{ Script:q }}
}

func (q Script) If(condition script.Boolean, block func(Script), ifelsechain ...script.IfElseChain) {
	var chain script.IfElseChain
	
	switch v := ifelsechain[0].(type) {
		case ElseIfChain:
			chain = script.ElseIfChain{ Chain: v.Chain }
		case *script.EndChain:
			chain = v
		default:
			panic(reflect.TypeOf(ifelsechain[0]).String())
	}
	
	q.Script.If(condition, func(q script.Script) {
		
		block(wrap(q))
		
	}, chain)
}

type ElseIfChain struct {
	script Script
	script.ElseIfChain
}

func (q Script) ElseIf(condition script.Boolean, block func(Script)) *ElseIfChain {
	var chain = new(ElseIfChain)
	chain.script = q
	return q.ElseIf(condition, block)
}

func (chain *ElseIfChain) ElseIf(condition script.Boolean, block func(Script)) *ElseIfChain {
	
	var elseifchain = script.ElseIfChain{ Chain: chain.Chain }

	elseifchain.ElseIf(condition, func(q script.Script) {
		block(wrap(q))
	})
	
	chain.Chain = elseifchain.Chain
	return chain
}


func (q Script) Else(block func(Script)) *script.EndChain {
	return new(ElseIfChain).Else(block)
}

func (chain *ElseIfChain) Else(block func(Script)) *script.EndChain {
	
	var elseifchain = script.ElseIfChain{ Chain: chain.Chain }

	return elseifchain.Else(func(q script.Script) {
		block(wrap(q))
	})
}
