package flash

import (
	html "github.com/multiverse-os/webkit/html"
)

// Notifications may be a better name
//

// TODO: IT shouldn't store the text here, or rather we should be specifying
// what yaml to pull the message from, or i guess pull it and populate this
type Message struct {
	Type  Type
	Value string
}
type Messages []*Message

type Type int

const (
	UndefinedType Type = iota
	AlertType
	AnnouncementType
	ErrorType
	SuccessType
	WarningType
)

// TODO: Then we want a flash.Messages(Alert("message"), Error("test"))

func Alert(message string) Message {
	return Message{Type: AlertType, Value: message}
}

func Announcement(message string) Message {
	return Message{Type: AnnouncementType, Value: message}
}

func Error(message string) Message {
	return Message{Type: ErrorType, Value: message}
}

func Success(message string) Message {
	return Message{Type: SuccessType, Value: message}
}

func Warning(message string) Message {
	return Message{Type: WarningType, Value: message}
}

// TODO: Needs to migrate to new HTML Element declaration
func (ms Messages) String() (output string) {
	if len(ms) > 0 {
		for _, message := range ms {
			output = html.Div.Class("container", "flash-messages").Text(message.String()).String()
		}
	}
	return output
}

// TODO: Obviously going to need a marshal function

func MarshalType(nType string) Type {
	switch nType {
	case AlertType.String():
		return AlertType
	case AnnouncementType.String():
		return AnnouncementType
	case ErrorType.String():
		return ErrorType
	case SuccessType.String():
		return SuccessType
	case WarningType.String():
		return WarningType
	default: // Undefined
		return UndefinedType
	}
}

// TODO: These types may be best definable so they can be established based on
// CSS framework.
func (nType Type) String() string {
	switch nType {
	case AlertType:
		return "alert"
	case AnnouncementType:
		return "announcement"
	case ErrorType:
		return "error"
	case SuccessType:
		return "success"
	case WarningType:
		return "warning"
	default: // Undefined
		return "undefined"
	}
}

func (m Message) String() string {
	return html.Div.Class("flash-message", m.Type.String()).Text(m.Value).String()
}
