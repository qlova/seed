//Package paymentbox provides a (stripe) secured input that accepts payment details.
package paymentbox

import (
	"qlova.org/seed"
	"qlova.org/seed/js"
	"qlova.org/seed/s/html/div"
	"qlova.org/seed/script"
	"qlova.org/seed/signal"
	"qlova.org/seed/state"
	"qlova.org/seed/user"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/setupintent"
)

var StripePublishableKey string

type StripeBox struct {
	BillingName, Value state.String
	confirmCardSetup   signal.Type
	confirmedCardSetup signal.Type
}

func NewStripe() StripeBox {
	return StripeBox{
		BillingName:        state.NewString("", state.Session()),
		Value:              state.NewString("", state.Session()),
		confirmCardSetup:   signal.New(),
		confirmedCardSetup: signal.New(),
	}
}

//ConfirmCardSetup tells the stripebox to confirm the user provided details.
func (s StripeBox) ConfirmCardSetup() script.Script {
	return signal.Emit(s.confirmCardSetup)
}

//OnCardConfirmed runs when the card has been confirmed.
func (s StripeBox) OnCardConfirmed(do script.Script) seed.Option {
	return signal.On(s.confirmedCardSetup, do)
}

//New returns a new stripe payment box.
func (s StripeBox) New(options ...seed.Option) seed.Seed {

	var PaymentBox = seed.NewLink()

	var Stripe = js.Function{js.NewValue(`Stripe`)}

	return div.New(
		js.Require("https://js.stripe.com/v3/", ""),
		PaymentBox.Link(),

		script.OnReady(func(q script.Ctx) {
			var element = script.Element(PaymentBox).Var(q)
			q(element.Set(`stripe`, js.Call(Stripe, js.NewString(StripePublishableKey))))
			var elements = element.Get("stripe").Call("elements").Var(q)
			q(element.Set(`card`, elements.Call(`create`, js.NewString("card"), js.NewObject(make(map[string]js.AnyValue)))))
			q(element.Get(`card`).Run(`mount`, element))
		}),

		signal.On(s.confirmCardSetup, func(q script.Ctx) {
			var element = script.Element(PaymentBox).Var(q)

			var secret = script.RPC(func(u user.Ctx) string {
				intent, err := setupintent.New(&stripe.SetupIntentParams{
					PaymentMethodTypes: []*string{
						stripe.String("card"),
					},
				})

				if err != nil {
					u.Report(err)
					return ""
				}

				return intent.ClientSecret
			})(q).Var(q)

			var result = js.Await(element.Get(`stripe`).Call(`confirmCardSetup`, secret, js.NewObject{
				"payment_method": js.NewObject{
					"card": element.Get("card"),
					"billing_details": js.NewObject{
						"name": s.BillingName,
					},
				},
			})).Var(q)

			q.If(js.Bool{Value: result.Get("error")},
				js.Throw(result.Get("error").Get("message")),
			)

			q(s.Value.Set(js.String{
				Value: result.Get("setupIntent").Get("payment_method"),
			}))
			q(signal.Emit(s.confirmedCardSetup))
		}),

		seed.Options(options),
	)
}
