package component

import (
	"strconv"

	html "github.com/multiverse-os/webkit/html"
)

func Columns(elements ...html.Element) html.Element {
	return html.Div.Tags(elements...).Class("columns")
}

func Column(size int, elements ...html.Element) html.Element {
	return html.Div.Tags(elements...).Class("column", ("is-" + strconv.Itoa(size)))
}

func EmptyColumn() html.Element { return html.Div.Class("column") }
