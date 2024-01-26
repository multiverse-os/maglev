package form

import (
	html "github.com/multiverse-os/webkit/html"
	//attribute "github.com/multiverse-os/webkit/html/framework/attribute"
)

func TextInput(name, placeholder string) html.Element {
	// TODO: Validate name/placeholder
	return html.Div.Class("field", "text-input").Containing(
		html.Div.Class("control").Containing(
			html.Input.Name(name).Placeholder(placeholder).Class("input", ("input-"+name), "is-primary", "login-form").Type("text"),
			html.Div.Class("error"),
		),
	)
}

func PasswordInput(name, placeholder string) html.Element {
	// TODO: Validate name/placeholder
	return html.Div.Class("field", "password-input").Containing(
		html.Div.Class("control").Containing(
			html.Input.Name(name).Placeholder(placeholder).Class("input", ("input-"+name), "is-primary", "login-form").Type("password"),
			html.Div.Class("error"),
		),
	)
}

func SubmitButton(name, text string) html.Element {
	// TODO: Validate name/placeholder
	return html.Div.Class("field", "button-input").Containing(
		html.Div.Class("control").Containing(
			//html.Input.Name(name).Value(text).Class("input", ("input-"+name), "login-form").Type("button"),
			html.Input.Value(text).Class("input", ("input-"+name), "login-form").Type("submit"),
			html.Div.Class("error"),
		),
	)
}

///////////////////////////////////////////////////////////////////////////////

//
// LEGACY
///////////////////////////////////////////////////////////////////////////////

//type InputForm struct {
//	Method          FormMethod
//	Action          string
//	class           html.Attribute // need to init this in NewForm now
//	FormTitle       string
//	FormDescription string
//	Inputs          []InputField
//}
//
//
//func (self FormMethod) Form(action string, inputs ...InputField) InputForm {
//	return InputForm{
//		Method: self,
//		Action: action, // TODO: If Valid!
//		Inputs: inputs,
//		class: html.Attribute{
//			Key:    attribute.Class,
//			Values: html.Values{"form"},
//		},
//	}
//}
//
//// TODO Placeholder because there's stuff calling this function already I guess
//func NewForm(method FormMethod, action string, class html.Attribute, formTitle string, formDesc string, inputs []InputField) InputForm {
//	return InputForm{
//		Method:          method,
//		Action:          action,
//		class:           class,
//		FormTitle:       formTitle,
//		FormDescription: formDesc,
//		Inputs:          inputs,
//	}
//}
//
//////////////////////////////////////////////////////////////////////////////////
//// Input Field HTML/CSS Structure
//type InputFieldComponent int
//
//// TODO: Need a way for horizontal style title/description in its own div
//// floated next to the input field and hint floated next to it. Need a div
//// around the title and description which shouldnt interfere with other styles
//// of layout. This will be the detailsComponent
//// Hint and Error might be combinable but then you cannot display both at the same time
//const (
//	ContainerComponent InputFieldComponent = iota
//	FormComponent                          // Form Tag
//	InputsComponent
//	DetailsComponent
//	TitleComponent
//	DescriptionComponent
//	FieldComponent // remaining portion from details component; input hint error
//	InputComponent
//	HintComponent
//	ErrorComponent
//)
//
//func (self InputFieldComponent) String() string {
//	switch self {
//	case ContainerComponent:
//		return "container"
//	case FormComponent:
//		return "form"
//	case InputsComponent:
//		return "inputs"
//	case DetailsComponent:
//		return "details"
//	case TitleComponent:
//		return "title"
//	case DescriptionComponent:
//		return "description"
//	case InputComponent:
//		return "input"
//	case HintComponent:
//		return "hint"
//	case ErrorComponent:
//		return "error"
//	default: // FieldComponent
//		return "field"
//	}
//}
//
//// TODO: These should be more generic to HTMLElement and inserting into
//// child objects.
//
//func (self InputForm) Class(component InputFieldComponent, class ...string) InputForm {
//	// check if component is valid
//	//self.class[component] = append(self.class[component], class...)
//	//self.class = attribute.Class.Value(class, ...self.class.Vals)
//	return self
//}
//
//func (self InputForm) InputDivClass(class ...string) InputForm {
//	//self.class[InputComponent] = append(self.class[InputComponent], class...)
//	return self
//}
//
//func (self InputForm) InputsClass(class ...string) InputForm {
//	//self.class[InputsComponent] = append(self.class[InputsComponent], class...)
//	return self
//}
//
//func (self InputForm) ContainerClass(class ...string) InputForm {
//	//self.class[ContainerComponent] = append(self.class[ContainerComponent], class...)
//	// TODO: I believe it errors because class map is not yet initiated, it needs
//	// an initialization call in NewForm
//	//self.class[ContainerComponent] = class
//	return self
//}
//
//////////////////////////////////////////////////////////////////////////////////
//
//type SelectOption struct {
//	Name     string
//	Value    string
//	Selected bool
//}
//
//// TODO: Make NewInputField function generic and add the more specific config
//// settings using something like: self.Placeholder("value").Wrap(SoftOption) ..
//// and the resulting chain outputs the InputField.
//// TODO: Going to want to model the nesting better using HTML elements object
//// eventually
//type InputField struct {
//	Type          attribute.TypeOption
//	Title         string
//	Description   string
//	Options       []SelectOption
//	MinLength     int
//	MaxLength     int
//	DefaultOption string
//	Name          string
//	Placeholder   string
//	value         string
//	Wrap          bool
//	WrapOption    WrapOption
//	Selected      bool
//	Errors        []error
//}
//
//func (self InputField) Class(component InputFieldComponent, class ...string) InputField {
//	// TODO: Downcase all class names coming in
//	//self.class[component] = append(self.class[component], class...)
//	return self
//}
//
//// NOTE: Name getting a bit much
//func (self InputField) InputDivClass(class ...string) InputField {
//	// TODO: Downcase all class names coming in
//	//self.class[InputComponent] = append(self.class[InputComponent], class...)
//	return self
//}
//
//func (self InputField) ContainerClass(class ...string) InputField {
//	// TODO: Downcase all class names coming in
//	//self.class[ContainerComponent] = append(self.class[ContainerComponent], class...)
//	return self
//}
//
//func (self InputField) Value(value string) InputField {
//	self.value = value
//	return self
//}
//
//// TODO: Need a way to populate a form with errors or returned data
//// so people dont have to retype every time
//func TextInput(name, placeholder string) InputField {
//	return InputField{
//		//Type:        TextType,
//		//class:       attribute.Class.Value("input", ("input-" + name)),
//		Name:        name,
//		Placeholder: placeholder,
//	}
//}
//
//func PasswordInput(name, placeholder string) InputField {
//	return InputField{
//		//Type:        PasswordType,
//		//class:       attribute.Class.Value("input", ("input-" + name)),
//		Name:        name,
//		Placeholder: placeholder,
//	}
//}
//
//func TextAreaInput(name, placeholder string) InputField {
//	return InputField{
//		//Type:        TextAreaType,
//		//class:       attribute.Class.Value("input", "input-textarea", ("input-" + name)),
//		Name:        name,
//		Placeholder: placeholder,
//	}
//}
//
//func ButtonInput(name string) InputField {
//	return InputField{
//		//Type:  ButtonType,
//		//class: attribute.Class.Value("button", ("input-" + name)),
//		Name: name,
//	}
//}
