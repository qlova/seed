package forms

import (
	"fmt"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientop"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/client/render"
	"qlova.org/seed/js"
	"qlova.org/seed/script"

	"qlova.org/seed/s/button"
	"qlova.org/seed/s/column"
	"qlova.org/seed/s/emailbox"
	"qlova.org/seed/s/numberbox"
	"qlova.org/seed/s/passwordbox"
	"qlova.org/seed/s/text"
	"qlova.org/seed/s/textbox"
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
	Title, Box, Column, ErrorText, ErrorBox seed.Options
}

type TextField struct {
	Title, Placeholder string

	Update *clientside.String

	Checker script.Script

	Required bool

	Theme FieldTheme
}

func (field TextField) AddTo(c seed.Seed) {
	var Error = new(clientside.String)

	c.With(column.New(
		field.Theme.Column,

		text.New(text.Set(field.Title), field.Theme.Title),
		textbox.New(textbox.Update(field.Update), field.Theme.Box,
			textbox.SetPlaceholder(field.Placeholder),

			seed.If(field.Required, SetRequired()),

			render.If(Error, field.Theme.ErrorBox),

			client.OnInput(Error.SetTo(js.NewString(""))),

			clientside.Catch(Error),

			client.OnChange(field.Checker),

			focusNextField(),

			//How to focus the next field?
			//script.OnEnter(textbox.Focus(EmailBox)),
		),

		render.If(Error,
			text.New(text.SetTo(Error), field.Theme.ErrorText),
		),
	))
}

type FloatField struct {
	Title, Placeholder string

	Update *clientside.Float64

	Checker script.Script

	Required bool

	Theme FieldTheme
}

func (field FloatField) AddTo(c seed.Seed) {
	var Error = new(clientside.String)

	c.With(column.New(
		field.Theme.Column,

		text.New(text.Set(field.Title), field.Theme.Title),
		numberbox.New(numberbox.Update(field.Update), field.Theme.Box,
			textbox.SetPlaceholder(field.Placeholder),

			seed.If(field.Required, SetRequired()),

			render.If(Error, field.Theme.ErrorBox),

			client.OnInput(Error.SetTo(js.NewString(""))),

			clientside.Catch(Error),

			client.OnChange(field.Checker),

			focusNextField(),

			//How to focus the next field?
			//script.OnEnter(textbox.Focus(EmailBox)),
		),

		render.If(Error,
			text.New(text.SetTo(Error), field.Theme.ErrorText),
		),
	))
}

type EmailField struct {
	Title, Placeholder string

	Update *clientside.String

	Required bool

	Theme FieldTheme
}

func (field EmailField) AddTo(c seed.Seed) {
	var Error = new(clientside.Bool)

	var Email = field.Update

	checkEmail := Error.SetTo(Email.GetString().Includes(js.NewString("@")).Not())

	c.With(column.New(
		field.Theme.Column,

		text.New(text.Set(field.Title), field.Theme.Title),
		emailbox.New(textbox.Update(field.Update), field.Theme.Box,
			textbox.SetPlaceholder(field.Placeholder),

			seed.If(field.Required, SetRequired()),

			render.If(clientop.And(Error, Email), field.Theme.ErrorBox),

			client.OnInput(Error.SetTo(js.NewString(""))),

			client.OnChange(checkEmail),

			focusNextField(),

			//How to focus the next field?
			//script.OnEnter(textbox.Focus(EmailBox)),
		),

		render.If(clientop.And(Error, Email),
			text.New(text.Set("please input a valid email address"), field.Theme.ErrorText),
		),
	))
}

type PasswordField struct {
	Title    string
	Required bool

	Theme FieldTheme

	Update  *clientside.Secret
	Confirm bool
}

func (field PasswordField) AddTo(c seed.Seed) {
	var Error = new(clientside.String)

	var Password = field.Update
	var PasswordToConfirm = &clientside.Secret{
		Pepper: Password.Pepper,

		CPU: Password.CPU,
		RAM: Password.RAM,
	}
	var PasswordMismatched = clientop.NotEq(Password, PasswordToConfirm)

	if field.Title == "" {
		field.Title = "Password"
	}

	c.With(column.New(
		field.Theme.Column,

		text.New(text.Set(field.Title), field.Theme.Title),
		passwordbox.New(field.Theme.Box,

			passwordbox.Update(field.Update),

			seed.If(field.Required, SetRequired()),

			render.If(Error, field.Theme.ErrorBox),

			client.OnInput(Error.SetTo(js.NewString(""))),

			clientside.Catch(Error),

			focusNextField(),

			//How to focus the next field?
			//script.OnEnter(textbox.Focus(EmailBox)),
		),

		render.If(Error,
			text.New(text.SetTo(Error), field.Theme.ErrorText),
		),

		seed.If(field.Confirm,
			text.New(text.Set("Confirm "+field.Title), field.Theme.Title),
			passwordbox.New(field.Theme.Box,

				passwordbox.Update(PasswordToConfirm),

				seed.If(field.Required, SetRequired()),

				render.If(PasswordMismatched, field.Theme.ErrorBox),

				focusNextField(),

				//How to focus the next field?
				//script.OnEnter(textbox.Focus(EmailBox)),
			),

			render.If(clientop.And(PasswordMismatched, Password.GetBool()),
				text.New(text.Set("this password is different from the one above"), field.Theme.ErrorText),
			),
		),
	))
}

type SubmitButton struct {
	Title             string
	Theme, ThemeError seed.Options

	OnSubmit client.Script

	Spinner seed.Seed
}

func (submit SubmitButton) AddTo(c seed.Seed) {
	var Error = new(clientside.String)
	var Processing = new(clientside.Bool)

	c.With(
		render.If(Error, text.New(text.SetTo(Error), submit.ThemeError)),

		Processing.Not().If(
			button.New(text.Set(submit.Title), submit.Theme,

				script.OnError(func(q script.Ctx, err script.Error) {
					q(Error.SetTo(err.String))
					q(Processing.Set(false))
				}),

				script.OnPress(func(q script.Ctx) {
					q.If(js.Func("s.form.reportValidity").Call(script.Element(c)),
						client.NewScript(
							Processing.Set(true),
							submit.OnSubmit,
							Processing.Set(false),
						).GetScript(),
					)
				}),
			),
		),

		Processing.If(
			submit.Spinner,
		),
	)
}
