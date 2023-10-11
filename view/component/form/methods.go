package form

type FormMethod int
type FormMethods map[FormMethod]bool

const (
	UNDEFINED FormMethod = iota
	POST
	GET
	PUT
	PATCH
	DELETE
)

// TODO: What about initializing the form based on the Method?
// As in, PostMethod.Form()
// TODO: Decide if we should only have the "real" POST/GET then
func (fm FormMethod) String() string {
	switch fm {
	case POST:
		return "POST"
	case GET:
		return "GET"
	case PUT:
		return "PUT"
	case PATCH:
		return "PATCH"
	case DELETE:
		return "DELETE"
	default: // UNDEFINED
		return ""
	}
}

func AllowedFormMethods() FormMethods {
	return FormMethods{
		GET:    true,
		PUT:    false,
		POST:   true,
		PATCH:  false,
		DELETE: false,
	}
}

// TODO: This may be the new method for doing marshals since it is much smaller
func MarshalMethod(methodName string) FormMethod {
	for method, allowed := range AllowedFormMethods() {
		if allowed && method.String() == methodName {
			return method
		}
	}
	return GET // Default to GET
}
