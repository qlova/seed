package client

import "qlova.org/seed/use/js"

//IfElseChain allows else to be called on it.
type IfElseChain struct {
	Script
}

//Else runs the provided scripts if the preceding conditions were false.
func (chain IfElseChain) Else(do ...Script) Script {
	return chain.Script.GetScript().Else(NewScript(do...).GetScript())
}

//If runs the provided scripts if the clients condition is true.
func If(condition Bool, do ...Script) IfElseChain {
	return IfElseChain{js.If(condition, NewScript(do...).GetScript())}
}
