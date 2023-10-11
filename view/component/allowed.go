package component

import (
	html "github.com/multiverse-os/maglev/html"
	attribute "github.com/multiverse-os/maglev/html/attribute"
)

// ALLOW LIST
///////////////////////////////////////////////////////////////////////////////
// |
// *-[_Elements_(Tags)_List____]: using map[html.Tag]bool
// |
// *-[_Entity_(Character)_List_]: using map[entity.Entity]bool
// |
// '---[_Attributes_(per_tag)__]: Allowed attributes defined per tag / Values for
//     |                           attributes per tag.
//     |
//     '----[Attribute Values (per tag)] Allowed attribute values
//
//
///////////////////////////////////////////////////////////////////////////////
// TODO: Should cache this in a k/v store, perhaps YAML/TOML/JSON -> K/V but the
// K/V would _absolutely have to be_ constantDB or immutable append only to
// avoid tampering during runtime.

// NOTE: This doesn't even need to be a text file, because we are working on
// a web serverw ecan just have a setup backend that lets you edit this text
// via a text form and submit it to possibly even modify the binary if not
// just use one of the embedded databases.

// The concept is to move towards a single allow list for managing the HTML
// rendering whitelist. Potentially this could be built ideally from scranning
// the registered templates, and CSS.
func AllowedListYAML() string {
	return `
allowed:
  no_attributes: # Need a way to allow just the tag without any attributes
	  - "html"
		- "body"
	attributes:
	  div:
	    class:
		    - "columns"
		  	- "coilumn
`
}

func AllowedAttributes(e html.Element) html.Attributes {
	switch e.Tag {
	case html.Meta:
		return html.Attributes{
			attribute.Type:    []string{"description", "keywords", "author", "viewport"},
			attribute.Content: []string{"socialscribe.to", "open,social,subscription,community,collective"},
		}
	default:
		return html.Attributes{}
	}
}

func AllowedTags() html.Tags {
	return html.Tags{
		html.HTML:     true,
		html.Head:     true,
		html.Style:    true,
		html.Meta:     true,
		html.Title:    true,
		html.Body:     true,
		html.Content:  true,
		html.Header:   true,
		html.Nav:      true,
		html.Aside:    true,
		html.Menu:     true,
		html.Main:     true,
		html.Footer:   true,
		html.UL:       true,
		html.OL:       true,
		html.LI:       true,
		html.H1:       true,
		html.H2:       true,
		html.H3:       true,
		html.H4:       true,
		html.P:        true,
		html.Span:     true,
		html.Strong:   true,
		html.Em:       true,
		html.A:        true,
		html.Img:      true,
		html.Div:      true,
		html.Form:     true,
		html.Input:    true,
		html.TextArea: true,
		html.Select:   true,
		html.Option:   true,
		html.Button:   true,
	}
}
