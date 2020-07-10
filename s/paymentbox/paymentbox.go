//Package paymentbox provides a (stripe) secured input that accepts payment details.
package paymentbox

import (
	"github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/setupintent"
	"qlova.org/seed"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/js"
	"qlova.org/seed/s/html/div"
	"qlova.org/seed/script"
	"qlova.org/seed/signal"
)

var StripePublishableKey string

type StripeBox struct {
	BillingName, Value *clientside.String
	confirmCardSetup   signal.Type
	confirmedCardSetup signal.Type
}

func NewStripe() StripeBox {
	return StripeBox{
		BillingName:        new(clientside.String),
		Value:              new(clientside.String),
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

			var secret = script.RPC(func() (string, error) {
				intent, err := setupintent.New(&stripe.SetupIntentParams{
					PaymentMethodTypes: []*string{
						stripe.String("card"),
					},
				})

				if err != nil {
					return "", err
				}

				return intent.ClientSecret, nil
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

			q(s.Value.SetTo(js.String{
				Value: result.Get("setupIntent").Get("payment_method"),
			}))
			q(signal.Emit(s.confirmedCardSetup))
		}),

		seed.Options(options),
	)
}
