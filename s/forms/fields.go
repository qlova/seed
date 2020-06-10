package forms

import (
	"fmt"

	"github.com/qlova/seed"
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/state"
	"github.com/qlova/seed/state/secret"

	"github.com/qlova/seed/s/button"
	"github.com/qlova/seed/s/emailbox"
	"github.com/qlova/seed/s/numberbox"
	"github.com/qlova/seed/s/passwordbox"
	"github.com/qlova/seed/s/text"
	"github.com/qlova/seed/s/textbox"
)

func focusNextField() seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		script.OnEnter(func(q script.Ctx) {
			q(fmt.Sprintf(`{let current = %v;`, script.Element(c)))
			q(`
				let inputs = document.querySelectorAll("form input, form button");
				let found = false;
				for (let input of inputs) {
					if (found && input.tagName == "BUTTON") {
						input.click();
						continue;
					}

					if (found && input.tagName == "INPUT") {
						input.focus();
						found = false;
						break;
					}
					if (input.id == current.id) {
						input.blur();
						found = true;
					}
				}
				if (found) inputs[inputs.length-1].focus();
			}`)
		}).AddTo(c)
	})
}

type FieldTheme struct {
	Title, Box, ErrorText, ErrorBox seed.Options
}

type TextField struct {
	Title, Placeholder string
	Update             state.String

	Checker script.Script

	Required bool

	Theme FieldTheme
}

func (field TextField) AddTo(c seed.Seed) {
	var Error = state.NewString("", state.Global())

	c.With(

		text.New(field.Title, field.Theme.Title),
		textbox.Var(field.Update, field.Theme.Box,
			textbox.SetPlaceholder(field.Placeholder),

			seed.If(field.Required, SetRequired()),

			Error.If(field.Theme.ErrorBox),

			script.OnInput(Error.Set(js.NewString(""))),

			state.Error(Error),

			script.OnChange(field.Checker),

			focusNextField(),

			//How to focus the next field?
			//script.OnEnter(textbox.Focus(EmailBox)),
		),

		Error.If(
			text.New(Error, field.Theme.ErrorText),
		),
	)
}

type FloatField struct {
	Title, Placeholder string
	Update             state.Float

	Checker script.Script

	Required bool

	Theme FieldTheme
}

func (field FloatField) AddTo(c seed.Seed) {
	var Error = state.NewString("", state.Global())

	c.With(

		text.New(field.Title, field.Theme.Title),
		numberbox.Var(field.Update, field.Theme.Box,
			textbox.SetPlaceholder(field.Placeholder),

			seed.If(field.Required, SetRequired()),

			Error.If(field.Theme.ErrorBox),

			script.OnInput(Error.Set(js.NewString(""))),

			state.Error(Error),

			script.OnChange(field.Checker),

			focusNextField(),

			//How to focus the next field?
			//script.OnEnter(textbox.Focus(EmailBox)),
		),

		Error.If(
			text.New(Error, field.Theme.ErrorText),
		),
	)
}

type EmailField struct {
	Title, Placeholder string
	Update             state.String

	Required bool

	Theme FieldTheme
}

func (field EmailField) AddTo(c seed.Seed) {
	var Error = state.NewBool(state.Global())

	var Email = field.Update

	checkEmail := Error.Set(Email.GetString().Includes(js.NewString("@")).Not())

	c.With(
		text.New(field.Title, field.Theme.Title),
		emailbox.Var(field.Update, field.Theme.Box,
			textbox.SetPlaceholder(field.Placeholder),

			seed.If(field.Required, SetRequired()),

			Error.And(Email).If(field.Theme.ErrorBox),

			script.OnInput(Error.Set(js.NewString(""))),

			script.OnChange(checkEmail),

			focusNextField(),

			//How to focus the next field?
			//script.OnEnter(textbox.Focus(EmailBox)),
		),

		Error.And(Email).If(
			text.New("please input a valid email address", field.Theme.ErrorText),
		),
	)
}

type PasswordField struct {
	Required bool

	Theme FieldTheme

	Update  secret.State
	Confirm bool
}

func (field PasswordField) AddTo(c seed.Seed) {
	var Error = state.NewString("", state.Global())

	var Password = field.Update
	var PasswordMismatched = state.NewBool(state.Session())
	var PasswordToConfirm = secret.New(Password.Salt, state.Session())

	checkPassword := PasswordMismatched.Set(Password.GetString().Equals(PasswordToConfirm).Not())

	c.With(

		text.New("Password", field.Theme.Title),
		passwordbox.Var(Password, field.Theme.Box,

			seed.If(field.Required, SetRequired()),

			Error.If(field.Theme.ErrorBox),

			script.OnInput(Error.Set(js.NewString(""))),

			state.Error(Error),

			script.OnChange(checkPassword),

			focusNextField(),

			//How to focus the next field?
			//script.OnEnter(textbox.Focus(EmailBox)),
		),

		Error.If(
			text.New(Error, field.Theme.ErrorText),
		),

		seed.If(field.Confirm,
			text.New("Confirm Password", field.Theme.Title),
			passwordbox.Var(PasswordToConfirm, field.Theme.Box,

				seed.If(field.Required, SetRequired()),

				PasswordMismatched.If(field.Theme.ErrorBox),

				script.OnChange(checkPassword),

				focusNextField(),

				//How to focus the next field?
				//script.OnEnter(textbox.Focus(EmailBox)),
			),

			PasswordMismatched.And(Password).If(
				text.New("this password is different from the one above", field.Theme.ErrorText),
			),
		),
	)
}

type SubmitButton struct {
	Title             string
	Theme, ThemeError seed.Options

	OnSubmit script.Script

	Spinner seed.Seed
}

func (submit SubmitButton) AddTo(c seed.Seed) {
	var Error = state.NewString("", state.Global())
	var Processing = state.New(state.Global())

	c.With(
		Error.If(text.New(Error, submit.ThemeError)),

		Processing.Not().If(
			button.New(submit.Title, submit.Theme,

				state.Error(Error),

				script.OnError(func(q script.Ctx, err script.Error) {
					Processing.Unset(q)
				}),

				script.OnPress(func(q script.Ctx) {
					q.If(script.Element(c).Call("reportValidity"),
						script.New(
							Processing.Set,
							submit.OnSubmit,
							Processing.Unset,
						),
					)
				}),
			),
		),

		Processing.If(
			submit.Spinner,
		),
	)
}
