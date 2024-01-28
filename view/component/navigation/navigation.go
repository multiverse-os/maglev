package navigation

import (
	"fmt"
	"strconv"

	html "github.com/multiverse-os/webframe/html"
)

type Sitemap struct {
	Categories []Category
}

type MenuType int
type MenuOptions []MenuOption

const (
	HeaderMenu MenuType = iota
	RightSideMenu
	LeftSideMenu
	FooterMenu
)

func (self MenuType) String() string {
	// TODO: Not sure how I feel about the use of - in the naming
	switch self {
	case RightSideMenu:
		return "right"
	case LeftSideMenu:
		return "left"
	case FooterMenu:
		return "footer"
	default: // HeaderMenu
		return "header"
	}
}

type Menu struct {
	Type        MenuType
	Path        string
	Columns     int
	Classes     []string
	MenuOptions []MenuOption
	Categories  []*Category
}

type MenuOption struct {
	Name    string
	Section string
	Address string
	Count   int
	Classes []string
}

type Category struct {
	Name          string
	Address       string
	Count         int
	Subcategories []*Category
	MenuOptions
}

func NewMenuCategory(name, address string, menuOptions ...MenuOption) *Category {
	return &Category{
		Name:        name,
		Address:     address,
		Count:       len(menuOptions),
		MenuOptions: menuOptions,
	}
}

// NOTE: This is so we cache the value instead of caculate it ever time
func (self *Category) UpdateCount() {
	self.Count = len(self.MenuOptions)
}

func (self *Category) AddSubcategory(category *Category) {
	self.Subcategories = append(self.Subcategories, category)
}

// TODO: Move HTML for navigation to this,
// also check where we are and identify it
// by adding a class to that `nav li`
func NewMenuOption(address, name string) MenuOption {
	return MenuOption{
		Address: address,
		Name:    name,
	}
}

func (self *Category) AddMenuOption(menuOption MenuOption) {
	self.MenuOptions = append(self.MenuOptions, menuOption)
}

func (self Menu) IsEmpty() bool {
	return (len(self.MenuOptions) == 0 && len(self.Categories) == 0)
}

func NewHeaderMenu(menuOptions ...MenuOption) Menu {
	return Menu{
		Type:        HeaderMenu,
		MenuOptions: menuOptions,
	}
}

func NewSideMenu(menuType MenuType, columns int, categories []*Category) Menu {
	return Menu{
		Type:       menuType,
		Columns:    columns,
		Categories: categories,
	}
}

func NewLeftSideMenu(columns int, categories ...*Category) Menu {
	return NewSideMenu(LeftSideMenu, columns, categories)
}

func NewRightSideMenu(columns int, categories ...*Category) Menu {
	return NewSideMenu(RightSideMenu, columns, categories)
}

//	TODO: A footer menu without categories can probably be acheived with a very simple
//
// single category menu that could leverage this func. This could be used for
// all categoriesless menus.
func NewFooterMenu(categories ...*Category) Menu {
	return Menu{
		Type:       FooterMenu,
		Categories: categories,
	}
}

func (self Menu) Render(path string) Menu {
	self.Path = path
	return self
}

func (self Menu) HTML() (menu html.Element) {
	fmt.Println("[MENU:" + self.Type.String() + "]")
	fmt.Println("  [MENU:CATS:COUNT" + strconv.Itoa(len(self.Categories)) + "]")
	var elements html.Elements
	for _, category := range self.Categories {
		fmt.Println("  [MENU:CATEGORY:NAME] " + category.Name)
		fmt.Println("  [MENU:CATEGORY:ITEMS:COUNT] " + strconv.Itoa(len(category.MenuOptions)))
		fmt.Println("  [MENU:SUBCATEGORIES:COUNT] " + strconv.Itoa(len(category.Subcategories)))

		elements = append(elements, html.P.Tags(html.A.Href(category.Address).Text(category.Name)).Class("menu-label"))
		var menuElements html.Elements
		for _, categoryItem := range category.MenuOptions {
			menuElements = append(menuElements, html.LI.Tags(html.A.Text(categoryItem.Name).Href(categoryItem.Address)))
		}
		// Make a list of element.LI, then do a normal element.UL.Tags(LIList) append to html output
		elements = append(elements, html.UL.Tags(menuElements...).Class("menu-list"))
	}

	switch self.Type {
	case LeftSideMenu, RightSideMenu:
		return html.Aside.Tags(elements...).Class("menu-list")
	case FooterMenu:
		return html.P.Text("footer menu goes here").Class("footer", "menu")
	default: // HeaderMenu
		return html.Nav.Tags(elements...)
	}
}

// NOTE: This may not be necessary because bulma at least provides an auto
// sizing column by simply not providing the `is-x` size class
// https://bulma.io/documentation/columns/gap/
// NOTE: Also columns under the 12 count can be centered with bulma using
// `is-centered` class.
// Maybe set max/min size and truncate size as variables since all of it should
// be dependent on CSS styling which will be custom to a specific web
// application or desktop UI project.

// NOTE: I hate this really
func (self Menu) ColumnSize() int {
	// NOTE: Should probably truncate after 5 or 6 since the top bar will not
	// support too many options without using drop-downs (which should be
	// supported via categories).
	switch count := len(self.MenuOptions); count {
	case 0:
		// one large empty placeholder
		return 12
	case 1, 2, 3:
		return 4
	case 4:
		return 3
	case 5, 6:
		return 2
	default: // >= 7
		return 1
	}
}

// TODO: Add support for header menu, side-menu, and footer menu differentiation
// since this portion is more important in render than the HTML() func since its
// adding the HTML that finalized in the HTML() func.
// TODO: Add support for `header` menu to have drop down menus handled using
// categories. Support click to open and close, hover types, CSS only (using
// checkbox).
// NOTE: Show count is used for messages or item counts
func (self Menu) MenuList(showCount bool) (elements html.Elements) {
	// determine what column size should be for horizontal menus
	// by taking the number of items and dividng or turncating
	for _, menuOption := range self.MenuOptions {
		if menuOption.Active(self.Path) {
			menuOption.Classes = append(menuOption.Classes, "active")
		}
		if showCount {
			elements = append(elements, html.Span.Text(strconv.Itoa(menuOption.Count)).Class(menuOption.Classes...))
		}
		// self.Type (MenuType) value, since all three types have different ways of
		// presenting at various screen sizes.
		elements = append(elements, Column(
			self.ColumnSize(),
			html.LI.Tags(
				html.A.Href(menuOption.Address).Text(menuOption.Name),
			).Class(menuOption.Classes...),
		),
		)
	}
	//return html.UL.Tags(Columns(html)).Class(self.Classes...)
	return elements
}

func (self MenuOption) Active(path string) bool {
	return (self.Address == path)
}
