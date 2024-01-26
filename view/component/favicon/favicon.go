package favicon

import (
	html "github.com/multiverse-os/webkit/html"
)

func (self *Template) FaviconTag() html.Element {
	return BlankFavicon()
}

func BlankFavicon() html.Element {
	return html.A.Id("favicon").Relative("shortcut icon").Type("image/png").Href("data:image/png;base64,....==")
}

func Favicon(imageData string) html.Element {
	return html.A.Id("favicon").Relative("shortcut icon").Type("image/png").Href(imageData)
}
