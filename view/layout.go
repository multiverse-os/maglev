package view

import (
	css "github.com/multiverse-os/maglev/view/asset/css"

	html "github.com/multiverse-os/webkit/html"
)

// TODO: Views should use this default template, and extend it for each page by
// adding sections.
// TODO: Needs a helper (reference navigation under template) for oooookkkk
func DefaultTemplate(title string, yield html.Element) html.Element {
	return html.HTML.Containing(
		html.Head.Containing(
			html.Meta.Charset("UTF-8"),
			html.Meta.Name("description").Content("A powerful web application framework"),
			html.Meta.Name("keywords").Content("powerful, go language, webapp, maglev"),
			html.Meta.Name("viewport").Content("width=device-width, initial-scale=1.0"),
			html.Title.Text(title),
			html.Style.Text(css.DefaultCSS()),
		),
		html.Body.Containing(
			html.Nav.Id("navbar").Class("navbar").Containing(
				html.Div.Class("container").Containing(
					html.Div.Class("navbar-brand").Containing(
						html.H1.Class("title").Containing(
							html.A.Class("site-title").Href("/").Text("Maglev"),
						),
					),
					html.Div.Class("navbar-menu").Containing(
						html.Div.Class("navbar-start").Containing(
							html.A.Class("navbar-item").Href("/").Text("Home"),
							html.A.Class("navbar-item").Href("/about").Text("About"),
							html.A.Class("navbar-item").Href("/contact").Text("Contact"),
						),
						html.Div.Class("navbar-end").Containing(
							html.A.Class("navbar-item").Href("/login").Text("Login"),
							html.A.Class("navbar-item").Href("/register").Text("Register"),
							html.A.Class("navbar-item").Href("https://github.com/multiverse-os/maglev").Containing(
								html.Span.Class("icon").Containing(
									html.I.Class("material-icons").Text("code"),
								),
							),
						),
					), // NavbarMenu
				), // Container
			), // Navbar
			html.Main.Class("bd-main").Containing(
				html.Div.Class("bd-side-background"),
				html.Div.Class("bd-main-container", "container").Containing(
					html.Div.Class("columns").Containing(
						html.Div.Class("column", "is-four-fifths").Containing(
							html.Div.Class("content").Containing(
								yield,
							),
						),
						html.Div.Class("column", "bd-notification", "is-wajofij", "sidebar-menu").Containing(
							html.UL.Containing(
								html.LI.Class("sidebar-header").Text("Community Projects"),
								html.LI.Class("sidebar-li").Text("Open Source Software"),
								html.LI.Class("sidebar-li").Text("Podcasts"),
								html.LI.Class("sidebar-li").Text("Writing"),
								html.LI.Class("sidebar-li").Text("Comics"),
								html.LI.Class("sidebar-li").Text("Music"),
								html.LI.Class("sidebar-li").Text("Illustrations"),
								html.LI.Class("sidebar-li").Text("Video"),
								html.LI.Class("sidebar-li").Text("Organizing"),
								html.LI.Class("sidebar-li").Text("Activism"),
								html.LI.Class("sidebar-header").Text("Socialscribe"),
								html.LI.Class("sidebar-li").Text("About Us"),
								html.LI.Class("sidebar-li").Text("Transparency"),
								html.LI.Class("sidebar-li").Text("Contact"),
								html.LI.Class("sidebar-li").Text("Get Involved"),
								html.LI.Class("sidebar-li").Text("Source Code"),
							),
						),
					),
					html.Div.Class("columns").Containing(
						html.Div.Class("column", "is-two-fifths").Containing(
							html.H4.Class("footer-title").Text("A web framework"),
							html.P.Class("footer-description").Text("Desktop user interfaces using webkit to render the single page responses that focus on HTML and CSS to avoid introducing potential security vulnerabilities."),
							html.P.Class("footer-description").Text("And maximize preformance but also provide a simple portable design. This text is essentially lorem text."),
						),
						html.Div.Class("column"),
						html.Div.Class("column").Containing(
							html.H4.Class("footer-title").Text("Community Projects"),
							html.UL.Class("footer-ul").Containing(
								html.LI.Class("footer-li").Text("Software"),
								html.LI.Class("footer-li").Text("Podcasts"),
								html.LI.Class("footer-li").Text("Writing"),
								html.LI.Class("footer-li").Text("Comics"),
							),
						),
						html.Div.Class("column").Containing(
							html.UL.Class("footer-ul").Containing(
								html.LI.Class("footer-li").Text("Music"),
								html.LI.Class("footer-li").Text("Illustrations"),
								html.LI.Class("footer-li").Text("Video"),
								html.LI.Class("footer-li").Text("Organizing"),
								html.LI.Class("footer-li").Text("Activism"),
							),
						),
						html.Div.Class("column").Containing(
							html.H4.Class("footer-title").Text("Socialscribe"),
							html.UL.Class("footer-ul").Containing(
								html.LI.Class("footer-li").Text("Transparency"),
								html.LI.Class("footer-li").Text("Contact"),
								html.LI.Class("footer-li").Text("Get Involved"),
								html.LI.Class("footer-li").Text("Source Code"),
							),
						),
					),
				), // bdMainContainer
			),
		), // Body
	)
}

////////////////////////////////////////////////////////////////////////////////
// Forms

//	return ""
//	//return form.NewForm(
//	//	action,
//	//	// NOTE: Not a big fan of this ContainerClass(), .. other classes
//	//	// that the user is expected to know but really are at least somewhat
//	//	// specific, and long ontop of it
//	//	form.TextInput("uid", "Username").ContainerClass("is-5"),
//	//	form.PasswordInput("password", "Password").ContainerClass("is-5"),
//	//	form.ButtonInput("Login").ContainerClass("is-1"),
//	//	form.ButtonInput("Register").ContainerClass("is-1"),
//	//).ContainerClass("field", "column", "is-1").InputDivClass("control").InputsClass("login", "columns", "is-gapless").HTML()
