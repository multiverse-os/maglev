package view

import (
	form "github.com/multiverse-os/maglev/view/component/form"

	html "github.com/multiverse-os/webkit/html"
)

// TODO: These need to be returning a view
func Login() html.Element {
	return DefaultTemplate("",
		html.Div.Class("columns").Containing(
			html.Div.Class("column"),
			html.Div.Class("column").Containing(
				html.Div.Class("top-spacer").Containing(
					html.H3.Class("header").Text("Login"),
					html.P.Class("subhead-text").Text("Login or register a new account"),
					LoginForm(),
					html.P.Class("forgot-password").Containing(
						html.A.Href("/users/password/recovery").Text("I forgot my password"),
					),
				),
			),
			html.Div.Class("column"),
		),
	)
}

func Register() html.Element {
	return html.Div.Class("top-spacer").Containing(
		html.H3.Class("header").Text("Register"),
		LoginForm(),
	)
}

//[ Partials ]/////////////////////////////////////////////////////////////////

func LoginForm() html.Element {
	return html.Form.Class("form").Method("POST").Action("/sessions/new").Containing(
		form.TextInput("uid", "Username"),
		form.PasswordInput("pwd", "Password"),
		html.Div.Class("columns").Containing(
			html.Div.Class("column").Containing(
				form.SubmitButton("submit", "Login"),
			),
			html.Div.Class("column").Containing(
				html.A.Class("button", "register-button").Href("/register").Text("Register"),
			),
		),
	)
}
