package the

import (
	"qlova.org/seed/client"
	"qlova.org/seed/use/js"
)

//timeExpression is an expression containing a client.Time.
type timeExpression struct {
	client.Time
	components []client.Value
}

type date struct {
	js.Value
}

func (d date) GetTime() js.Value {
	return d.Value
}

func Time(base client.Time, durations ...client.Duration) client.Time {
	var expression string = "(%v"
	var values = []client.Value{base}

	for _, n := range durations {
		expression += " + %v"
		values = append(values, n)
	}

	expression += ")"

	sum := Number(js.Number{Value: js.NewValue(expression, values...)}, values...)

	return timeExpression{date{sum.GetValue()}, sum.components}
}
