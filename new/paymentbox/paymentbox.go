//Package paymentbox provides a (stripe) secured input that accepts payment details.
package paymentbox

import (
	"github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/setupintent"
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/web/html"
	"qlova.org/seed/web/js"
	"qlova.org/seed/new/html/div"
)

var StripePublishableKey string

type StripeBox struct {
	BillingName, Value *clientside.String
	confirmCardSetup   *clientside.Signal
	confirmedCardSetup *clientside.Signal
}

func NewStripe() StripeBox {
	return StripeBox{
		BillingName:        new(clientside.String),
		Value:              new(clientside.String),
		confirmCardSetup:   new(clientside.Signal),
		confirmedCardSetup: new(clientside.Signal),
	}
}

//ConfirmCardSetup tells the stripebox to confirm the user provided details.
func (s StripeBox) ConfirmCardSetup() client.Script {
	return s.confirmCardSetup
}

//OnCardConfirmed runs when the card has been confirmed.
func (s StripeBox) OnCardConfirmed(do js.Script) seed.Option {
	return s.confirmedCardSetup.On(do)
}

//New returns a new stripe payment box.
func (s StripeBox) New(options ...seed.Option) seed.Seed {

	var PaymentBox = div.New()

	var Stripe = js.Function{js.NewValue(`Stripe`)}

	return PaymentBox.With(
		js.Require("https://js.stripe.com/v3/", ""),

		client.OnLoad(js.Script(func(q js.Ctx) {
			var element = html.Element(PaymentBox).Var(q)
			q(element.Set(`stripe`, js.Call(Stripe, js.NewString(StripePublishableKey))))
			var elements = element.Get("stripe").Call("elements").Var(q)
			q(element.Set(`card`, elements.Call(`create`, js.NewString("card"), js.NewObject(make(map[string]js.AnyValue)))))
			q(element.Get(`card`).Run(`mount`, element))
		})),

		s.confirmCardSetup.On(js.Script(func(q js.Ctx) {
			var element = html.Element(PaymentBox).Var(q)

			var secret = client.Call(func() (string, error) {
				intent, err := setupintent.New(&stripe.SetupIntentParams{
					PaymentMethodTypes: []*string{
						stripe.String("card"),
					},
				})

				if err != nil {
					return "", err
				}

				return intent.ClientSecret, nil
			}).GetValue().Var(q)

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
			q(s.confirmedCardSetup)
		})),

		seed.Options(options),
	)
}
