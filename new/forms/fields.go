package forms

import (
	"fmt"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/set/change"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/client/if/all"
	"qlova.org/seed/client/if/not"
	"qlova.org/seed/client/if/the"
	"qlova.org/seed/set/visible"
	"qlova.org/seed/use/html"
	"qlova.org/seed/use/js"

	"qlova.org/seed/new/button"
	"qlova.org/seed/new/column"
	"qlova.org/seed/new/emailbox"
	"qlova.org/seed/new/numberbox"
	"qlova.org/seed/new/passwordbox"
	"qlova.org/seed/new/text"
	"qlova.org/seed/new/textarea"
	"qlova.org/seed/new/textbox"
)

func focusNextField() seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		client.OnEnterKey(js.Script(func(q js.Ctx) {
			q(fmt.Sprintf(`{let current = %v;`, html.Element(c)))
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
		})).AddTo(c)
	})
}

type FieldTheme struct {
	Title, Box, Area, Column, ErrorText, ErrorBox seed.Options
}

type TextField struct {
	Title, Placeholder string

	Update *clientside.String

	Checker client.Script

	Required, Multiline bool

	Theme FieldTheme
}

func (field TextField) AddTo(c seed.Seed) {
	var Error = new(clientside.String)

	var box = textbox.New
	var theme = field.Theme.Box

	if field.Multiline {
		box = textarea.New
		theme = field.Theme.Area
	}

	c.With(column.New(
		field.Theme.Column,

		text.New(text.SetString(field.Title), field.Theme.Title),
		box(textbox.Update(field.Update), theme,
			textbox.SetPlaceholder(field.Placeholder),

			seed.If(field.Required, SetRequired()),

			change.When(Error, field.Theme.ErrorBox),

			client.OnInput(Error.SetTo(js.NewString(""))),

			clientside.Catch(Error),

			client.OnChange(field.Checker),

			focusNextField(),

			//How to focus the next field?
			//script.OnEnter(textbox.Focus(EmailBox)),
		),

		change.When(Error,
			text.New(text.SetStringTo(Error), field.Theme.ErrorText),
		),
	))
}

type FloatField struct {
	Title, Placeholder string

	Update *clientside.Float64

	Checker client.Script

	Required bool

	Theme FieldTheme
}

func (field FloatField) AddTo(c seed.Seed) {
	var Error = new(clientside.String)

	c.With(column.New(
		field.Theme.Column,

		text.New(text.SetString(field.Title), field.Theme.Title),
		numberbox.New(numberbox.Update(field.Update), field.Theme.Box,
			textbox.SetPlaceholder(field.Placeholder),

			seed.If(field.Required, SetRequired()),

			change.When(Error, field.Theme.ErrorBox),

			client.OnInput(Error.SetTo(js.NewString(""))),

			clientside.Catch(Error),

			client.OnChange(field.Checker),

			focusNextField(),

			//How to focus the next field?
			//script.OnEnter(textbox.Focus(EmailBox)),
		),

		change.When(Error,
			text.New(text.SetStringTo(Error), field.Theme.ErrorText),
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

		text.New(text.SetString(field.Title), field.Theme.Title),
		emailbox.New(textbox.Update(field.Update), field.Theme.Box,
			textbox.SetPlaceholder(field.Placeholder),

			seed.If(field.Required, SetRequired()),

			change.When(all.AreTrue(Error, Email), field.Theme.ErrorBox),

			client.OnInput(Error.SetTo(js.NewString(""))),

			client.OnChange(checkEmail),

			focusNextField(),

			//How to focus the next field?
			//script.OnEnter(textbox.Focus(EmailBox)),
		),

		visible.When(all.AreTrue(Error, Email),
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
	var PasswordMismatched = the.ValueOf(Password).IsNot(PasswordToConfirm)

	if field.Title == "" {
		field.Title = "Password"
	}

	c.With(column.New(
		field.Theme.Column,

		text.New(text.SetString(field.Title), field.Theme.Title),
		passwordbox.New(field.Theme.Box,

			passwordbox.Update(field.Update),

			seed.If(field.Required, SetRequired()),

			change.When(Error, field.Theme.ErrorBox),

			client.OnInput(Error.SetTo(js.NewString(""))),

			clientside.Catch(Error),

			focusNextField(),

			//How to focus the next field?
			//script.OnEnter(textbox.Focus(EmailBox)),
		),

		change.When(Error,
			text.New(text.SetStringTo(Error), field.Theme.ErrorText),
		),

		seed.If(field.Confirm,
			text.New(text.SetString("Confirm "+field.Title), field.Theme.Title),
			passwordbox.New(field.Theme.Box,

				passwordbox.Update(PasswordToConfirm),

				seed.If(field.Required, SetRequired()),

				change.When(PasswordMismatched, field.Theme.ErrorBox),

				focusNextField(),

				//How to focus the next field?
				//script.OnEnter(textbox.Focus(EmailBox)),
			),

			visible.When(all.AreTrue(PasswordMismatched, Password.GetBool()),
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
		visible.When(Error, text.New(text.SetStringTo(Error), submit.ThemeError)),

		visible.When(not.True(Processing),
			button.New(text.SetString(submit.Title), submit.Theme,

				client.OnError(func(err client.String) client.Script {
					return client.NewScript(
						Error.SetTo(err),
						Processing.Set(false),
					)
				}),

				client.OnClick(js.Script(func(q js.Ctx) {
					q.If(js.Func("s.form.reportValidity").Call(html.Element(c)),
						client.NewScript(
							Processing.Set(true),
							Error.Set(""),
							submit.OnSubmit,
							Processing.Set(false),
						).GetScript(),
					)
				})),
			),
		),

		visible.When(Processing,
			submit.Spinner,
		),
	)
}
