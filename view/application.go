package view

import (
	html "github.com/multiverse-os/webframe/html"
)

// New paradigm is views return a VIEW and helpers and other such stuff return
// html.Element; Because a view can be a json output!!! this metaphor is
// starting to go places!

// TODO: I want to prefer returning bytes instead of our type so developers can
// easily use their own system long as it turns in the expected bytes
func Root() html.Element {
	return DefaultTemplate("title",
		html.Div.Class("content").Containing(
			html.Section.Class("section is-fullwidth is-primary").Containing(
				html.H1.Class("title").Text("maglev: Go webframework"),
				html.H5.Class("subtitle").Text("A web application framework designed for simplicity, single binary, single response, inspired by rails"),
			),
		),
	)
}
